package main

import (
	"crowdfunding/config"
	"crowdfunding/databases"
	"crowdfunding/handler"
	"crowdfunding/repository"
	"crowdfunding/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	db := databases.DBconnection(cfg)

	if db == nil {
		log.Fatal("Failed to connect DB")
	}

	userRepository := repository.UserRepository(db)
	userService := services.ServicesUser(userRepository)
	userHandler := handler.HandlerUser(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
	}

	router.Run()
}
