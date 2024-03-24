package models

import "time"

type Payment struct {
	ID             int       `json:"id"  gorm:"primary_key"`
	Order_ID       int       `json:"order_id"`
	Amount         float64   `json:"amount"`
	Payment_Date   string    `json:"payment_date"`
	Payment_Method string    `json:"payment_method"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreatePayment struct {
	Order_ID       int     `json:"order_id"`
	Amount         float64 `json:"amount" binding:"required"`
	Payment_Date   string  `json:"payment_date" binding:"required"`
	Payment_Method string  `json:"payment_method" binding:"required"`
}

type UpdatePayment struct {
	Order_ID       int     `json:"order_id"`
	Amount         float64 `json:"amount"`
	Payment_Date   string  `json:"payment_date"`
	Payment_Method string  `json:"payment_method"`
}
