package repository

import (
	"errors"
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers() ([]domain.User, error)
	GetUser(id int) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(u domain.User) (*domain.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) GetUsers() (users []domain.User, err error) {
	return
}

func (r *userRepo) GetUser(id int) (user *domain.User, err error) {
	return
}

func (r *userRepo) GetUserByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	result := r.DB.Where("username = ?", username).First(user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		err := fmt.Errorf("[repo.User.GetUserByUsername] failed: %w", result.Error)
		return nil, err
	}

	return user, nil
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
