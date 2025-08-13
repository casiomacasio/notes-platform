package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/casiomacasio/notes-platform/services/notification/internal/model"
)

type notificationMongo struct {
	collection *mongo.Collection
}

func NewNotificationMongo(db *mongo.Database) *notificationMongo {
	return &notificationMongo{
		collection: db.Collection("notifications"),
	}
}

func (r *notificationMongo) GetAllNotifications() ([]model.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []model.Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationMongo) GetAllNotificationsByUserId(userId int) ([]model.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []model.Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationMongo) SaveNotification(n model.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, n)
	return err
}
