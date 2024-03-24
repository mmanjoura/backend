package models

import "time"

type HotelImage struct {
	ID        int       `json:"id"  gorm:"primary_key"`
	Hotel_ID  int       `json:"hotel_ID"`
	SlideImg  string    `json:"slide_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateHotelImage struct {
	Product_ID  int    `json:"product_ID"`
	Customer_ID int    `json:"customer_ID"`
	Rating      int    `json:"rating" binding:"required"`
	Review_Text string `json:"reviewText" binding:"required"`
}
type UpdateHotelImage struct {
	Product_ID  int    `json:"product_ID"`
	Customer_ID int    `json:"customer_ID"`
	Rating      int    `json:"rating"`
	Review_Text string `json:"reviewText"`
}
