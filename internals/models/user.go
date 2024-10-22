package models

import "time"

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
	JoinedAt time.Time `json:"joinedAt"`
}

type UserRegisterOrUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=32,startsWithLetter"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Address   string `json:"address" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32,startsWithLetter"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
