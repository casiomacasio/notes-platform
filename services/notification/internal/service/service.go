package service

import (
	"github.com/casiomacasio/notes-platform/services/notification/internal/model"
	"github.com/casiomacasio/notes-platform/services/notification/internal/repository"
)

type Notification interface {
	GetAllNotifications() ([]model.Notification, error)
	GetAllNotificationsByUserId(userId int) ([]model.Notification, error)
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
		// Notification: NewNotificationService(repos.Note),
		Authorization: NewAuthService(),
	}
}
