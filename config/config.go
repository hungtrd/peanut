package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	TmpPath       string
	JwtSecretKey  string
	MaxSizeUpload int
	TokenLifespan int
	ApiSecret     string
)

func getConfig() {
	TmpPath = os.Getenv("TMP_PATH")
	JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	ApiSecret = os.Getenv("API_SECRET")

	// Get int
	var err error

	if MaxSizeUpload, err = strconv.Atoi(os.Getenv("MAX_SIZE_UPLOAD")); err != nil {
		panic(fmt.Errorf("get MAX_SIZE_UPLOAD error: %w", err))
	}

	TokenLifespan, err = strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		panic(fmt.Errorf("get MAX_SIZE_UPLOAD error: %w", err))
	}
}
