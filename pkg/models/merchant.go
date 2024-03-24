package models

import "time"

type Merchant struct {
	ID                int          `json:"id"  gorm:"primary_key"`
	BusinessName      string       `json:"business_name "`
	UserName          string       `json:"user_name"`
	FirstName         string       `json:"firs_tname"`
	LastName          string       `json:"last_name"`
	Email             string       `json:"email"`
	PhoneNumber       string       `json:"phone_number"`
	BirthDate         string       `json:"birth_date"`
	About             string       `json:"about"`
	LocationInfo      LocationInfo `json:"location_info" gorm:"foreignKey:LocationInfoRefer"`
	ChangePass        ChangePass   `json:"change_pass" gorm:"foreignKey:ChangePassRefer"`
	LocationInfoRefer int          `json:"location_info_id"`
	ChangePassRefer   int          `json:"change_pass_id"`
	CreatedAt         time.Time    `json:"created_at"  gorm:"autoCreateTime"`
	UpdatedAt         time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateMerchant struct {
	BusinessName string `json:"business_name "`
	UserName     string `json:"user_name"`
	FirstName    string `json:"firs_tname"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"birth_date"`
	About        string `json:"about"`
}

type UpdateMerchant struct {
	BusinessName string `json:"business_name "`
	UserName     string `json:"user_name"`
	FirstName    string `json:"firs_tname"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	BirthDate    string `json:"birth_date"`
	About        string `json:"about"`
}

type LocationInfo struct {
	ID           int    `json:"id"  gorm:"primary_key"`
	MerchantID   int    `json:"merchant_id" gorm:"foreignKey:M"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type ChangePass struct {
	ID               int    `json:"id"  gorm:"primary_key"`
	MerchantID       int    `json:"merchant_id" gorm:"foreignKey:M"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}

type CreateLocation struct {
	MerchantID   int    `json:"merchant_id" gorm:"foreignKey:M"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type CreatePassword struct {
	MerchantID       int    `json:"merchant_id" gorm:"foreignKey:M"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}
type UpdateLocation struct {
	MerchantID   int    `json:"merchant_id" gorm:"foreignKey:M"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}

type UpdatePassword struct {
	MerchantID       int    `json:"merchant_id" gorm:"foreignKey:M"`
	CurrentPassword  string `json:"current_password"`
	NewPassword      string `json:"new_password"`
	NewPasswordAgain string `json:"new_password_again"`
}
