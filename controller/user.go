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

func (h *UserHandler) GetOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID Not Found"))
		return
	}

	user, err := h.repository.GetOne(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch user"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data user successfully", user))
}

func (h *UserHandler) CreateOne(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createdUser, err := h.repository.CreateOne(context.Background(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create user"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create user successfully", createdUser))
}

func (h *UserHandler) UpdateOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updatedUser, err := h.repository.UpdateOne(context.Background(), uint(userId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update user"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update user successfully", updatedUser))
}

func (h *UserHandler) DeleteOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete user"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete user successfully", nil))
}
