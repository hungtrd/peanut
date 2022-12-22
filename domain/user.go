package domain

import (
	"gorm.io/gorm"
	"peanut/pkg/hash"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,usernameAllow"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required,passwordAllow"`
	Todos    []Todo
}

func (u *User) HashPassword(password string) {
	u.Password = hash.GenerateFromPassword(password)
}

type RequestLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
