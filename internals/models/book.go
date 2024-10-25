package models

import "time"

type Book struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CategoryId    int       `json:"categoryId"`
	CoverId       int       `json:"coverId"`
	Price         float64   `json:"price"`
	Quantity      int       `json:"quantity"`
	Discount      float64   `json:"discount"`
	AddedAt       time.Time `json:"addedAt"`
	PurchaseCount int       `json:"purchaseCount"`
}

type BookCreateRequest struct {
	Title         string  `json:"title" validate:"required,notBlank"`
	Description   string  `json:"description" validate:"required,notBlank"`
	CategoryId    int     `json:"categoryId" validate:"required,number"`
	CoverId       int     `json:"coverId" validate:"required,number"`
	Price         float64 `json:"price" validate:"required,number,gte=0"`
	Quantity      int     `json:"quantity" validate:"required,number,gte=0"`
	Discount      float64 `json:"discount" validate:"required,number,gte=0"`
}

type BookUpdateRequest struct {
	Title         string  `json:"title" validate:"required,notBlank"`
	Description   string  `json:"description" validate:"required,notBlank"`
	CategoryId    int     `json:"categoryId" validate:"required,number"`
	Price         float64 `json:"price" validate:"required,number,gte=0"`
	Quantity      int     `json:"quantity" validate:"required,number,gte=0"`
	Discount      float64 `json:"discount" validate:"required,number,gte=0"`
}
