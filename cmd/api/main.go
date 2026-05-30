package main

import (
	"encurtador-url-go/internal/database"
	"encurtador-url-go/internal/domain"
	"encurtador-url-go/internal/routes"
)

func main() {
	database.ConnectionRedis()
	database.ConnectionDatabase()
	database.DB.AutoMigrate(&domain.URL{})

	router := routes.ConfigRouter()

	router.Run(":8080")
}
