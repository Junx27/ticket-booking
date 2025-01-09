package controller

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("JWT_SECRET")

type ScheduleHandler struct {
	repository entity.ScheduleRepository
}

func NewScheduleHandler(repo entity.ScheduleRepository) *ScheduleHandler {
	return &ScheduleHandler{
		repository: repo,
	}
}

func (h *ScheduleHandler) GetMany(ctx *gin.Context) {
	schedules, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch schedules"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data schedules successfully", schedules))
}

func (h *ScheduleHandler) GetOne(ctx *gin.Context) {
	scheduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid schedule ID"))
		return
	}

	schedule, err := h.repository.GetOne(context.Background(), uint(scheduleId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch schedule"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data schedule successfully", schedule))
}
func (h *ScheduleHandler) CreateOne(ctx *gin.Context) {
	var schedule entity.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create schedule"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Create schedule successfully", createdSchedule))
}

func (h *ScheduleHandler) UpdateOne(ctx *gin.Context) {
	scheduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid schedule ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updatedSchedule, err := h.repository.UpdateOne(context.Background(), uint(scheduleId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update schedule"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update schedule successfully", updatedSchedule))
}

func (h *ScheduleHandler) UpdateSeatsStatus(ctx *gin.Context) {
	scheduleIdParam := ctx.Param("id")
	scheduleId, err := strconv.Atoi(scheduleIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid schedule ID"))
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
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid schedule ID"))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(scheduleId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete schedule"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete schedule successfully", nil))
}
