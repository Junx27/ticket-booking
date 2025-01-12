package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupNotificationRouter(r *gin.Engine, db *gorm.DB) {
	notificationRepository := repository.NewNotificationRepository(db)
	notificationHandler := controller.NewNotificationHandler(notificationRepository)

	notificationGroup := r.Group("/notifications")
	notificationGroup.Use(middleware.AuthProtected(db))
	{
		notificationGroup.GET("/", middleware.RoleRequired("admin"), notificationHandler.GetMany)
		notificationGroup.GET("/user/:user_id", notificationHandler.GetManyByUser)
		notificationGroup.GET("/:id", notificationHandler.GetOne)
		notificationGroup.PUT("/:id", middleware.RoleRequired("admin"), notificationHandler.UpdateOne)
		notificationGroup.DELETE("/:id", notificationHandler.DeleteOne)
		notificationGroup.DELETE("/user/:user_id", notificationHandler.DeleteAllByUser)
	}
}
