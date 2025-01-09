package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRefundRouter(r *gin.Engine, db *gorm.DB) {
	refundRepository := repository.NewRefundRepository(db)
	refundHandler := controller.NewRefundHandler(refundRepository)

	refundGroup := r.Group("/refunds")
	refundGroup.Use(middleware.AuthProtected(db))
	{
		refundGroup.GET("/", refundHandler.GetMany)
		refundGroup.GET("/:id", refundHandler.GetOne)
		refundGroup.POST("/", refundHandler.CreateOne)
		refundGroup.PUT("/:id", refundHandler.UpdateOne)
		refundGroup.DELETE("/:id", refundHandler.DeleteOne)
	}
}
