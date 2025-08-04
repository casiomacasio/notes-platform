package repository

import (
	"time"
    "github.com/casiomacasio/notes-platform/services/auth/internal/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	refreshTokensTable = "refresh_tokens"
)

type Authorization interface {
	CreateUser(user model.CreateUserRequest) (int, error) 
	GetUser(email, password string) (model.User, error) 
	SaveRefreshToken(hashed_token string, userId int, expires_at time.Time) (uuid.UUID, error)
	GetUserIdAndHashByRefreshTokenId(refreshToken uuid.UUID) (int, string, error) 
	DeleteRefreshToken(refreshToken uuid.UUID) error 
	RevokeRefreshToken(tokenUUID uuid.UUID) (bool, error)
	RevokeRefreshTokenByUserId(userId int) (bool, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}