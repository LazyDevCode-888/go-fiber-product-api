package models 

import "time"

// struct user
type User struct {

	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Age int `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}