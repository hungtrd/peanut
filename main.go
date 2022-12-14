package main

import (
	"fmt"
	"golang_first_pj/config"
	"golang_first_pj/domain/model"
	"golang_first_pj/infra"
	"gorm.io/gorm"
	"log"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)

	server.Router.Run()
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
