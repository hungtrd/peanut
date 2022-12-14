package infra

import (
	"gorm.io/gorm"
	"peanut/domain"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
	)
	return
}
