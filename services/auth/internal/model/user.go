package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreatedEvent struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
