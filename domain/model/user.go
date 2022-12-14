package model

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
