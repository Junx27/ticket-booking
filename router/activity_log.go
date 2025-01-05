package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupActivityLogRouter(r *gin.Engine, db *gorm.DB) {
	activityLogRepository := repository.NewActivityLogRepository(db)
	activityLogHandler := controller.NewActivityLogHandler(activityLogRepository)

	activityLogGroup := r.Group("/activity-logs")
	{
		activityLogGroup.GET("/", activityLogHandler.GetMany)
		activityLogGroup.GET("/activity-log/:user_id", activityLogHandler.GetManyByUser)
		activityLogGroup.GET("/:id", activityLogHandler.GetOne)
		activityLogGroup.POST("/", activityLogHandler.CreateOne)
		activityLogGroup.DELETE("/:id", activityLogHandler.DeleteOne)
		activityLogGroup.DELETE("/", activityLogHandler.DeleteMany)
	}
}
