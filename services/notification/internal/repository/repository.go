package repository

import (
	"github.com/casiomacasio/notes-platform/services/notification/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Notification interface {
	SaveNotification(data model.Notification) error
	GetAllNotifications() ([]model.Notification, error)
	GetAllNotificationsByUserId(userID int) ([]model.Notification, error)
}

type Repository struct {
	Notification
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Notification: NewNotificationMongo(db),
	}
}
