package user

import "time"

type CreateUserReq struct {
	ID        uint   `json:"id" gorm:"autoIncrement"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
