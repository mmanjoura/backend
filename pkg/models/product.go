package models

import "time"

type Product struct {
	ID           int       `json:"id"  gorm:"primary_key"`
	Product_Name string    `json:"product_name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateProduct struct {
	Product_Name string  `json:"product_name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Category     string  `json:"category" binding:"required"`
}

type UpdateProduct struct {
	Product_Name string  `json:"product_name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
}
