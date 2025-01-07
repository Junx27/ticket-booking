package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	repository         entity.BookingRepository
	scheduleHandler    *ScheduleHandler
	activityLogHandler *ActivityLogHandler
}

func NewBookingHandler(
	bookingRepo entity.BookingRepository,
	scheduleHandler *ScheduleHandler,
	activityLogHandler *ActivityLogHandler,

) *BookingHandler {
	return &BookingHandler{
		repository:         bookingRepo,
		scheduleHandler:    scheduleHandler,
		activityLogHandler: activityLogHandler,
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

	if schedule.TicketPrice == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ticket price cannot be zero"})
		return
	}

	if len(booking.SeatNumbers) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "At least one seat must be selected"})
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

	booking.TicketCode = helper.GenerateTicketNumber(booking.UserID, booking.ScheduleID)

	booking.TotalAmount = schedule.TicketPrice * float64(len(booking.SeatNumbers))

	if booking.TotalAmount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Total amount must be greater than zero"})
		return
	}

	createdBooking, err := h.repository.CreateOne(context.Background(), &booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	seatsData := make(map[int]interface{})
	for _, seat := range booking.SeatNumbers {
		seatsData[int(seat)] = "booked"
	}

	scheduleId := schedule.ID
	_, err = h.scheduleHandler.repository.UpdateSeatsStatus(context.Background(), scheduleId, seatsData)
	if err != nil {
		h.repository.DeleteOne(context.Background(), createdBooking.ID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		ActionType:  "CREATE BOOKING",
		Description: "Create booking is successfully",
	}
	createdActivityLog, err := h.activityLogHandler.repository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":      "Booking created successfully",
		"booking":      createdBooking,
		"activity_log": createdActivityLog,
	})
}
func (h *BookingHandler) UpdateOne(ctx *gin.Context) {
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
	var updateData entity.Booking
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}
	if updateData.BookingStatus == "" {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Booking status is required"))
		return
	}
	updateFields := map[string]interface{}{
		"id":             booking.ID,
		"user_id":        booking.UserID,
		"schedule_id":    booking.ScheduleID,
		"ticket_code":    booking.TicketCode,
		"total_amount":   booking.TotalAmount,
		"booking_status": updateData.BookingStatus,
		"seat_numbers":   booking.SeatNumbers,
		"created_at":     booking.CreatedAt,
	}
	updatedBooking, err := h.repository.UpdateOne(context.Background(), uint(bookingId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update booking"))
		return
	}
	if updateData.BookingStatus == "cancel" {
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleHandler.repository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse("Schedule not found"))
			return
		}
		_, err = h.scheduleHandler.repository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update schedule"))
			return
		}
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data booking successfully", updatedBooking))
}

func (h *BookingHandler) DeleteOne(ctx *gin.Context) {
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
	if booking.BookingStatus == "pending" {
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleHandler.repository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse("Schedule not found"))
			return
		}
		_, err = h.scheduleHandler.repository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update schedule"))
			return
		}
	}
	err = h.repository.DeleteOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete booking"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data booking successfully", nil))
}
