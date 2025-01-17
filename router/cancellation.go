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
	cancellationMiddleware := cancellationRepository.(*repository.CancellationRepository)

	cancellationGroup := r.Group("/cancellations")
	cancellationGroup.Use(middleware.AuthProtected(db))
	{
		cancellationGroup.GET("/", middleware.AccessPermission(cancellationMiddleware), cancellationHandler.GetMany)
		cancellationGroup.GET("/booking/:booking_id", middleware.AccessPermission(cancellationMiddleware), cancellationHandler.GetManyByBookingID)
		cancellationGroup.GET("/:id", middleware.AccessPermission(cancellationMiddleware), cancellationHandler.GetOne)
		cancellationGroup.POST("/", middleware.RoleRequired("customer"), cancellationHandler.CreateOne)
		cancellationGroup.DELETE("/:id", middleware.RoleRequired("admin"), cancellationHandler.DeleteOne)
		cancellationGroup.DELETE("/", middleware.RoleRequired("admin"), cancellationHandler.DeleteMany)
	}
}
