package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	repository entity.BookingRepository
}

func NewBookingHandler(repo entity.BookingRepository) *BookingHandler {
	return &BookingHandler{
		repository: repo,
	}
}

func (h *BookingHandler) GetMany(ctx *gin.Context) {
	bookings, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch bookings"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data bookings successfully", bookings))
}

func (h *BookingHandler) GetOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid booking ID Not Found"))
		return
	}

	booking, err := h.repository.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch booking"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data booking successfully", booking))
}

func (h *BookingHandler) CreateOne(ctx *gin.Context) {
	var booking entity.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createdBooking, err := h.repository.CreateOne(context.Background(), &booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create booking"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create data booking successfully", createdBooking))
}

func (h *BookingHandler) UpdateOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid booking ID Not Found"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updatedBooking, err := h.repository.UpdateOne(context.Background(), uint(bookingId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update booking"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data booking successfully", updatedBooking))
}

func (h *BookingHandler) DeleteOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid booking ID Not Found"))
		return
	}

	err = h.repository.DeleteOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete booking"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data booking successfully", nil))
}
