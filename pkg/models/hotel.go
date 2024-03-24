package models

import (
	"time"
)

type Hotel struct {
	ID            int            `json:"id"  gorm:"primary_key"`
	Category_ID   int            `json:"category_id"`
	Tag           string         `json:"tag"`
	Img           string         `json:"img"`
	Title         string         `json:"title"`
	Location      string         `json:"location"`
	Ratings       float32        `json:"ratings"`
	Reviews       int            `json:"reviews"`
	Price         float32        `json:"price"`
	Animation     int            `json:"animation"`
	City          string         `json:"city"`
	HotelInfo     HotelInfo      `json:"hotel_info" gorm:"foreignKey:HotelID"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:HotelID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateHotel struct {
	Category_ID   int            `json:"category_id"`
	Tag           string         `json:"tag"`
	Img           string         `json:"img"`
	Title         string         `json:"title"`
	Location      string         `json:"location"`
	Ratings       float32        `json:"ratings"`
	Reviews       int            `json:"reviews"`
	Price         float32        `json:"price"`
	Animation     int            `json:"animation"`
	City          string         `json:"city"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:HotelID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
}

type UpdateHotel struct {
	ID            int            `json:"id"  gorm:"primary_key"`
	Category_ID   int            `json:"category_id"`
	Tag           string         `json:"tag"`
	Img           string         `json:"img"`
	Title         string         `json:"title"`
	Location      string         `json:"location"`
	Ratings       float32        `json:"ratings"`
	Reviews       int            `json:"reviews"`
	Price         float32        `json:"price"`
	Animation     int            `json:"animation"`
	City          string         `json:"city"`
	SlideImages   []SlideImage   `json:"SlideImages" gorm:"foreignKey:HotelID"`
	GalleryImages []GalleryImage `json:"GalleryImages" gorm:"foreignKey:HotelID"`
	SlideImg      []string       `json:"slide_img"`
	GalleryImg    []string       `json:"gallery_img"`
}

// HotelInfo represents the Hotel_Infos table
type HotelInfo struct {
	ID              int             `json:"id"`
	HotelID         int             `json:"hotel_id"`
	RoomType        string          `json:"room_type"`
	Overview        string          `json:"overview"`
	Img             string          `json:"img"`
	Price           string          `json:"price"`
	HotelFacilities []HotelFacility `json:"hotel_facility" gorm:"foreignKey:HotelID"`
	HotelBenefits   []HotelBenefit  `json:"hotel_benefit" gorm:"foreignKey:HotelID"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// HotelFacility represents the Hotel_Facilities table
type HotelFacility struct {
	ID           int    `json:"id"`
	HotelID      int    `json:"hotel_id"`
	ClassName    string `json:"class_name"`
	FacilityName string `json:"facility_name"`
	Exist        int    `json:"exist"`
	IsHighlight  int    `json:"is_highlight"`
}

// HotelBenefit represents the Hotel_Benefits table
type HotelBenefit struct {
	ID          int    `json:"id"`
	HotelID     int    `json:"hotel_id"`
	ClassName   string `json:"class_name"`
	BenefitName string `json:"benefit_name"`
	Exist       int    `json:"exist"`
}
