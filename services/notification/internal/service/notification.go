package service

import (
	"encoding/json"
	"github.com/casiomacasio/notes-platform/services/notification/internal/model"
	"github.com/casiomacasio/notes-platform/services/notification/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetAllNotifications() ([]model.Notification, error) {
	return s.repo.GetAllNotifications()
}

func (s *NotificationService) GetAllNotificationsByUserId(userId int) ([]model.Notification, error) {
	return s.repo.GetAllNotificationsByUserId(userId)
}

func (s *NotificationService) HandleNotificationMessage(msg amqp.Delivery) {
	var n model.Notification
	if err := json.Unmarshal(msg.Body, &n); err != nil {
		logrus.Printf("failed to unmarshal notification: %v", err)
		return
	}
	if err := s.repo.SaveNotification(n); err != nil {
		logrus.Printf("failed to save notification: %v", err)
	}
}
