package main

import (
	"context"
	"os/signal"
	"syscall"
    "github.com/casiomacasio/notes-platform/services/note/internal/handler"
    "github.com/casiomacasio/notes-platform/services/note/internal/service"
    "github.com/casiomacasio/notes-platform/services/note/internal/repository"
    "github.com/casiomacasio/notes-platform/services/note/server"
    "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
)

func main() {
    if err := initConfig(); err != nil {
        logrus.Fatalf("failed to read config: %s", err.Error())
    }

    if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
    } 
   
    db, err := repository.NewPostgresDB(repository.Config{ 
        Host:     viper.GetString("postgres.db.host"),
        Port:     viper.GetString("postgres.db.port"),
        Username: viper.GetString("postgres.db.username"),
		Password: os.Getenv("DB_PASSWORD"),
        DBName:   viper.GetString("postgres.db.dbname"),
        SSLMode:  viper.GetString("postgres.db.sslmode"),
    })
    if err != nil {
        logrus.Fatalf("failed to connect to db: %s", err.Error())
    }
    userRepos := repository.NewRepository(db)
    userService := service.NewService(userRepos)
    userHandler := handler.NewHandler(userService)
	srv := new(server.Server)
	go func () {
		if err := srv.Run(viper.GetString("port"), userHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %v", err.Error())
		}
	}()
	logrus.Print("Note Microservice Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("Note Microservice Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured when shutting down the server %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured when closing db connection %s", err.Error())
	}
}

func initConfig() error {
    viper.AddConfigPath("./configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}