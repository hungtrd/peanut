package middleware

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateFunction() {
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
