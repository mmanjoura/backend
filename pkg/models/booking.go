package models

import "time"

type Booking struct {
	ID          int       `json:"id"`
	ProductID   string    `json:"product_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	ProductType string    `json:"product_type"`
	NumAdult    string    `json:"num_adult"`
	NumChildren string    `json:"num_children"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
