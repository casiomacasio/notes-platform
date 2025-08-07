package repository

import (
	"database/sql"
	"fmt"
	"errors"
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserNotFound    = errors.New("user not found")
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUser(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT id, name, email, bio, avatar_url, created_at, updated_at FROM %s WHERE id=$1`, usersTable)
	err := r.db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *UserPostgres) UpdateUser(userId int, input model.UpdateUserInput) error {
	query := fmt.Sprintf("UPDATE %s SET", usersTable)
	args := []interface{}{}
	argIdx := 1

	if input.Name != nil {
		query += fmt.Sprintf(" name = $%d,", argIdx)
		args = append(args, *input.Name)
		argIdx++
	}
	if input.Email != nil {
		query += fmt.Sprintf(" email = $%d,", argIdx)
		args = append(args, *input.Email)
		argIdx++
	}
	if input.Bio != nil {
		query += fmt.Sprintf(" bio = $%d,", argIdx)
		args = append(args, *input.Bio)
		argIdx++
	}
	if input.AvatarURL != nil {
		query += fmt.Sprintf(" avatar_url = $%d,", argIdx)
		args = append(args, *input.AvatarURL)
		argIdx++
	}

	query += " updated_at = NOW()"
	query += fmt.Sprintf(" WHERE id = $%d", argIdx)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
