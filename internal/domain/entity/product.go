package entity

import (
	"time"
)

type Product struct {
	Uuid        string   `json:"uuid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Photos      []string `json:"photo"`
	Category    string   `json:"category"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

func NewProduct(
	uuid string,
	title string,
	description string,
	price float64,
	photos []string,
	category string,
) *Product {
	entity := &Product{
		Uuid:        uuid,
		Title:       title,
		Description: description,
		Price:       price,
		Photos:      photos,
		Category:    category,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   "",
	}
	return entity
}
