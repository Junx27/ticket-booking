package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type CancellationHandler struct {
	repository entity.CancellationRepository
}

func NewCancellationHandler(repo entity.CancellationRepository) *CancellationHandler {
	return &CancellationHandler{
		repository: repo,
	}
}

func (h *CancellationHandler) GetMany(ctx *gin.Context) {
	cancellations, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch cancellations"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data cancellations successfully", cancellations))
}

func (h *CancellationHandler) GetManyByBookingID(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid booking ID"))
		return
	}

	cancellations, err := h.repository.GetManyByBookingID(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch cancellations"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data cancellations successfully", cancellations))
}

func (h *CancellationHandler) GetOne(ctx *gin.Context) {
	cancellationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid cancellation ID"))
		return
	}

	cancellation, err := h.repository.GetOne(context.Background(), uint(cancellationId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch cancellation"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data cancellation successfully", cancellation))
}

func (h *CancellationHandler) CreateOne(ctx *gin.Context) {
	var cancellation entity.Cancellation
	if err := ctx.ShouldBindJSON(&cancellation); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createdCancellation, err := h.repository.CreateOne(context.Background(), &cancellation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create cancellation"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse("Create data cancellation successfully", createdCancellation))
}

func (h *CancellationHandler) DeleteOne(ctx *gin.Context) {
	cancellationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid cancellation ID"))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(cancellationId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete cancellation"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete data cancellation successfully", nil))
}

func (h *CancellationHandler) DeleteMany(ctx *gin.Context) {
	if err := h.repository.DeleteMany(context.Background()); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete all cancellations"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Deleted all cancellations successfully", nil))
}