package repository

import "gorm.io/gorm"

type UserRepo interface {
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}
