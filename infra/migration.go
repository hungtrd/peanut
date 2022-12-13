package infra

import (
	"gorm.io/gorm"
	"peanut/domain"
)

func Migration(db *gorm.DB) {
	if !db.Migrator().HasTable(&domain.User{}) {
		db.Migrator().CreateTable(&domain.User{})
	}
	return
}
