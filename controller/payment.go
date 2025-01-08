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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch payments"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data payments successfully", payments))
}

func (h *PaymentHandler) GetOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid payment ID"))
		return
	}

	payment, err := h.paymentRepository.GetOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch payment"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data payment successfully", payment))
}

func (h *PaymentHandler) CreateOne(ctx *gin.Context) {
	var payment entity.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	booking, err := h.bookingRepository.GetOne(context.Background(), payment.BookingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch booking"))
		return
	}

	payment.PaymentAmount = booking.TotalAmount
	createdPayment, err := h.paymentRepository.CreateOne(context.Background(), &payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create payment"))
		return
	}

	ticketUsage := entity.TicketUsage{
		BookingID: payment.BookingID,
	}
	_, err = h.ticketUsageRepositiry.CreateOne(context.Background(), &ticketUsage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create ticket usage"))
		return
	}

	activityLog := entity.ActivityLog{
		UserID:      booking.UserID,
		Description: fmt.Sprintf("Booking id %d status %s payment successfully", booking.ID, booking.BookingStatus),
	}

	_, err = h.activityLogRepository.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create activity log"))
		return
	}
	notification := entity.Notification{
		UserID:  booking.UserID,
		Message: fmt.Sprintf("Booking id %d status %s payment completed", booking.ID, booking.BookingStatus),
	}

	_, err = h.notificationRepository.CreateOne(context.Background(), &notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create notification"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create data payment successfully", createdPayment))
}

func (h *PaymentHandler) UpdateOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid payment ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updatedPayment, err := h.paymentRepository.UpdateOne(context.Background(), uint(paymentId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update payment"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update data payment successfully", updatedPayment))
}

func (h *PaymentHandler) DeleteOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid payment ID"))
		return
	}

	err = h.paymentRepository.DeleteOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete payment"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data payment successfully", nil))
}
