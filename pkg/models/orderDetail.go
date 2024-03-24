package models

import "time"

type OrderDetail struct {
	ID         int       `json:"id"  gorm:"primary_key"`
	Order_ID   int       `json:"order_id"`
	Product_ID int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Subtotal   float64   `json:"subtotal"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateOrderDetail struct {
	Order_ID   int     `json:"order_id"`
	Product_ID int     `json:"product_id"`
	Quantity   int     `json:"quantity" binding:"required"`
	Subtotal   float64 `json:"subtotal" binding:"required"`
}
type UpdateOrderDetail struct {
	Order_ID   int     `json:"order_id"`
	Product_ID int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Subtotal   float64 `json:"subtotal"`
}
