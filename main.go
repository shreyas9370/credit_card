package main

import (
	"credit-card-notification/config"
	"credit-card-notification/controller"
	"credit-card-notification/repository"
	"credit-card-notification/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.POST("/register", userController.RegisterUser)

	router.Run(":8080")
}
