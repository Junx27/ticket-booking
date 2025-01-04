package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	repository entity.PaymentRepository
}

func NewPaymentHandler(repo entity.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{
		repository: repo,
	}
}

func (h *PaymentHandler) GetMany(ctx *gin.Context) {
	payments, err := h.repository.GetMany(context.Background())
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

	payment, err := h.repository.GetOne(context.Background(), uint(paymentId))
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

	createdPayment, err := h.repository.CreateOne(context.Background(), &payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create payment"))
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

	updatedPayment, err := h.repository.UpdateOne(context.Background(), uint(paymentId), updateData)
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

	err = h.repository.DeleteOne(context.Background(), uint(paymentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete payment"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data payment successfully", nil))
}
