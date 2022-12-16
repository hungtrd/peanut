package domain

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title" binding:"required" gorm:"size:30"`
	Content string `json:"content" binding:"required" gorm:"size:500"`
}

type ListTodo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"column:username"`
	Title     string `json:"title" binding:"required" gorm:"size:30"`
	Content   string `json:"content" binding:"required" gorm:"size:500"`
	CreatedAt time.Time
}
