package models

import (
	"time"
)

type SlideImage struct {
	ID         int       `json:"id"`
	CategoryID string    `json:"category_id"`
	ReferrerID string    `json:"referrer_id"`
	Img        string    `json:"img"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateSlideImage struct {
	CategoryID string `json:"category_id"`
	ReferrerID string `json:"referrer_id"`
	Img        string `json:"img"`
}

type UpdateSlideImage struct {
	ID         int    `json:"id"`
	CategoryID string `json:"category_id"`
	ReferrerID string `json:"referrer_id"`
	Img        string `json:"img"`
}
