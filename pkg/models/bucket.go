package models

type Bucket struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProjectID  string `json:"project_id"`
	FolderName string `json:"folder_name"`
}

type ProductImage struct {
	ID         int    `json:"id"`
	RefferId   string `json:"reffer_id"`
	CategoryId string `json:"category_id"`
}
