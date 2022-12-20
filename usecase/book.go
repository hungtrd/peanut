package usecase

import (
	"peanut/domain"
	"peanut/repository"

	"gorm.io/gorm"
)

type BookUsecase interface {
	GetBooks() ([]domain.Book, error)
	GetBook(id int) (*domain.Book, error)
	CreateBook(b domain.Book) (*domain.Book, error)
	UpdateBook(updateForm domain.UpdateBookForm, ID int) (*domain.Book, error)
	DeleteBook(id int) error
}

type bookUsecase struct {
	BookRepo repository.BookRepo
}

func NewBookUsecase(db *gorm.DB) BookUsecase {
	return &bookUsecase{
		BookRepo: repository.NewBookRepo(db),
	}
}

func (uc *bookUsecase) GetBooks() (books []domain.Book, err error) {
	books, err = uc.BookRepo.GetBooks()
	if err != nil {
		return nil, err
	}
	return
}
func (uc *bookUsecase) GetBook(id int) (book *domain.Book, err error) {
	book, err = uc.BookRepo.GetBookById(id)
	if err != nil {
		return nil, err
	}
	return
}
func (uc *bookUsecase) CreateBook(b domain.Book) (book *domain.Book, err error) {
	book, err = uc.BookRepo.CreateBook(b)
	if err != nil {
		return nil, err
	}
	return
}
func (uc *bookUsecase) UpdateBook(updateForm domain.UpdateBookForm, ID int) (book *domain.Book, err error) {
	updatingBook, err := uc.BookRepo.GetBookById(ID)
	if err != nil {
		return nil, err
	}
	if updateForm.Name != nil {
		updatingBook.Name = *updateForm.Name
	}
	if updateForm.Year != nil {
		updatingBook.Year = *updateForm.Year
	}
	if updateForm.Price != nil {
		updatingBook.Price = *updateForm.Price
	}
	book, err = uc.BookRepo.UpdateBook(*updatingBook, ID)
	if err != nil {
		return nil, err
	}
	return
}
func (uc *bookUsecase) DeleteBook(id int) (err error) {
	err = uc.BookRepo.DeleteBook(id)
	if err != nil {
		return
	}
	return
}
