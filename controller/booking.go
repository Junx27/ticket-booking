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
	repositoryBooking      entity.BookingRepository
	scheduleRepository     entity.ScheduleRepository
	activityLogRepository  entity.ActivityLogRepository
	cancellationRepository entity.CancellationRepository
	notificationRepository entity.NotificationRepository
}

func NewBookingHandler(
	repositoryBooking entity.BookingRepository,
	scheduleRepository entity.ScheduleRepository,
	activityLogRepository entity.ActivityLogRepository,
	cancellationRepository entity.CancellationRepository,
	notificationRepository entity.NotificationRepository,

) *BookingHandler {
	return &BookingHandler{
		repositoryBooking:      repositoryBooking,
		scheduleRepository:     scheduleRepository,
		activityLogRepository:  activityLogRepository,
		cancellationRepository: cancellationRepository,
		notificationRepository: notificationRepository,
	}
}

func (h *BookingHandler) GetMany(ctx *gin.Context) {
	bookings, err := h.repositoryBooking.GetMany(context.Background())
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

	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
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

	schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
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

	createdBooking, err := h.repositoryBooking.CreateOne(context.Background(), &booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	seatsData := make(map[int]interface{})
	for _, seat := range booking.SeatNumbers {
		seatsData[int(seat)] = "booked"
	}

	scheduleId := schedule.ID
	_, err = h.scheduleRepository.UpdateSeatsStatus(context.Background(), scheduleId, seatsData)
	if err != nil {
		h.repositoryBooking.DeleteOne(context.Background(), createdBooking.ID)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: "Create booking is successfully",
	}
	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
		return
	}

	notification := entity.Notification{
		UserID:  booking.UserID,
		Message: "Booking ticket successfully, please payment to complete!",
	}

	_, err = h.notificationRepository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create notification"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Booking created successfully", createdBooking))
}
func (h *BookingHandler) UpdateOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid booking ID Not Found"))
		return
	}

	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch booking"))
		return
	}

	switch booking.BookingStatus {
	case "pending":
		activityLog := entity.ActivityLog{
			UserID:      booking.UserID,
			Description: fmt.Sprintf("Booking id %d updated from %s to confirm successfully", booking.ID, booking.BookingStatus),
		}

		_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
			return
		}
	case "cancel":
		activityLog := entity.ActivityLog{
			UserID:      booking.UserID,
			Description: fmt.Sprintf("Booking id %d updated from %s to payment successfully", booking.ID, booking.BookingStatus),
		}

		_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
			return
		}
	case "payment":
		activityLog := entity.ActivityLog{
			UserID:      booking.UserID,
			ActionType:  "NOTICE",
			Description: fmt.Sprintf("Booking id %d updated from %s to cancel successfully", booking.ID, booking.BookingStatus),
		}

		_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
			return
		}
	default:
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
	updatedBooking, err := h.repositoryBooking.UpdateOne(context.Background(), uint(bookingId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update booking"))
		return
	}
	if updateData.BookingStatus == "cancel" {
		cancellation := entity.Cancellation{
			BookingID:          booking.ID,
			CancellationReason: fmt.Sprintf("Booking id %d cancellation is successfully", booking.ID),
		}
		_, err := h.cancellationRepository.CreateOne(context.Background(), &cancellation)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create cancellation"))
			return
		}
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse("Schedule not found"))
			return
		}
		_, err = h.scheduleRepository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
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
	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch booking"))
		return
	}
	if booking.BookingStatus == "pending" {
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse("Schedule not found"))
			return
		}
		_, err = h.scheduleRepository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update schedule"))
			return
		}
	}
	err = h.repositoryBooking.DeleteOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete booking"))
		return
	}
	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: fmt.Sprintf("Deleted booking id %d status %s successfully", booking.ID, booking.BookingStatus),
	}

	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data booking successfully", nil))
}
