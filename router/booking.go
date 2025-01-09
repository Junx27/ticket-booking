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

	bookingGroup := r.Group("/bookings")
	bookingGroup.Use(middleware.AuthProtected(db))
	{
		bookingGroup.GET("/", middleware.RoleRequired("admin", "customer"), bookingHandler.GetMany)
		bookingGroup.GET("/:id", bookingHandler.GetOne)
		bookingGroup.POST("/", bookingHandler.CreateOne)
		bookingGroup.PUT("/:id", bookingHandler.UpdateOne)
		bookingGroup.DELETE("/:id", bookingHandler.DeleteOne)
	}
}
