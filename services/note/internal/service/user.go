package service

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/casiomacasio/notes-platform/services/note/internal/repository"
)

type UserService struct {
	repo        repository.User
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
		return  err
	}
	return  nil
}