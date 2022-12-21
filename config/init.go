package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const pjDirName = "peanut"

func init() {
	loadEnv()
}

func Setup() {
	setEnv()
	setGinMode()
	getConfig()
}

func loadEnv() {
	re := regexp.MustCompile(`^(.*` + pjDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("Error loading .env file: %w", err)
	}
}
