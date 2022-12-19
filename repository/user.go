package repository

import (
	"context"
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUser(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, u domain.User) (*domain.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) GetUsers(ctx context.Context) (users []domain.User, err error) {
	result := r.DB.Find(&users)

	return users, result.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	result := r.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *userRepo) GetUser(ctx context.Context, id int) (user *domain.User, err error) {
	if err = r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return
}

func (r *userRepo) CreateUser(ctx context.Context, u domain.User) (user *domain.User, err error) {
	user = &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
	result := r.DB.Create(user)

	if result.Error != nil {
		err = fmt.Errorf("[repo.User.CreateUser] failed: %w", result.Error)
		return nil, err
	}

	return
}
