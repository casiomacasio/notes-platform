package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/casiomacasio/notes-platform/services/notification/internal/handler"
	"github.com/casiomacasio/notes-platform/services/notification/internal/server"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("failed to read config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}

	s, err := server.NewRabbitMQConsumer(
		viper.GetString("rabbitmq.url"),
		viper.GetString("rabbitmq.queue"),
	)
	if err != nil {
		logrus.Fatalf("could not start consumer: %v", err)
	}

	go func() {
		s.Listen(func(event, payload string) {
			consumer.HandleMessage(event, payload)
		})
	}()

	logrus.Println("Notification Microservice Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Notification Microservice Shutting Down")

	if err := s.Shutdown(); err != nil {
		logrus.Errorf("error while shutting down: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
