package main

import (
	"github.com/Junx27/ticket-booking/config"
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/database"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewEnvConfig()

	db := database.Init(cfg, database.DBMigrator)
	database.SeedUsers(db)

	userRepository := repository.NewUserRepository(db)

	userHandler := controller.NewUserHandler(userRepository)

	r := gin.Default()

	r.GET("/users", userHandler.GetMany)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "system ok",
		})
	})

	r.Run(":8080")
}
