package repository

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
)

type User interface{
	GetUser(userId int) (model.User, error)
	UpdateUser(useId int, input model.UpdateUserInput) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}