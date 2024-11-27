package routes

import (
	"dto/dao"
	"dto/handler"
	service "dto/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitUserRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	userDAO := dao.NewUserDAO(DB)
	userService := service.NewUserService(userDAO)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	userRoutes.POST("/create", userHandler.CreateUser)
	userRoutes.GET("/all", userHandler.GetAllUsers)
	userRoutes.GET("/:id", userHandler.GetUserByID)
	userRoutes.DELETE("/:id", userHandler.DeleteUserByID)
}
