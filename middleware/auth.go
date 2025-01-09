package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Missing Authorization Header",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid Token Format",
			})
			ctx.Abort()
			return
		}

		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": fmt.Sprintf("Unauthorized - %v", err),
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid Claims",
			})
			ctx.Abort()
			return
		}
		userId, ok := claims["id"].(float64)
		if !ok {
			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - Invalid User ID in Token",
			})
			ctx.Abort()
			return
		}
		var user entity.User
		if err := db.Where("id = ?", int64(userId)).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			ctx.JSON(401, gin.H{
				"status":  "fail",
				"message": "Unauthorized - User Not Found",
			})
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
