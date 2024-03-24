package models

import "time"

type Destination struct {
	ID         int       `json:"id"  gorm:"primary_key"`
	Class      string    `json:"class"`
	Title      string    `json:"title"`
	Location   string    `json:"location"`
	Travellers string    `json:"travellers"`
	Hover      string    `json:"hover"`
	Img        string    `json:"img"`
	City       string    `json:"city"`
	Properties string    `json:"properties"`
	Region     string    `json:"region"`
	Animation  string    `json:"animation"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateDestination struct {
	Class      string `json:"class"`
	Title      string `json:"title"`
	Location   string `json:"location"`
	Travellers string `json:"travellers"`
	Hover      string `json:"hover"`
	Img        string `json:"img"`
	City       string `json:"city"`
	Properties string `json:"properties"`
	Region     string `json:"region"`
	Animation  string `json:"animation"`
	Name       string `json:"name"`
}

type UpdateDestination struct {
	Class      string `json:"class"`
	Title      string `json:"title"`
	Location   string `json:"location"`
	Travellers string `json:"travellers"`
	Hover      string `json:"hover"`
	Img        string `json:"img"`
	City       string `json:"city"`
	Properties string `json:"properties"`
	Region     string `json:"region"`
	Animation  string `json:"animation"`
	Name       string `json:"name"`
}
