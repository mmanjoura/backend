package models

import (
	"time"
)

type Image struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	ReferrerID int       `json:"referrer_id"`
	Img        string    `json:"img"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateImage struct {
	CategoryID int    `json:"category_id"`
	ReferrerID int    `json:"referrer_id"`
	Img        string `json:"img"`
}

type UpdateImage struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	ReferrerID int    `json:"referrer_id"`
	Img        string `json:"img"`
}
