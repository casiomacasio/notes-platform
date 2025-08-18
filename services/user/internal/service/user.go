package service

import (
	"encoding/json"
	"github.com/casiomacasio/notes-platform/services/user/internal/model"
	"github.com/casiomacasio/notes-platform/services/user/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(userId int) (model.User, error) {
	user, err := s.repo.GetUser(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(userId int, input model.UpdateUserInput) error {
	err := s.repo.UpdateUser(userId, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) HandleUserCreated(msg amqp.Delivery) {
	var event model.UserCreatedEvent
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		logrus.Printf("failed to parse UserCreatedEvent: %v", err)
		return
	}
	err := s.repo.CreateUser(event.UserId, event.Name, event.Email)
	if err != nil {
		logrus.Printf("failed to save user: %v", err)
		return
	}
	logrus.Printf("User created successfully in User service: %s", event.UserId)
}
