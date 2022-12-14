package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"
	"peanut/pkg/i18n"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	i18n.SetupI18n()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)

	server.Router.Run(":8080")
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
