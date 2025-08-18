package model

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Bio       *string   `json:"bio,omitempty" db:"bio"`
	AvatarURL *string   `json:"avatar_url,omitempty" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserInput struct {
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Bio       *string `json:"bio,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type UpdateUserInput struct {
	Name      *string `json:"name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserCreatedEvent struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
