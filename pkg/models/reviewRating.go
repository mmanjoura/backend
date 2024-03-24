package models

import "time"

type ReviewRating struct {
	ID          int       `json:"id"  gorm:"primary_key"`
	Product_ID  int       `json:"product_ID"`
	Customer_ID int       `json:"customer_ID"`
	Rating      int       `json:"rating"`
	Review_Text string    `json:"review_Text"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateReviewRating struct {
	Product_ID  int    `json:"product_ID"`
	Customer_ID int    `json:"customer_ID"`
	Rating      int    `json:"rating" binding:"required"`
	Review_Text string `json:"reviewText" binding:"required"`
}
type UpdateReviewRating struct {
	Product_ID  int    `json:"product_ID"`
	Customer_ID int    `json:"customer_ID"`
	Rating      int    `json:"rating"`
	Review_Text string `json:"reviewText"`
}

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	ProductID int       `json:"productID"`
	Text      string    `json:"text"`
	Rating    int       `json:"rating"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
