package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	repository entity.NotificationRepository
}

func NewNotificationHandler(repo entity.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{
		repository: repo,
	}
}

func (h *NotificationHandler) GetMany(ctx *gin.Context) {
	notifications, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch notifications"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data notifications successfully", notifications))
}

func (h *NotificationHandler) GetManyByUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	notifications, err := h.repository.GetManyByUser(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch notifications"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data notifications successfully", notifications))
}

func (h *NotificationHandler) GetOne(ctx *gin.Context) {
	notificationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid notification ID"))
		return
	}

	notification, err := h.repository.GetOne(context.Background(), uint(notificationId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch notification"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data notification successfully", notification))
}

func (h *NotificationHandler) CreateOne(ctx *gin.Context) {
	var notification entity.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createdNotification, err := h.repository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create notification"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create notification successfully", createdNotification))
}

func (h *NotificationHandler) UpdateOne(ctx *gin.Context) {
	notificationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid notification ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updatedNotification, err := h.repository.UpdateOne(context.Background(), uint(notificationId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update notification"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update notification successfully", updatedNotification))
}
func (h *NotificationHandler) DeleteOne(ctx *gin.Context) {
	notificationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid notification ID"))
		return
	}

	err = h.repository.DeleteOne(context.Background(), uint(notificationId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete notification"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete notification successfully", nil))
}

func (h *NotificationHandler) DeleteAllByUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	err = h.repository.DeleteAllByUser(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete notifications"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete notifications successfully", nil))
}
