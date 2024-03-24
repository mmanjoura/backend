package models

import "time"

type AddBanner struct {
	ID             int       `json:"id"`
	Img            string    `json:"img"`
	Title          string    `json:"title"`
	Meta           string    `json:"meta"`
	RouterPath     string    `json:"routerPath"`
	DelayAnimation string    `json:"delayAnimation"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateAddBanner struct {
	Img            string `json:"img"`
	Title          string `json:"title"`
	Meta           string `json:"meta"`
	RouterPath     string `json:"routerPath"`
	DelayAnimation string `json:"delayAnimation"`
}

type UpdateAddBanner struct {
	ID             int    `json:"id"`
	Img            string `json:"img"`
	Title          string `json:"title"`
	Meta           string `json:"meta"`
	RouterPath     string `json:"routerPath"`
	DelayAnimation string `json:"delayAnimation"`
}
