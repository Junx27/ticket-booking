package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTicketUsageRouter(r *gin.Engine, db *gorm.DB) {
	ticketUsageRepository := repository.NewTicketUsageRepository(db)
	ticketUsageHandler := controller.NewTicketUsageHandler(ticketUsageRepository)
	ticketMiddleware := ticketUsageRepository.(*repository.TicketUsageRepository)

	ticketUsageGroup := r.Group("/ticket-usages")
	ticketUsageGroup.Use(middleware.AuthProtected(db))
	{
		ticketUsageGroup.GET("/success", middleware.RoleRequired("provider", "admin"), ticketUsageHandler.GetManyProvider)
		ticketUsageGroup.GET("/", middleware.AccessPermission(ticketMiddleware), ticketUsageHandler.GetMany)
		ticketUsageGroup.GET("/:id", middleware.AccessPermission(ticketMiddleware), ticketUsageHandler.GetOne)
		ticketUsageGroup.PUT("/:id", middleware.RoleRequired("provider"), ticketUsageHandler.UpdateOne)
		ticketUsageGroup.DELETE("/:id", middleware.RoleRequired("admin"), ticketUsageHandler.DeleteOne)
	}
}
