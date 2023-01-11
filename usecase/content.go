package usecase

import (
	"context"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/repository"
	"time"
)

type ContentUsecase interface {
	GetContents(ctx context.Context) ([]domain.Content, error)
	CreateContent(ctx context.Context, content domain.CreateContent, filepath string) error
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func (c contentUsecase) GetContents(ctx context.Context) (contents []domain.Content, err error) {
	contents, err = c.ContentRepo.GetContents()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (c contentUsecase) CreateContent(ctx context.Context, content domain.CreateContent, filePath string) (err error) {
	data := domain.Content{
		Name:        content.Name,
		Thumbnail:   filePath,
		Description: content.Description,
		Tag:         content.Tag,
		Category:    content.Category,
		Resolution:  content.Resolution,
		AspectRatio: content.AspectRatio,
	}
	data.PlayTime, err = time.Parse("2006-01-02 15:04:05", content.PlayTime)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	err = c.ContentRepo.CreateContent(data)
	if err != nil {
		return
	}
	return nil
}

func NewContentUsecase(repo repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: repo,
	}
}
