package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProviderRouter(r *gin.Engine, db *gorm.DB) {
	providerRepository := repository.NewProviderRepository(db)
	providerHandler := controller.NewProviderHandler(providerRepository)

	providerGroup := r.Group("/providers")
	{
		providerGroup.GET("/", providerHandler.GetMany)
		providerGroup.GET("/:id", providerHandler.GetOne)
		providerGroup.POST("/", providerHandler.CreateOne)
		providerGroup.PUT("/:id", providerHandler.UpdateOne)
		providerGroup.DELETE("/:id", providerHandler.DeleteOne)
	}
}
