package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/Junx27/ticket-booking/repository"
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

func RoleRequired(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Authorization token is missing",
			})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		role, ok := (*claims)["role"].(string)
		if !ok || role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": "Role is missing or invalid in token",
			})
			c.Abort()
			return
		}
		if !isRoleAllowed(role, allowedRoles) {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "fail",
				"message": "You do not have permission to access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isRoleAllowed(role string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

func AccessPermissionAction(repo repository.HasUserID) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resourceID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid resource ID"))
			ctx.Abort()
			return
		}

		userID, err := helper.GetUserIDFromCookie(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		ownerID, err := repo.GetUserID(uint(resourceID))
		if err != nil {
			if err.Error() == "provider not found" {
				ctx.JSON(http.StatusNotFound, helper.FailedResponse("Resource not found"))
			} else {
				ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Internal server error"))
			}
			ctx.Abort()
			return
		}
		fmt.Printf("UserID: %d, OwnerID: %d\n", userID, ownerID)
		if userID != ownerID && !isAdmin(userID) {
			ctx.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AccessPermissionView(repo *repository.ProviderRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramID := ctx.Param("user_id")
		userID, err := helper.GetUserIDFromCookie(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		if paramID != strconv.Itoa(int(userID)) && !isAdmin(userID) {
			ctx.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func isAdmin(userID uint) bool {
	return userID == 1
}
