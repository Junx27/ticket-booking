package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProviderRouter(r *gin.Engine, db *gorm.DB) {
	providerRepository := repository.NewProviderRepository(db)
	activityLogRepository := repository.NewActivityLogRepository(db)
	providerHandler := controller.NewProviderHandler(providerRepository, activityLogRepository)
	providerMiddleware := providerRepository.(*repository.ProviderRepository)

	providerGroup := r.Group("/providers")
	providerGroup.Use(middleware.AuthProtected(db))
	{
		providerGroup.GET("/available", middleware.RoleRequired("admin", "customer"), providerHandler.GetManyCustomer)
		providerGroup.GET("/", middleware.AccessPermission(providerMiddleware), providerHandler.GetMany)
		providerGroup.GET("/:id", middleware.AccessPermission(providerMiddleware), providerHandler.GetOne)
		providerGroup.POST("/", middleware.RoleRequired("provider"), providerHandler.CreateOne)
		providerGroup.PUT("/:id", middleware.AccessPermission(providerMiddleware), middleware.RoleRequired("admin", "provider"), providerHandler.UpdateOne)
		providerGroup.DELETE("/:id", middleware.AccessPermission(providerMiddleware), middleware.RoleRequired("admin", "provider"), providerHandler.DeleteOne)
	}
}
