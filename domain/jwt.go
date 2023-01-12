package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"peanut/config"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(config.JwtSecretKey)

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
