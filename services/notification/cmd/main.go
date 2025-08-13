package main

import (
	"context"
	"github.com/casiomacasio/notes-platform/services/notification/internal/events"
	"github.com/casiomacasio/notes-platform/services/notification/internal/handler"
	"github.com/casiomacasio/notes-platform/services/notification/internal/repository"
	"github.com/casiomacasio/notes-platform/services/notification/internal/service"
	"github.com/casiomacasio/notes-platform/services/notification/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to read config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}

	conn, err := amqp.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		logrus.Fatalf("failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	eventBus, err := events.NewRabbitMQBus(conn)
	if err != nil {
		logrus.Fatalf("failed to init event bus: %s", err)
	}

	mongoClient, mongoDB, err := repository.ConnectMongo()
	defer mongoClient.Disconnect(context.Background())
	notificationRepos := repository.NewRepository(mongoDB)
	notificationService := service.NewService(notificationRepos)
	authHandler := handler.NewHandler(notificationService)
	go func() {
		if err := eventBus.Consume("notifications", notificationService.HandleNotificationMessage); err != nil {
			logrus.Fatalf("failed to consume messages: %s", err)
		}
	}()
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), authHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %v", err.Error())
		}
	}()
	logrus.Print("Notification Microservice Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Notification Microservice Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occurred when shutting down the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
