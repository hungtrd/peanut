package repository

import (
	"peanut/domain"

	"gorm.io/gorm"
)

type BookRepo interface {
	GetBooks() ([]domain.Book, error)
	GetBookById(id int) (*domain.Book, error)
	CreateBook(b domain.Book) (*domain.Book, error)
	UpdateBook(updatingBook domain.Book, id int) (*domain.Book, error)
	DeleteBook(id int) error
}

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{DB: db}
}

func (r *bookRepo) GetBooks() (books []domain.Book, err error) {
	r.DB.Raw("SELECT * FROM books").Scan(&books)
	return
}

func (r *bookRepo) GetBookById(id int) (book *domain.Book, err error) {
	result := r.DB.First(&book, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

func (r *bookRepo) CreateBook(b domain.Book) (book *domain.Book, err error) {
	book = &domain.Book{Name: b.Name, Price: b.Price, Year: b.Year}
	result := r.DB.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

func (r *bookRepo) UpdateBook(updatingBook domain.Book, id int) (*domain.Book, error) {
	var book domain.Book
	result := r.DB.First(&book, "id =?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	book = updatingBook
	r.DB.Save(&book)
	return &book, nil
}

func (r *bookRepo) DeleteBook(id int) error {
	var book domain.Book
	result := r.DB.First(&book, "id=?", id)
	if result.Error != nil {
		return result.Error
	}
	deleteResult := r.DB.Delete(&book)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}
	return nil
}
