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
	user := uc.UserRepo.GetUserByEmail(ctx, req.Email)
	result := hash.CompareHashAndPassword(user.Password, req.Password)
	if !result {
		return "", fmt.Errorf("login fail")
	}
	return jwt.GenerateToken(user.ID)
}

func (uc *userUsecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return
}

func (uc *userUsecase) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	return
}

func (uc *userUsecase) CreateUser(ctx context.Context, u domain.User) (err error) {
	u.HashPassword(u.Password)
	_, err = uc.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return
}
