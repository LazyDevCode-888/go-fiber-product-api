
package models 

import "time"

type Product struct {

	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	Category string `json:"category"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
}

