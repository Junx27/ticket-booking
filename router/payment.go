package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPaymentRouter(r *gin.Engine, db *gorm.DB) {
	paymentRepository := repository.NewPaymentRepository(db)
	ticketUsageRepository := repository.NewTicketUsageRepository(db)
	ticketUsageHandler := controller.NewTicketUsageHandler(ticketUsageRepository)
	paymentHandler := controller.NewPaymentHandler(paymentRepository, ticketUsageHandler)

	paymentGroup := r.Group("/payments")
	{
		paymentGroup.GET("/", paymentHandler.GetMany)
		paymentGroup.GET("/:id", paymentHandler.GetOne)
		paymentGroup.POST("/", paymentHandler.CreateOne)
		paymentGroup.PUT("/:id", paymentHandler.UpdateOne)
		paymentGroup.DELETE("/:id", paymentHandler.DeleteOne)
	}
}
