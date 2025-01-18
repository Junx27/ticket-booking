package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/middleware"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := controller.NewUserHandler(userRepository)

	userGroup := r.Group("/users")
	userGroup.Use(middleware.AuthProtected(db))
	{
		userGroup.GET("/", middleware.RoleRequired("admin"), userHandler.GetMany)
		userGroup.GET("/:id", middleware.RoleRequired("admin"), userHandler.GetOne)
		userGroup.PUT("/:id", middleware.RoleRequired("admin"), userHandler.UpdateOne)
		userGroup.PUT("/provider/:id", middleware.RoleRequired("admin"), userHandler.UpdateOneProvider)
		userGroup.DELETE("/:id", middleware.RoleRequired("admin"), userHandler.DeleteOne)
	}
}
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
