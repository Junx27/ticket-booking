package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseRefundName = "refund"
var responseRefund helper.ResponseMessage

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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.GetFailed(responseRefundName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseRefund.GetSuccessfully(responseRefundName), refunds))
}

func (h *RefundHandler) GetOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.IdFailed(responseRefundName)))
		return
	}

	refund, err := h.repository.GetOne(context.Background(), uint(refundId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.GetFailed(responseRefundName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseRefund.GetSuccessfully(responseRefundName), refund))
}

func (h *RefundHandler) CreateOne(ctx *gin.Context) {
	var refund entity.Refund
	if err := ctx.ShouldBindJSON(&refund); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseRefund.RequestFailed(responseRefundName)))
		return
	}

	createRefund, err := h.repository.CreateOne(context.Background(), &refund)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.CreateFailed(responseRefundName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseRefund.CreateSuccessfully(responseRefundName), createRefund))
}

func (h *RefundHandler) UpdateOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseRefund.IdFailed(responseRefundName)))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseRefund.RequestFailed(responseRefundName)))
		return
	}

	updateRefund, err := h.repository.UpdateOne(context.Background(), uint(refundId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.UpdateFailed(responseRefundName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseRefund.UpdateSuccessfully(responseRefundName), updateRefund))

}

func (h *RefundHandler) DeleteOne(ctx *gin.Context) {
	refundId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseRefund.IdFailed(responseRefundName)))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(refundId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseRefund.DeleteFailed(responseRefundName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseRefund.DeleteSuccessfully(responseRefundName), nil))
}
