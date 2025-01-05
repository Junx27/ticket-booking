package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/Junx27/ticket-booking/utils"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	repository      entity.BookingRepository
	scheduleHandler *ScheduleHandler
}

func NewBookingHandler(
	bookingRepo entity.BookingRepository,
	scheduleHandler *ScheduleHandler,
) *BookingHandler {
	return &BookingHandler{
		repository:      bookingRepo,
		scheduleHandler: scheduleHandler,
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	schedule, err := h.scheduleHandler.repository.GetOne(context.Background(), booking.ScheduleID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	var unavailableSeats []int
	for _, seat := range booking.SeatNumbers {
		if !schedule.IsSeatAvailable(seat) {
			unavailableSeats = append(unavailableSeats, int(seat))
		}
	}

	if len(unavailableSeats) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Seats %v are not available", unavailableSeats),
		})
		return
	}

	booking.NumberOfTickets = helper.GenerateTicketNumber()
	createdBooking, err := h.repository.CreateOne(context.Background(), &booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}
	seatsData := make(map[string]string)
	for _, seat := range booking.SeatNumbers {
		seatsData[fmt.Sprint(seat)] = "booked"
	}
	err = utils.UpdateSeatsStatus(schedule.ID, seatsData)
	if err != nil {
		h.repository.DeleteOne(context.Background(), createdBooking.ID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seat status"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"data":    createdBooking,
	})
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
