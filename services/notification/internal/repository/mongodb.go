package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMongo() (*mongo.Client, *mongo.Database) {
	mongoURI := viper.GetString("mongodb.url")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logrus.Fatalf("failed to connect to MongoDB: %s", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logrus.Fatalf("failed to ping MongoDB: %s", err)
	}

	logrus.Println("Connected to MongoDB")
	db := client.Database(viper.GetString("mongodb.dbname"))
	return client, db
}
