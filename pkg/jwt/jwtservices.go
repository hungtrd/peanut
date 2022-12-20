package jwtservices

import (
	"fmt"
	"os"
	"peanut/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user *domain.User) (string, error) {
	// generate a jwt token string
	mapClaims := jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(), //30days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	hmac_secret_signing_key := os.Getenv("HMAC_SECRET_SIGNING_KEY")
	// sign token
	tokenString, err := token.SignedString([]byte(hmac_secret_signing_key))
	if err != nil {
		err = fmt.Errorf("Incorrect email or password")
		return "", err
	}
	return tokenString, err
}
