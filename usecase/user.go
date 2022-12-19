package usecase

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
	"peanut/pkg/hash"
	"peanut/pkg/jwt"
	"peanut/repository"
)

type UserUsecase interface {
	Login(ctx context.Context, req domain.RequestLogin) (string, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, u domain.User) error
}

type userUsecase struct {
	UserRepo repository.UserRepo
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		UserRepo: repository.NewUserRepo(db),
	}
}

func (uc *userUsecase) Login(ctx context.Context, req domain.RequestLogin) (token string, err error) {
	user, err := uc.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("login fail")
	}
	result := hash.CompareHashAndPassword(user.Password, req.Password)
	if !result {
		return "", fmt.Errorf("login fail")
	}
	return jwt.GenerateToken(user.ID)
}

func (uc *userUsecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	users, err = uc.UserRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *userUsecase) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	user, err = uc.UserRepo.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.User) (err error) {
	u.HashPassword(u.Password)
	_, err = uc.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return
}
