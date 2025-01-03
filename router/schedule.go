package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupScheduleRouter(r *gin.Engine, db *gorm.DB) {
	scheduleRepository := repository.NewScheduleRepository(db)
	scheduleHandler := controller.NewScheduleHandler(scheduleRepository)

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.GET("/", scheduleHandler.GetMany)
		scheduleGroup.GET("/:id", scheduleHandler.GetOne)
		scheduleGroup.POST("/", scheduleHandler.CreateOne)
		scheduleGroup.PUT("/:id", scheduleHandler.UpdateOne)
		scheduleGroup.DELETE("/:id", scheduleHandler.DeleteOne)
		scheduleGroup.PUT("/:id/seats", scheduleHandler.UpdateSeatsStatus)
	}
}
