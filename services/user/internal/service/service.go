package service

import (
	"github.com/casiomacasio/notes-platform/services/user/internal/model"
	"github.com/casiomacasio/notes-platform/services/user/internal/repository"
	"github.com/streadway/amqp"
)

type User interface {
	GetUser(userId int) (model.User, error)
	UpdateUser(userId int, input model.UpdateUserInput) error
	HandleUserCreated(msg amqp.Delivery)
}

type Authorization interface {
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	User
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:          NewUserService(repos.User),
		Authorization: NewAuthService(),
	}
}
