package usecase

import (
	"context"
	"peanut/domain"
	"peanut/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, u domain.User) error
}

type userUsecase struct {
	UserRepo repository.UserRepo
}

func NewUserUsecase(r repository.UserRepo) UserUsecase {
	return &userUsecase{
		UserRepo: r,
	}
}

func (uc *userUsecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return
}

func (uc *userUsecase) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	return
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.User) (err error) {
	// TODO: hash password
	_, err = uc.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return
}
