package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type ActivityLogHandler struct {
	repository entity.ActivityLogRepository
}

func NewActivityLogHandler(repo entity.ActivityLogRepository) *ActivityLogHandler {
	return &ActivityLogHandler{
		repository: repo,
	}
}

func (h *ActivityLogHandler) GetMany(ctx *gin.Context) {
	activityLogs, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch activity logs"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data activity logs successfully", activityLogs))
}

func (h *ActivityLogHandler) GetManyByUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid user ID"))
		return
	}

	activityLogs, err := h.repository.GetManyByUser(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch activity logs"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data activity logs successfully", activityLogs))
}

func (h *ActivityLogHandler) GetOne(ctx *gin.Context) {
	activityLogId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid activity log ID"))
		return
	}

	activityLog, err := h.repository.GetOne(context.Background(), uint(activityLogId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch activity log"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data activity log successfully", activityLog))
}

func (h *ActivityLogHandler) CreateOne(ctx *gin.Context) {
	var activityLog entity.ActivityLog
	if err := ctx.ShouldBindJSON(&activityLog); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createdActivityLog, err := h.repository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create notification successfully", createdActivityLog))
}

func (h *ActivityLogHandler) DeleteOne(ctx *gin.Context) {
	activityLogId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid activity log ID"))
		return
	}

	err = h.repository.DeleteOne(context.Background(), uint(activityLogId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete activity log"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete activity log successfully", nil))
}

func (h *ActivityLogHandler) DeleteMany(ctx *gin.Context) {
	if err := h.repository.DeleteMany(context.Background()); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete all activity logs"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Deleted all activity logs successfully", nil))
}
