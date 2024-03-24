package models

import "time"

type Flight struct {
	ID         int          `json:"id"  gorm:"primary_key"`
	Price      int          `json:"price"`
	Deals      string       `json:"deals"`
	Animation  string       `json:"animation"`
	SelectId   string       `json:"select_Id"`
	FlightList []FlightList `gorm:"foreignKey:flight_id"`
	CreatedAt  time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
type CreateFlight struct {
	Price      int          `json:"price"`
	Deals      string       `json:"deals"`
	Animation  string       `json:"animation"`
	SelectId   string       `json:"select_Id"`
	FlightList []FlightList `gorm:"foreignKey:flight_id"`
}

type UpdateFlight struct {
	Price      int          `json:"price"`
	Deals      string       `json:"deals"`
	Animation  string       `json:"animation"`
	SelectId   string       `json:"select_Id"`
	FlightList []FlightList `gorm:"foreignKey:flight_id"`
}

type FlightList struct {
	ID               int    `json:"id"  gorm:"primary_key"`
	Flight_ID        int    `json:"flight_id" gorm:"foreignkey:flight_id"`
	Avatar           string `json:"avatar"`
	ArrivalAirport   string `json:"arrival_airport"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalTime      string `json:"arrival_time"`
	Duration         string `json:"duration"`
}
