package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupActivityLogRouter(r *gin.Engine, db *gorm.DB) {
	activityLogRepository := repository.NewActivityLogRepository(db)
	activityLogHandler := controller.NewActivityLogHandler(activityLogRepository)

	activityLogGroup := r.Group("/activity-logs")
	activityLogGroup.Use(middleware.AuthProtected(db))
	{
		activityLogGroup.GET("/", middleware.RoleRequired("admin"), activityLogHandler.GetMany)
		activityLogGroup.GET("/activity-log/:user_id", middleware.RoleRequired("admin"), activityLogHandler.GetManyByUser)
		activityLogGroup.GET("/:id", middleware.RoleRequired("admin"), activityLogHandler.GetOne)
		activityLogGroup.DELETE("/:id", middleware.RoleRequired("admin"), activityLogHandler.DeleteOne)
		activityLogGroup.DELETE("/", middleware.RoleRequired("admin"), activityLogHandler.DeleteMany)
	}
}
