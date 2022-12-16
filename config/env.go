package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	local = iota
	test
	development
	staging
	production
)

var env = local

func setEnv() {
	envName := os.Getenv("PJ_ENV")
	switch strings.ToLower(envName) {
	case "local":
		env = local
	case "test":
		env = test
	case "development":
		env = development
	case "staging":
		env = staging
	case "production":
		env = production
	default:
		panic("PJ_ENV unknown: " + envName + " (available env: local development staging production test)")
	}
	fmt.Println("Running in [" + envName + "] environment")
}

func setGinMode() {
	switch env {
	case production:
		gin.SetMode(gin.ReleaseMode)
	case test:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func IsDevelopment() bool {
	return env == development || env == local
}

func IsTest() bool {
	return env == test
}

func IsProduction() bool {
	return env == production
}
