package models

import "time"

type Blog struct {
	ID             int       `json:"id"  gorm:"primary_key"`
	Title          string    `json:"title"`
	Img            string    `json:"img"`
	Details        string    `json:"details"`
	Tag            string    `json:"tag"`
	Tags           []string  `json:"tags"`
	DelayAnimation int       `json:"delay_animation"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateBlog struct {
	Title          string   `json:"title"`
	Img            string   `json:"img"`
	Details        string   `json:"details"`
	Tag            string   `json:"tag"`
	Tags           []string `json:"tags"`
	DelayAnimation int      `json:"delay_animation"`
}

type UpdateBlog struct {
	Title          string   `json:"title"`
	Img            string   `json:"img"`
	Details        string   `json:"details"`
	Tag            string   `json:"tag"`
	Tags           []string `json:"tags"`
	DelayAnimation int      `json:"delay_animation"`
}
