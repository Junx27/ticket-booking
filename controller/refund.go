package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type RefundHandler struct {
	repository entity.RefundRepository
}

func NewRefundHandler(repo entity.RefundRepository) *RefundHandler {
	return &RefundHandler{
		repository: repo,
	}
}

func (h *RefundHandler) GetMany(ctx *gin.Context) {
	refunds, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch refunds"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data refunds successfully", refunds))
}

func (h *RefundHandler) GetOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch refund"))
		return
	}

	refund, err := h.repository.GetOne(context.Background(), uint(refundId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch refund"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data refund succesfully", refund))
}

func (h *RefundHandler) CreateOne(ctx *gin.Context) {
	var refund entity.Refund
	if err := ctx.ShouldBindJSON(&refund); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request patload"))
		return
	}

	createRefund, err := h.repository.CreateOne(context.Background(), &refund)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create refund"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Craate refund successfully", createRefund))
}

func (h *RefundHandler) UpdateOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid refund ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updateRefund, err := h.repository.UpdateOne(context.Background(), uint(refundId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update refund"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update refund successfully", updateRefund))

}

func (h *RefundHandler) DeleteOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid refund ID"))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(refundId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete refund"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete refund successfully", nil))
}
