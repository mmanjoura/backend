package models

import "time"

type Order struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Customer_ID  int       `json:"customer_id"`
	Total_Amount float64   `json:"total_amount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateOrder struct {
	Customer_ID  int     `json:"customer_id"`
	Total_Amount float64 `json:"total_amount" binding:"required"`
	Status       string  `json:"status" binding:"required"`
}
type UpdateOrder struct {
	Customer_ID  int     `json:"customer_id"`
	Total_Amount float64 `json:"total_amount"`
	Status       string  `json:"status"`
}
