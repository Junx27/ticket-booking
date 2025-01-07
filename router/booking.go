package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookingRouter(r *gin.Engine, db *gorm.DB) {
	bookingRepository := repository.NewBookingRepository(db)
	scheduleRepository := repository.NewScheduleRepository(db)
	activityLogRepository := repository.NewActivityLogRepository(db)
	cancellationRepository := repository.NewCancellationRepository(db)
	scheduleHandler := controller.NewScheduleHandler(scheduleRepository)
	activityLogHandler := controller.NewActivityLogHandler(activityLogRepository)
	cancellatonHandler := controller.NewCancellationHandler(cancellationRepository)
	bookingHandler := controller.NewBookingHandler(bookingRepository, scheduleHandler, activityLogHandler, cancellatonHandler)

	bookingGroup := r.Group("/bookings")
	{
		bookingGroup.GET("/", bookingHandler.GetMany)
		bookingGroup.GET("/:id", bookingHandler.GetOne)
		bookingGroup.POST("/", bookingHandler.CreateOne)
		bookingGroup.PUT("/:id", bookingHandler.UpdateOne)
		bookingGroup.DELETE("/:id", bookingHandler.DeleteOne)
	}
}
