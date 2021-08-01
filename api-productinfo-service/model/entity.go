package model

import "time"

type ProductInfo struct {
	Name        string   `json:"name"`
	ImagesURL   string   `json:"imageURL"`
	Description string `json:"description"`
	Price       string   `json:"price"`
	Reviews     int      `json:"totalReviews"`
}
type ProductRequest struct {
	Url         string      `json:"url"`
	ProductInfo *ProductInfo `json:"product"`
}
type Product struct {
	Url         string      `json:"url"`
	ProductInfo *ProductInfo `json:"product"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
