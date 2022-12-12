package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

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
