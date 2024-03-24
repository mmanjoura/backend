package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Subject   string    `json:"subject"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
