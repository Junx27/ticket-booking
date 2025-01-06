package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTicketUsageRouter(r *gin.Engine, db *gorm.DB) {
	ticketUsageRepository := repository.NewTicketUsageRepository(db)
	ticketUsageHandler := controller.NewTicketUsageHandler(ticketUsageRepository)

	ticketUsageGroup := r.Group("/ticket-usages")
	{
		ticketUsageGroup.GET("/", ticketUsageHandler.GetMany)
		ticketUsageGroup.GET("/:id", ticketUsageHandler.GetOne)
		ticketUsageGroup.POST("/", ticketUsageHandler.CreateOne)
		ticketUsageGroup.PUT("/:id", ticketUsageHandler.UpdateOne)
		ticketUsageGroup.DELETE("/:id", ticketUsageHandler.DeleteOne)
	}
}
