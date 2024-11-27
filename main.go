package main

import (
	"dto/config"
	userRoutes "dto/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	app := gin.Default()
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api, db)

	log.Println("Starting the app in release mode on localhost:8000")

	if err := app.Run("localhost:8000"); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}
}
