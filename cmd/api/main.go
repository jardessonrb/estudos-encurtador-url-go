package main

import (
	"encurtador-url-go/internal/database"
	"encurtador-url-go/internal/domain"
	"encurtador-url-go/internal/routes"
	"os"
)

func main() {
	database.ConnectionRedis()
	database.ConnectionDatabase()
	database.DB.AutoMigrate(&domain.URL{})

	router := routes.ConfigRouter()

	porta := os.Getenv("PORTA")

	if porta == "" {
		porta = "8080"
	}

	router.Run(":" + porta)
}
