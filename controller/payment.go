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

var responsePayment helper.ResponsePayment
var responeNotification helper.ResponseNotification
var responseActivityLog helper.ResponseActivityLog
var responseTicketUsage helper.ResponseTicketUsage
var responseBooking helper.ResponseBooking

type PaymentHandler struct {
	paymentRepository      entity.PaymentRepository
	ticketUsageRepositiry  entity.TicketUsageRepository
	bookingRepository      entity.BookingRepository
	activityLogRepository  entity.ActivityLogRepository
	notificationRepository entity.NotificationRepository
}

func NewPaymentHandler(
	paymentRepository entity.PaymentRepository,
	ticketUsageRepositiry entity.TicketUsageRepository,
	bookingRepository entity.BookingRepository,
	activityLogRepository entity.ActivityLogRepository,
	notificationRepository entity.NotificationRepository,
) *PaymentHandler {
	return &PaymentHandler{
		paymentRepository:      paymentRepository,
		ticketUsageRepositiry:  ticketUsageRepositiry,
		bookingRepository:      bookingRepository,
		activityLogRepository:  activityLogRepository,
		notificationRepository: notificationRepository,
	}
}

func (h *PaymentHandler) GetMany(ctx *gin.Context) {
	payments, err := h.paymentRepository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed()))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.GetSuccessfully(), payments))
}

func (h *PaymentHandler) GetOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed()))
		return
	}

	payment, err := h.paymentRepository.GetOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.GetSuccessfully(), payment))
}

func (h *PaymentHandler) CreateOne(ctx *gin.Context) {
	var payment entity.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.RequestFailed()))
		return
	}

	booking, err := h.bookingRepository.GetOne(context.Background(), payment.BookingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.GetFailed()))
		return
	}

	payment.PaymentAmount = booking.TotalAmount
	createdPayment, err := h.paymentRepository.CreateOne(context.Background(), &payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.CreateFailed()))
		return
	}

	ticketUsage := entity.TicketUsage{
		BookingID: payment.BookingID,
	}
	_, err = h.ticketUsageRepositiry.CreateOne(context.Background(), &ticketUsage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.CreateFailed()))
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: fmt.Sprintf("Booking id %d status %s payment successfully", booking.ID, booking.BookingStatus),
	}

	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseActivityLog.CreateFailed()))
		return
	}
	notification := entity.Notification{
		UserID:  booking.UserID,
		Message: fmt.Sprintf("Booking id %d status %s payment completed", booking.ID, booking.BookingStatus),
	}

	_, err = h.notificationRepository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responeNotification.InvalidCreate()))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse(responsePayment.CreateSuccessfully(), createdPayment))
}

func (h *PaymentHandler) UpdateOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed()))
		return
	}

	payment, err := h.paymentRepository.GetOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed()))
		return
	}

	var updateData entity.Payment
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.RequestFailed()))
		return
	}

	updateFields := map[string]interface{}{
		"id":             payment.ID,
		"booking_id":     payment.BookingID,
		"payment_amount": payment.PaymentAmount,
		"payment_status": updateData.PaymentStatus,
		"payment_method": payment.PaymentMethod,
	}

	updatedPayment, err := h.paymentRepository.UpdateOne(context.Background(), uint(paymentId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.UpdateFailed()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.UpdateSuccessfully(), updatedPayment))
}

func (h *PaymentHandler) DeleteOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed()))
		return
	}

	err = h.paymentRepository.DeleteOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.DeleteFailed()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.DeleteSuccessfully(), nil))
}
