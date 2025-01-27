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
	paymentMiddleware := paymentRepository.(*repository.PaymentRepository)

	paymentGroup := r.Group("/payments")
	paymentGroup.Use(middleware.AuthProtected(db))
	{
		paymentGroup.GET("/success", middleware.RoleRequired("provider", "admin"), paymentHandler.GetManyProvider)
		paymentGroup.GET("/", middleware.AccessPermission(paymentMiddleware), paymentHandler.GetMany)
		paymentGroup.GET("/:id", middleware.AccessPermission(paymentMiddleware), paymentHandler.GetOne)
		paymentGroup.POST("/", middleware.RoleRequired("customer"), paymentHandler.CreateOne)
		paymentGroup.PUT("/:id", middleware.RoleRequired("admin"), paymentHandler.UpdateOne)
		paymentGroup.DELETE("/:id", middleware.RoleRequired("admin"), paymentHandler.DeleteOne)
	}
}
