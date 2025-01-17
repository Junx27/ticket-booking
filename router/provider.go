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
		providerGroup.GET("/", middleware.RoleRequired("admin", "customer"), providerHandler.GetMany)
		providerGroup.GET("/all/:user_id", middleware.AccessPermissionView(providerMiddleware), providerHandler.GetManyByUser)
		providerGroup.GET("/:id", providerHandler.GetOne)
		providerGroup.POST("/", middleware.RoleRequired("provider"), providerHandler.CreateOne)
		providerGroup.PUT("/:id", middleware.AccessPermissionAction(providerMiddleware), middleware.RoleRequired("admin", "provider"), providerHandler.UpdateOne)
		providerGroup.DELETE("/:id", middleware.AccessPermissionAction(providerMiddleware), middleware.RoleRequired("admin", "provider"), providerHandler.DeleteOne)
	}
}
