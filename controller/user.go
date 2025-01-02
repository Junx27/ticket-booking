package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch users"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data users successfully", users))
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	user, err := h.repository.GetByID(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch user"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data user successfully", user))
}
