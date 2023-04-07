package config

import (
	"fmt"
	"log"
	ct "ojam-test/c2/s4/app/constant"
	"ojam-test/c2/s4/app/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabaseConnection() *gorm.DB {
	var err error

	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		ct.HOST, ct.PORT, ct.USER, ct.PASSWORD, ct.DB_NAME, ct.SSL_MODE, ct.TIMEZONE,
	)

	db, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migration.DatabaseMigration(db)

	return db
}