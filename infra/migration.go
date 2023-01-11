package infra

import (
	"gorm.io/gorm"
	"log"
	"peanut/domain"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Todo{},
		&domain.Content{},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
