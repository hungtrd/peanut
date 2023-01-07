package usecase

import (
	"context"
	"peanut/domain"
	"peanut/repository"
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

func (c contentUsecase) CreateContent(ctx context.Context, content domain.CreateContent, filePath string) error {
	data := domain.Content{
		Name:        content.Name,
		Thumbnail:   filePath,
		Description: content.Description,
		Tag:         content.Tag,
		Category:    content.Category,
		Resolution:  content.Resolution,
		AspectRatio: content.AspectRatio,
		PlayTime:    content.PlayTime,
	}

	err := c.ContentRepo.CreateContent(data)
	if err != nil {
		return err
	}
	return nil
}

func NewContentUsecase(repo repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: repo,
	}
}
