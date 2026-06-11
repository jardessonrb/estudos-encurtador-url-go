package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	DB_PORT := os.Getenv("DB_HOST")

	uriConnection := fmt.Sprintf("host=%s user=postgres password=admin dbname=encurtador_url_db port=5432 sslmode=disable", DB_PORT)

	connection, err := gorm.Open(postgres.Open(uriConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection
}
