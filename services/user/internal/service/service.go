package service

import (
	"github.com/casiomacasio/notes-platform/services/user/internal/repository"
	"github.com/casiomacasio/notes-platform/services/user/internal/model"
)

type User interface {
	GetUser(userId int) (model.User, error)
	UpdateUser(userId int, input model.UpdateUserInput) error
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
		User: NewUserService(repos.User),
		Authorization: NewAuthService(),
	}
}