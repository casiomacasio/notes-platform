package service

import (
	"github.com/casiomacasio/notes-platform/services/notification/internal/model"
	"github.com/casiomacasio/notes-platform/services/notification/internal/repository"
	"github.com/streadway/amqp"
)

type Notification interface {
	GetAllNotifications() ([]model.Notification, error)
	GetAllNotificationsByUserId(userId int) ([]model.Notification, error)
	HandleNotificationMessage(msg amqp.Delivery)
}

type Authorization interface {
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Notification
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Notification:  NewNotificationService(repos.Notification),
		Authorization: NewAuthService(),
	}
}
