package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseCancellationName = "cancellation"
var responseCancellation helper.ResponseMessage

type CancellationHandler struct {
	repository entity.CancellationRepository
}

func NewCancellationHandler(repo entity.CancellationRepository) *CancellationHandler {
	return &CancellationHandler{
		repository: repo,
	}
}

func (h *CancellationHandler) GetMany(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	cancellations, err := h.repository.GetMany(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.GetFailed(responseCancellationName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseCancellation.GetSuccessfully(responseCancellationName), cancellations))
}

func (h *CancellationHandler) GetManyByBookingID(ctx *gin.Context) {
	bookingId, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseCancellation.IdFailed(responseCancellationName)))
		return
	}

	cancellations, err := h.repository.GetManyByBookingID(context.Background(), uint(bookingId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.GetFailed(responseCancellationName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseCancellation.GetSuccessfully(responseCancellationName), cancellations))
}

func (h *CancellationHandler) GetOne(ctx *gin.Context) {
	cancellationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseCancellation.IdFailed(responseCancellationName)))
		return
	}

	cancellation, err := h.repository.GetOne(context.Background(), uint(cancellationId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.GetFailed(responseCancellationName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseCancellation.GetSuccessfully(responseCancellationName), cancellation))
}

func (h *CancellationHandler) CreateOne(ctx *gin.Context) {
	var cancellation entity.Cancellation
	if err := ctx.ShouldBindJSON(&cancellation); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseCancellation.RequestFailed(responseCancellationName)))
		return
	}

	createdCancellation, err := h.repository.CreateOne(context.Background(), &cancellation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.CreateFailed(responseCancellationName)))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessResponse(responseCancellation.CreateSuccessfully(responseCancellationName), createdCancellation))
}

func (h *CancellationHandler) DeleteOne(ctx *gin.Context) {
	cancellationId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseCancellation.IdFailed(responseCancellationName)))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(cancellationId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.DeleteFailed(responseCancellationName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseCancellation.DeleteSuccessfully(responseCancellationName), nil))
}

func (h *CancellationHandler) DeleteMany(ctx *gin.Context) {
	if err := h.repository.DeleteMany(context.Background()); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseCancellation.DeleteAllFailed(responseCancellationName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseCancellation.DeleteAllSuccessfully(responseCancellationName), nil))
}
