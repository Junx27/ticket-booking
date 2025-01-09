package controller

import (
	"context"
	"time"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	service entity.AuthService
}

func NewAuthHandler(
	service entity.AuthService,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	creds := &entity.AuthCredentials{}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	token, user, err := h.service.Login(ctxTimeout, creds)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Successfully logged in",
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	creds := &entity.AuthCredentials{}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": "Please, provide a valid name, email, and password",
		})
		return
	}

	token, user, err := h.service.Register(ctxTimeout, creds)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "success",
		"message": "Successfully registered",
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": "Token is required",
		})
		return
	}

	err := h.service.Logout(ctx, token)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "fail",
			"message": "Logout failed: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Successfully logged out",
	})
}
