package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       float64 `json:"price"`
	Description string `json:"description"`
	Quantity    int `json:"quantity"`
}
