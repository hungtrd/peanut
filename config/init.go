package config

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"regexp"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Setup() {
	fmt.Println("Initial configuration")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordAllow", passwordAllow)
	}
}

var passwordAllow validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	pattern := "[\\w]{8,}"
	re, _ := regexp.MatchString(pattern, password)
	return re
}
