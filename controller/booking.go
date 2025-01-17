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

var responseBookingName = "booking"
var responseBooking helper.ResponseMessage

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
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	bookings, err := h.repositoryBooking.GetMany(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed(responseBookingName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseBooking.GetSuccessfully(responseBookingName), bookings))
}

func (h *BookingHandler) GetOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseBooking.IdFailed(responseBookingName)))
		return
	}

	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed(responseBookingName)))
		return
	}

	response := &entity.BookingWithRelation{
		ID:            booking.ID,
		UserID:        booking.UserID,
		ScheduleID:    booking.ScheduleID,
		TicketCode:    booking.TicketCode,
		TotalAmount:   booking.TotalAmount,
		BookingStatus: booking.BookingStatus,
		SeatNumbers:   booking.SeatNumbers,
		CreatedAt:     booking.CreatedAt,
		UpdatedAt:     booking.UpdatedAt,
		User:          booking.User,
		Schedule:      booking.Schedule,
	}

	if booking.Payment != nil && booking.Payment.ID != 0 {
		response.Payment = booking.Payment
	} else {
		response.Payment = nil
	}

	if booking.Cancellation != nil && booking.Cancellation.ID != 0 {
		response.Cancellation = booking.Cancellation
	} else {
		response.Cancellation = nil
	}

	if booking.TicketUsage != nil && booking.TicketUsage.ID != 0 {
		response.TicketUsage = booking.TicketUsage
	} else {
		response.TicketUsage = nil
	}

	if booking.Refund != nil && booking.Refund.ID != 0 {
		response.Refund = booking.Refund
	} else {
		response.Refund = nil
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseBooking.GetSuccessfully(responseBookingName), response))
}

func (h *BookingHandler) CreateOne(ctx *gin.Context) {
	var booking entity.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseBooking.RequestFailed(responseBookingName)))
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

	schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseBooking.GetFailed("schedule")))
		return
	}

	if schedule.TicketPrice == 0 {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Ticket price cannot be zero"))
		return
	}

	if len(booking.SeatNumbers) == 0 {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("At least one seat must be selected"))
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

	booking.TicketCode = helper.GenerateTicketNumber(userID, booking.ScheduleID)
	booking.UserID = userID
	booking.TotalAmount = schedule.TicketPrice * float64(len(booking.SeatNumbers))

	if booking.TotalAmount <= 0 {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Total amount must be greater than zero"))
		return
	}

	createdBooking, err := h.repositoryBooking.CreateOne(context.Background(), &booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed(responseBookingName)))
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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.UpdateFailed("schedule")))
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: "Create booking is successfully",
	}
	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
		return
	}

	notification := entity.Notification{
		UserID:  booking.UserID,
		Message: "Booking ticket successfully, please payment to complete!",
	}

	_, err = h.notificationRepository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("notification")))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse(responseBooking.CreateSuccessfully(responseBookingName), createdBooking))
}
func (h *BookingHandler) UpdateOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseBooking.IdFailed(responseBookingName)))
		return
	}

	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed(responseBookingName)))
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
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
			return
		}
	case "cancel":
		activityLog := entity.ActivityLog{
			UserID:      booking.UserID,
			Description: fmt.Sprintf("Booking id %d updated from %s to payment successfully", booking.ID, booking.BookingStatus),
		}

		_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
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
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
			return
		}
	}
	var updateData entity.Booking
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseBooking.RequestFailed(responseBookingName)))
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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.UpdateFailed(responseBookingName)))
		return
	}
	if updateData.BookingStatus == "cancel" {
		cancellation := entity.Cancellation{
			BookingID:          booking.ID,
			CancellationReason: fmt.Sprintf("Booking id %d cancellation is successfully", booking.ID),
		}
		_, err := h.cancellationRepository.CreateOne(context.Background(), &cancellation)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("cancellation")))
			return
		}
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseBooking.GetFailed("schedule")))
			return
		}
		_, err = h.scheduleRepository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.UpdateFailed(responseBookingName)))
			return
		}
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseBooking.UpdateSuccessfully(responseBookingName), updatedBooking))
}

func (h *BookingHandler) DeleteOne(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseBooking.IdFailed(responseBookingName)))
		return
	}
	booking, err := h.repositoryBooking.GetOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed(responseBookingName)))
		return
	}
	if booking.BookingStatus == "pending" {
		var seatsData = make(map[int]interface{})
		for _, seat := range booking.SeatNumbers {
			seatsData[int(seat)] = "cancel"
		}
		schedule, err := h.scheduleRepository.GetOne(context.Background(), booking.ScheduleID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseBooking.IdFailed("schedule")))
			return
		}
		_, err = h.scheduleRepository.UpdateSeatsStatus(context.Background(), schedule.ID, seatsData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed("schedule")))
			return
		}
	}
	err = h.repositoryBooking.DeleteOne(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.DeleteFailed(responseBookingName)))
		return
	}
	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: fmt.Sprintf("Deleted booking id %d status %s successfully", booking.ID, booking.BookingStatus),
	}

	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseBooking.DeleteSuccessfully(responseBookingName), nil))
}
