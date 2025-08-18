package service

import (
	"github.com/casiomacasio/notes-platform/services/auth/internal/model"
	"github.com/casiomacasio/notes-platform/services/auth/internal/repository"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(user model.CreateUserRequest) (int, error)
	GetUserByRefreshTokenAndRefreshTokenId(refresh_token string, refreshTokenUUID uuid.UUID) (int, error)
	GetUserByRefreshTokenId(refreshTokenUUID uuid.UUID) (int, error)
	ParseToken(token string) (int, error)
	GetUser(email, password string) (model.User, error)
	GenerateToken(userId int) (string, error)
	GenerateRefreshToken(userId int) (string, string, error)
	RevokeRefreshToken(uuid.UUID) error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
