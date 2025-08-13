package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Notification interface {
	SaveNotification(data interface{}) error
	GetNotificationsByUserId(userID string) ([]interface{}, error)
}

type Repository struct {
	Notification
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Notification: NewNotificationMongo(db),
	}
}
