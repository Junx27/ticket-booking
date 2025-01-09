package helper

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromCookie(ctx *gin.Context) (uint, error) {

	cookie, err := ctx.Cookie("token")
	if err != nil {
		return 0, fmt.Errorf("Authorization token is missing")
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Failed to parse token claims")
	}
	userIDFloat, ok := (*claims)["id"].(float64)
	if !ok || userIDFloat == 0 {
		return 0, fmt.Errorf("Missing or invalid user ID in token")
	}
	userID := uint(userIDFloat)
	return userID, nil
}
