package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationMongo struct {
	collection *mongo.Collection
}

func NewNotificationMongo(db *mongo.Database) *notificationMongo {
	return &notificationMongo{
		collection: db.Collection("notifications"),
	}
}

func (r *notificationMongo) SaveNotification(data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, data)
	return err
}

func (r *notificationMongo) GetNotificationsByUserId(userID string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []interface{}
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}
