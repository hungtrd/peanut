package domain

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Content struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:30"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description" gorm:"size:500"`
	PlayTime    time.Time `json:"play_time"`
	Resolution  string    `json:"resolution" gorm:"size:500"`
	AspectRatio string    `json:"aspect_ratio" gorm:"size:500"`
	Tag         string    `json:"tag" gorm:"size:500"`
	Category    string    `json:"category" gorm:"size:500"`
} //@name Content

type CreateContent struct {
	Name        string                `form:"name" binding:"required,max=30" gorm:"size:30"`
	Thumbnail   *multipart.FileHeader `form:"thumbnail" binding:"required"`
	Description string                `form:"description" gorm:"size:500"`
	PlayTime    string                `form:"play_time" gorm:"size:500" binding:"datetime=2006-01-02 15:04:05"`
	Resolution  string                `form:"resolution" gorm:"size:500"`
	AspectRatio string                `form:"aspect_ratio" gorm:"size:500"`
	Tag         string                `form:"tag" gorm:"size:500"`
	Category    string                `form:"category" gorm:"size:500"`
} //@name CreateContent

type ListContent struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"column:username"`
	Title     string `json:"title" binding:"required" gorm:"size:30"`
	Content   string `json:"content" binding:"required" gorm:"size:500"`
	CreatedAt time.Time
} //@name ListContent
