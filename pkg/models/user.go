package models

import "time"

type LoginUser struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
