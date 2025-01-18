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

var responsePaymentName = "payment"
var responsePayment helper.ResponseMessage

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
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	payments, err := h.paymentRepository.GetMany(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed(responsePaymentName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.GetSuccessfully(responsePaymentName), payments))
}

func (h *PaymentHandler) GetManyProvider(ctx *gin.Context) {
	payments, err := h.paymentRepository.GetManyProvider(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed(responsePaymentName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.GetSuccessfully(responsePaymentName), payments))
}

func (h *PaymentHandler) GetOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed(responsePaymentName)))
		return
	}

	payment, err := h.paymentRepository.GetOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed(responsePaymentName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.GetSuccessfully(responsePaymentName), payment))
}

func (h *PaymentHandler) CreateOne(ctx *gin.Context) {
	var payment entity.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.RequestFailed(responsePaymentName)))
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

	booking, err := h.bookingRepository.GetOne(context.Background(), payment.BookingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed("booking")))
		return
	}
	payment.UserID = userID
	payment.PaymentAmount = booking.TotalAmount
	createdPayment, err := h.paymentRepository.CreateOne(context.Background(), &payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.CreateFailed(responsePaymentName)))
		return
	}

	ticketUsage := entity.TicketUsage{
		BookingID: payment.BookingID,
		UserID:    payment.UserID,
	}
	_, err = h.ticketUsageRepositiry.CreateOne(context.Background(), &ticketUsage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.CreateFailed("ticket usage")))
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: fmt.Sprintf("Booking id %d status %s payment successfully", booking.ID, booking.BookingStatus),
	}

	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.CreateFailed("activity log")))
		return
	}
	notification := entity.Notification{
		UserID:  booking.UserID,
		Message: fmt.Sprintf("Booking id %d status %s payment completed", booking.ID, booking.BookingStatus),
	}

	_, err = h.notificationRepository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.CreateFailed("notification")))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse(responsePayment.CreateSuccessfully(responsePaymentName), createdPayment))
}

func (h *PaymentHandler) UpdateOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed(responsePaymentName)))
		return
	}

	payment, err := h.paymentRepository.GetOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.GetFailed(responsePaymentName)))
		return
	}

	var updateData entity.Payment
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.RequestFailed(responsePaymentName)))
		return
	}

	updateFields := map[string]interface{}{
		"id":             payment.ID,
		"booking_id":     payment.BookingID,
		"user_id":        payment.UserID,
		"payment_amount": payment.PaymentAmount,
		"payment_status": updateData.PaymentStatus,
		"payment_method": payment.PaymentMethod,
	}

	updatedPayment, err := h.paymentRepository.UpdateOne(context.Background(), uint(paymentId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.UpdateFailed(responsePaymentName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.UpdateSuccessfully(responsePaymentName), updatedPayment))
}

func (h *PaymentHandler) DeleteOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responsePayment.IdFailed(responsePaymentName)))
		return
	}

	err = h.paymentRepository.DeleteOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responsePayment.DeleteFailed(responsePaymentName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responsePayment.DeleteSuccessfully(responsePaymentName), nil))
}
