package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

func init() {
	loadEnv()
}

func Setup() {
	fmt.Println("Initial configuration")
}

var projectDirName = "peanut"

func loadEnv() {
	projectName := regexp.MustCompile("^(.*" + projectDirName + ")")
	currentWork, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWork))
	err := godotenv.Load(string(rootPath) + "/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
