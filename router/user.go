package router

import (
	"github.com/Junx27/ticket-booking/controller"
	"github.com/Junx27/ticket-booking/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userHandler := controller.NewUserHandler(userRepository)

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", userHandler.GetMany)
		userGroup.GET("/:id", userHandler.GetOne)
		userGroup.PUT("/:id", userHandler.UpdateOne)
		userGroup.DELETE("/:id", userHandler.DeleteOne)
	}
}
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
