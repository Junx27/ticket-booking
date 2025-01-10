package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupScheduleRouter(r *gin.Engine, db *gorm.DB) {
	scheduleRepository := repository.NewScheduleRepository(db)
	scheduleHandler := controller.NewScheduleHandler(scheduleRepository)

	scheduleGroup := r.Group("/schedules")
	scheduleGroup.Use(middleware.AuthProtected(db))
	{
		scheduleGroup.GET("/", scheduleHandler.GetMany)
		scheduleGroup.GET("/:id", scheduleHandler.GetOne)
		scheduleGroup.POST("/", middleware.RoleRequired("provider"), scheduleHandler.CreateOne)
		scheduleGroup.PUT("/:id", middleware.RoleRequired("provider"), scheduleHandler.UpdateOne)
		scheduleGroup.DELETE("/:id", middleware.RoleRequired("provider", "admin"), scheduleHandler.DeleteOne)
	}
}
