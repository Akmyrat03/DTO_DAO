package main

import (
	"dto/config"
	_ "dto/docs"
	userRoutes "dto/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Management API
// @version 1.0
// @description This is a sample API for managing users
// @host localhost:8000
// @BasePath /api
func main() {
	db := config.ConnectDB()
	defer db.Close()

	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api, db)

	log.Println("Starting the app in release mode on localhost:8000")

	if err := app.Run("localhost:8000"); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}
}
