package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name  string  `json:"name" binding:"required" gorm:"unique"`
	Year  int     `json:"year" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateBookForm struct {
	Name  *string
	Year  *int
	Price *float64
}
