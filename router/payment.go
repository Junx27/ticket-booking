package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPaymentRouter(r *gin.Engine, db *gorm.DB) {
	paymentRepository := repository.NewPaymentRepository(db)
	ticketUsageRepository := repository.NewTicketUsageRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	activityLogRepository := repository.NewActivityLogRepository(db)
	bookingRepository := repository.NewBookingRepository(db)
	paymentHandler := controller.NewPaymentHandler(paymentRepository, ticketUsageRepository, bookingRepository, activityLogRepository, notificationRepository)

	paymentGroup := r.Group("/payments")
	paymentGroup.Use(middleware.AuthProtected(db))
	{
		paymentGroup.GET("/", paymentHandler.GetMany)
		paymentGroup.GET("/:id", paymentHandler.GetOne)
		paymentGroup.POST("/", middleware.RoleRequired("customer"), paymentHandler.CreateOne)
		paymentGroup.PUT("/:id", paymentHandler.UpdateOne)
		paymentGroup.DELETE("/:id", paymentHandler.DeleteOne)
	}
}
