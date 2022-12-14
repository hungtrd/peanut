package repository

import (
	"context"
	"fmt"
	"golang_first_pj/domain/model"
	user2 "golang_first_pj/domain/request/user"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByGmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, u user2.CreateUserReq) (*model.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

func (r *userRepo) GetUsers(ctx context.Context) (users []model.User, err error) {
	return
}

func (r *userRepo) GetUser(ctx context.Context, id int) (user *model.User, err error) {
	return
}
func (r *userRepo) GetUserByGmail(ctx context.Context, email string) (user *model.User, err error) {
	user = &model.User{}

	result := r.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		err = fmt.Errorf("[repo.auth.Login] failed: %w", result.Error)
		return nil, err
	}
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u user2.CreateUserReq) (user *model.User, err error) {
	user = &model.User{
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
