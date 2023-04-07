package config

import (
	"database/sql"
	"fmt"
	ct "ojam-test/c2/s3/app/constant"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDatabaseConnection() *sql.DB {
	var err error

	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		ct.HOST, ct.PORT, ct.USER, ct.PASSWORD, ct.DB_NAME, ct.SSL_MODE,
	)

	db, err = sql.Open(ct.DRIVER, dataSource)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}