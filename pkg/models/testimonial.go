package models

import "time"

type Testimonial struct {
	ID             int       `json:"id"  gorm:"primary_key"`
	Hotel_ID       int       `json:"hotel_id"`
	Meta           string    `json:"meta"`
	Avatar         string    `json:"avatar"`
	Name           string    `json:"name"`
	Designation    string    `json:"designation"`
	Text           string    `json:"text"`
	DelayAnimation int       `json:"delay_animation"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateTestimonial struct {
	Hotel_ID       int    `json:"hotel_id"`
	Meta           string `json:"meta"`
	Avatar         string `json:"avatar"`
	Name           string `json:"name"`
	Designation    string `json:"designation"`
	Text           string `json:"text"`
	DelayAnimation int    `json:"delay_animation"`
}

type UpdateTestimonial struct {
	Hotel_ID       int    `json:"hotel_id"`
	Meta           string `json:"meta"`
	Avatar         string `json:"avatar"`
	Name           string `json:"name"`
	Designation    string `json:"designation"`
	Text           string `json:"text"`
	DelayAnimation int    `json:"delay_animation"`
}
