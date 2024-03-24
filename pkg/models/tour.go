package models

import "time"

type Tour struct {
	ID                    int            `json:"id"`
	UserID                string         `json:"user_id"`
	Tag                   string         `json:"tag"`
	Title                 string         `json:"title"`
	NumberOfReviews       string         `json:"number_of_reviews"`
	ReviewsComment        string         `json:"reviews_comment"`
	Location              string         `json:"location"`
	Latitude              string         `json:"latitude"`
	Longitude             string         `json:"longitude"`
	MapUrl                string         `json:"map_url"`
	MinimumDuration       string         `json:"minimum_duration"`
	GroupSize             string         `json:"group_size"`
	Overview              string         `json:"overview"`
	CancellationPolicy    string         `json:"cancellation_policy"`
	WhatsIncluded         string         `json:"whats_included"`
	Highlights            string         `json:"highlights"`
	AdditionalInformation string         `json:"additional_information"`
	ImportantInformation  string         `json:"important_information"`
	Price                 string         `json:"price"`
	TourType              string         `json:"tour_type"`
	Animation             string         `json:"animation"`
	Images                []Image        `json:"images"`
	GalleryImages         []GalleryImage `json:"gallery_images"`
	SlideImages           []SlideImage   `json:"slide_images"`
	Itineraries           []Itinerary    `json:"itineraries"`
	Faqs                  []Faq          `json:"faqs"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

type CreateTour struct {
	ID                    int       `json:"id"`
	UserID                string    `json:"user_id"`
	Tag                   string    `json:"tag"`
	Title                 string    `json:"title"`
	NumberOfReviews       string    `json:"number_of_reviews"`
	ReviewsComment        string    `json:"reviews_comment"`
	Location              string    `json:"location"`
	Latitude              string    `json:"latitude"`
	Longitude             string    `json:"longitude"`
	MapUrl                string    `json:"map_url"`
	MinimumDuration       string    `json:"minimum_duration"`
	GroupSize             string    `json:"group_size"`
	Overview              string    `json:"overview"`
	CancellationPolicy    string    `json:"cancellation_policy"`
	WhatsIncluded         string    `json:"whats_included"`
	Highlights            string    `json:"highlights"`
	AdditionalInformation string    `json:"additional_information"`
	ImportantInformation  string    `json:"important_information"`
	Price                 string    `json:"price"`
	TourType              string    `json:"tour_type"`
	Animation             string    `json:"animation"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type UpdateTour struct {
	ID                    int       `json:"id"`
	UserID                string    `json:"user_id"`
	Tag                   string    `json:"tag"`
	Title                 string    `json:"title"`
	NumberOfReviews       string    `json:"number_of_reviews"`
	ReviewsComment        string    `json:"reviews_comment"`
	Location              string    `json:"location"`
	Latitude              string    `json:"latitude"`
	Longitude             string    `json:"longitude"`
	MapUrl                string    `json:"map_url"`
	MinimumDuration       string    `json:"minimum_duration"`
	GroupSize             string    `json:"group_size"`
	Overview              string    `json:"overview"`
	CancellationPolicy    string    `json:"cancellation_policy"`
	WhatsIncluded         string    `json:"whats_included"`
	Highlights            string    `json:"highlights"`
	AdditionalInformation string    `json:"additional_information"`
	ImportantInformation  string    `json:"important_information"`
	Price                 string    `json:"price"`
	TourType              string    `json:"tour_type"`
	Animation             string    `json:"animation"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type TourFilter struct {
	// Filtering fields.
	ID   int    `json:"id"`
	Type string `json:"type"`

	// Restrict to subset of results.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
