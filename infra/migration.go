package infra

import (
	"peanut/domain"

	"gorm.io/gorm"
)

func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Book{})
	return
}
