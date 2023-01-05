package infra

import (
	"peanut/domain"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
		&domain.Book{},
	)
}
