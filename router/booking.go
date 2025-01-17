package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookingRouter(r *gin.Engine, db *gorm.DB) {
	bookingRepository := repository.NewBookingRepository(db)
	scheduleRepository := repository.NewScheduleRepository(db)
	activityLogRepository := repository.NewActivityLogRepository(db)
	cancellationRepository := repository.NewCancellationRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	bookingHandler := controller.NewBookingHandler(bookingRepository, scheduleRepository, activityLogRepository, cancellationRepository, notificationRepository)
	bookingMiddleware := bookingRepository.(*repository.BookingRepository)

	bookingGroup := r.Group("/bookings")
	bookingGroup.Use(middleware.AuthProtected(db))
	{
		bookingGroup.GET("/", middleware.AccessPermission(bookingMiddleware), bookingHandler.GetMany)
		bookingGroup.GET("/:id", middleware.AccessPermission(bookingMiddleware), bookingHandler.GetOne)
		bookingGroup.POST("/", middleware.RoleRequired("customer"), bookingHandler.CreateOne)
		bookingGroup.PUT("/:id", middleware.AccessPermission(bookingMiddleware), middleware.RoleRequired("customer"), bookingHandler.UpdateOne)
		bookingGroup.DELETE("/:id", middleware.AccessPermission(bookingMiddleware), middleware.RoleRequired("customer", "admin"), bookingHandler.DeleteOne)
	}
}
