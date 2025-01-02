package controller

import (
	"context"
	"net/http"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository entity.UserRepository
}

func NewUserHandler(repo entity.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repo,
	}
}

func (h *UserHandler) GetMany(ctx *gin.Context) {
	users, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
