package main

import (
	"github.com/Junx27/ticket-booking/config"
	"github.com/Junx27/ticket-booking/database"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/Junx27/ticket-booking/router"
	"github.com/Junx27/ticket-booking/service"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewEnvConfig()

	db := database.Init(cfg, database.DBMigrator)
	database.SeedUsers(db)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	router.SetupAuthRouter(r, authService.(*service.AuthService))
	router.SetupUserRouter(r, db)
	router.SetupProviderRouter(r, db)
	router.SetupScheduleRouter(r, db)
	router.SetupBookingRouter(r, db)
	router.SetupPaymentRouter(r, db)
	router.SetupNotificationRouter(r, db)
	router.SetupCancellationRouter(r, db)
	router.SetupActivityLogRouter(r, db)
	router.SetupRefundRouter(r, db)
	router.SetupTicketUsageRouter(r, db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "booking service is work! - author Tri Saptono(dibimbing)",
		})
	})

	r.Run(":8080")
}
