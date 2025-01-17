package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseScheduleName = "schedule"
var responseSchedule helper.ResponseMessage

type ScheduleHandler struct {
	repository entity.ScheduleRepository
}

func NewScheduleHandler(repo entity.ScheduleRepository) *ScheduleHandler {
	return &ScheduleHandler{
		repository: repo,
	}
}

func (h *ScheduleHandler) GetMany(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	schedules, err := h.repository.GetMany(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseSchedule.GetFailed(responseScheduleName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseSchedule.GetSuccessfully(responseScheduleName), schedules))
}

func (h *ScheduleHandler) GetOne(ctx *gin.Context) {
	scheduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.IdFailed(responseScheduleName)))
		return
	}

	schedule, err := h.repository.GetOne(context.Background(), uint(scheduleId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseSchedule.GetFailed(responseScheduleName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseSchedule.GetSuccessfully(responseScheduleName), schedule))
}
func (h *ScheduleHandler) CreateOne(ctx *gin.Context) {
	var schedule entity.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.RequestFailed(responseScheduleName)))
		return
	}
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	schedule.UserID = userID
	createdSchedule, err := h.repository.CreateOne(context.Background(), &schedule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseSchedule.CreateFailed(responseScheduleName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseSchedule.CreateSuccessfully(responseScheduleName), createdSchedule))
}

func (h *ScheduleHandler) UpdateOne(ctx *gin.Context) {
	scheduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.IdFailed(responseScheduleName)))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.RequestFailed(responseScheduleName)))
		return
	}

	updatedSchedule, err := h.repository.UpdateOne(context.Background(), uint(scheduleId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseSchedule.UpdateFailed(responseScheduleName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseSchedule.UpdateSuccessfully(responseScheduleName), updatedSchedule))
}

func (h *ScheduleHandler) UpdateSeatsStatus(ctx *gin.Context) {
	scheduleIdParam := ctx.Param("id")
	scheduleId, err := strconv.Atoi(scheduleIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.IdFailed(responseScheduleName)))
		return
	}

	var seatsData map[string]string
	if err := ctx.ShouldBindJSON(&seatsData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid JSON format"))
		return
	}

	convertedSeatsData := make(map[int]interface{})
	for key, value := range seatsData {
		intKey, err := strconv.Atoi(key)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid seat ID"))
			return
		}

		if value == "booked" {
			convertedSeatsData[intKey] = "booked"
		} else if value == "cancel" {
			convertedSeatsData[intKey] = "cancel"
		} else {
			ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid seat status"))
			return
		}
	}

	updatedSchedule, err := h.repository.UpdateSeatsStatus(ctx, uint(scheduleId), convertedSeatsData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update seats status"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Seats updated successfully", updatedSchedule))
}

func (h *ScheduleHandler) DeleteOne(ctx *gin.Context) {
	scheduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseSchedule.IdFailed(responseScheduleName)))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(scheduleId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseSchedule.DeleteFailed(responseScheduleName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseSchedule.DeleteSuccessfully(responseScheduleName), nil))
}
