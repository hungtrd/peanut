package domain

import "time"

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Bob      time.Time `json:"dob"`
}
