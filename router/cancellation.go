package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCancellationRouter(r *gin.Engine, db *gorm.DB) {
	cancellationRepository := repository.NewCancellationRepository(db)
	cancellationHandler := controller.NewCancellationHandler(cancellationRepository)

	cancellationGroup := r.Group("/cancellations")
	cancellationGroup.Use(middleware.AuthProtected(db))
	{
		cancellationGroup.GET("/", cancellationHandler.GetMany)
		cancellationGroup.GET("/booking/:booking_id", cancellationHandler.GetManyByBookingID)
		cancellationGroup.GET("/:id", cancellationHandler.GetOne)
		cancellationGroup.POST("/", middleware.RoleRequired("customer"), cancellationHandler.CreateOne)
		cancellationGroup.DELETE("/:id", cancellationHandler.DeleteOne)
		cancellationGroup.DELETE("/", cancellationHandler.DeleteMany)
	}
}
