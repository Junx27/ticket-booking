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
	scheduleMiddleware := scheduleRepository.(*repository.ScheduleRepository)

	scheduleGroup := r.Group("/schedules")
	scheduleGroup.Use(middleware.AuthProtected(db))
	{
		scheduleGroup.GET("/available", middleware.RoleRequired("admin", "customer"), scheduleHandler.GetManyCustomer)
		scheduleGroup.GET("/", middleware.AccessPermission(scheduleMiddleware), scheduleHandler.GetMany)
		scheduleGroup.GET("/:id", middleware.AccessPermission(scheduleMiddleware), scheduleHandler.GetOne)
		scheduleGroup.POST("/", middleware.RoleRequired("provider"), scheduleHandler.CreateOne)
		scheduleGroup.PUT("/:id", middleware.AccessPermission(scheduleMiddleware), middleware.RoleRequired("provider", "admin"), scheduleHandler.UpdateOne)
		scheduleGroup.DELETE("/:id", middleware.AccessPermission(scheduleMiddleware), middleware.RoleRequired("provider", "admin"), scheduleHandler.DeleteOne)
	}
}
