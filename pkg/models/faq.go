package models

import "time"

type Faq struct {
	ID             int       `json:"id"`
	CollapseTarget string    `json:"collapse_target"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	CategoryID     string    `json:"category_id"`
	ReferrerID     string    `json:"referrer_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
