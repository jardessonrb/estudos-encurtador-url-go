package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	uriConnection := "host=localhost user=postgres password=admin dbname=encurtador_url_db port=5432 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(uriConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection
}
