package migration

import (	
	"ojam-test/c3/s4/app/entity"
	
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	user(db)
	socialMedia(db)
	photo(db)
	comment(db)	
}

func user(db *gorm.DB) {
	db.AutoMigrate(
		entity.User{},
	)
}

func socialMedia(db *gorm.DB) {
	db.AutoMigrate(
		entity.SocialMedia{},
	)
}

func photo(db *gorm.DB) {
	db.AutoMigrate(
		entity.Photo{},
	)
}

func comment(db *gorm.DB) {
	db.AutoMigrate(
		entity.Comment{},
	)
}


