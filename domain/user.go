package domain

import (
	"peanut/pkg/bcrypt"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  string `json:"password" binding:"required,passwordAllow"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword(password string) {
	u.Password = bcrypt.GenerateFromPassword(password)
}
