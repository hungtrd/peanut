package repository

import (
	"context"
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (*domain.User, error)
	CreateUser(u domain.User) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return
}

func (r *userRepo) GetUserById(ctx context.Context, id int) (user *domain.User, err error) {
	result := r.DB.First(&user, "id=?", id)

	if result.Error != nil {
		err = fmt.Errorf("[repo.User.First] failed: %w", result.Error)
		return nil, err
	}
	return
}

func (r *userRepo) GetUserByUsername(username string) (user *domain.User, err error) {
	result := r.DB.First(&user, "username=?", username)
	if result.Error != nil {
		err = fmt.Errorf("[repo.User.First] failed: %w", result.Error)
		return nil, err
	}
	fmt.Println(user)
	if user.ID == 0 {
		err = fmt.Errorf("Incorrect email or password")
		return nil, err
	}
	return
}

func (r *userRepo) CreateUser(u domain.User) (user *domain.User, err error) {
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
