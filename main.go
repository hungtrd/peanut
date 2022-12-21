package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"
	"peanut/pkg/i18n"

	"gorm.io/gorm"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/api

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	i18n.SetupI18n()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)
	infra.Migration(dbClient)

	server.Router.Run(":8080")
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
