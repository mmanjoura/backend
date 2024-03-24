package models

import (
	"time"
)

type Itinerary struct {
	ID         int       `json:"id"`
	CategoryID string    `json:"category_id"`
	ReferrerID string    `json:"referrer_id"`
	Img        string    `json:"img"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateItinerary struct {
	ID         int    `json:"id"`
	CategoryID string `json:"category_id"`
	ReferrerID string `json:"referrer_id"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

type UpdateItinerary struct {
	ID         int    `json:"id"`
	CategoryID string `json:"category_id"`
	ReferrerID string `json:"referrer_id"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}
