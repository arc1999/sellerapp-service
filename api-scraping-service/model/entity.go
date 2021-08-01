package model

import "time"

type ProductDetailsRequest struct {
	Name        string   `json:"name"`
	ImagesURL   string   `json:"imageURL"`
	Description string `json:"description"`
	Price       string   `json:"price"`
	Reviews     int      `json:"totalReviews"`
}
type ScrapingRequest struct {
	Url string `json:"url"`
}
type ProductInfoRequest struct {
	Url         string      `json:"url"`
	ProductInfo *ProductDetailsRequest `json:"product"`
}

type Product struct {
	Url         string      `json:"url"`
	ProductInfo *ProductDetailsRequest `json:"product"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
