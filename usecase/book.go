package usecase

import (
	"context"
	"fmt"
	"peanut/domain"
	"peanut/repository"
)

type BookUsecase interface {
	CreateBook(ctx context.Context, req domain.CreateBookReq) (domain.CreateBookResp, error)
}

type bookUsecase struct {
	BookRepo repository.BookRepo
}

func NewBookUsecase(r repository.BookRepo) BookUsecase {
	return &bookUsecase{
		BookRepo: r,
	}
}

func (uc *bookUsecase) CreateBook(ctx context.Context, req domain.CreateBookReq) (resp domain.CreateBookResp, err error) {
	b := domain.Book{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
	}

	book, err := uc.BookRepo.CreateBook(b)
	if err != nil {
		err = fmt.Errorf("[usecase.bookUsecase.CreateBook] failed: %w", err)
		return
	}

	resp = domain.CreateBookResp{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}

	return
}
