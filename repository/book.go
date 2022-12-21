package repository

import (
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type BookRepo interface {
	CreateBook(b domain.Book) (*domain.Book, error)
}

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{DB: db}
}

func (r *bookRepo) CreateBook(b domain.Book) (book *domain.Book, err error) {
	book = &domain.Book{
		Title:       b.Title,
		Description: b.Description,
		Author:      b.Author,
	}
	result := r.DB.Create(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.CreateBook] failed: %w", result.Error)
		return nil, err
	}

	return
}
