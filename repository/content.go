package repository

import (
	"peanut/domain"

	"gorm.io/gorm"
)

type ContentRepo interface {
	GetContents() (content []domain.Content, error error)
	CreateContent(u domain.Content) error
}

type contentRepo struct {
	DB *gorm.DB
}

func (c contentRepo) GetContents() (content []domain.Content, error error) {
	result := c.DB.Find(&content)

	return content, result.Error
}

func (c contentRepo) CreateContent(content domain.Content) error {
	data := &domain.Content{
		Name:        content.Name,
		Thumbnail:   content.Thumbnail,
		Description: content.Description,
		Tag:         content.Tag,
		Category:    content.Category,
		Resolution:  content.Resolution,
		AspectRatio: content.AspectRatio,
		PlayTime:    content.PlayTime,
	}
	result := c.DB.Create(data)
	return result.Error
}

func NewContentRepo(db *gorm.DB) ContentRepo {
	return &contentRepo{DB: db}
}
