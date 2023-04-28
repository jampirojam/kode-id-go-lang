package migration

import (	
	"ojam-test/c3/s2/app/entity"
	
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	product(db)
	user(db)
}

func product(db *gorm.DB) {
	db.AutoMigrate(
		entity.ProductType{},
		entity.Product{},
	)
}

func user(db *gorm.DB) {
	db.AutoMigrate(
		entity.UserRole{},
		entity.User{},
	)
}