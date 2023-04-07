package migration

import (	
	"ojam-test/c2/s4/app/entity"
	
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	book(db)
}

func book(db *gorm.DB) {
	db.AutoMigrate(
		entity.BookCategory{},
		entity.Book{},
	)
}