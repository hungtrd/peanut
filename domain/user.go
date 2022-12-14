package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"peanut/pkg/hash"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" binding:"required,usernameAllow"`
	Email     string `json:"email" binding:"required,email" gorm:"unique"`
	Password  string `json:"password" binding:"required,passwordAllow"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (u *User) HashPassword(password string) {
	u.Password = hash.GenerateFromPassword(password)
}

type RequestLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
