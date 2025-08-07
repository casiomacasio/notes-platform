package model

import "time"

type User struct {
    Id        int    `json:"id" db:"id"`
    Name  string    `json:"name" db:"name"`
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
