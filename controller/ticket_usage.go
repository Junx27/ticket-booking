package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseTicketUsageName = "ticket usage"
var responseTicketUsage helper.ResponseMessage

type TicketUsageHanlder struct {
	repository entity.TicketUsageRepository
}

func NewTicketUsageHandler(repo entity.TicketUsageRepository) *TicketUsageHanlder {
	return &TicketUsageHanlder{
		repository: repo,
	}
}

func (h *TicketUsageHanlder) GetMany(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ticketUsages, err := h.repository.GetMany(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.GetFailed(responseTicketUsageName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseTicketUsage.GetSuccessfully(responseTicketUsageName), ticketUsages))
}

func (h *TicketUsageHanlder) GetOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.GetFailed(responseTicketUsageName)))
		return
	}

	ticketUsage, err := h.repository.GetOne(context.Background(), uint(ticketUsageId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.GetFailed(responseTicketUsageName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseTicketUsage.GetSuccessfully(responseTicketUsageName), ticketUsage))
}

func (h *TicketUsageHanlder) CreateOne(ctx *gin.Context) {
	var ticketUsage entity.TicketUsage
	if err := ctx.ShouldBindJSON(&ticketUsage); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseTicketUsage.RequestFailed(responseTicketUsageName)))
		return
	}

	createTicketUsage, err := h.repository.CreateOne(context.Background(), &ticketUsage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.GetFailed(responseTicketUsageName)))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseTicketUsage.CreateSuccessfully(responseTicketUsageName), createTicketUsage))
}

func (h *TicketUsageHanlder) UpdateOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseTicketUsage.IdFailed(responseTicketUsageName)))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseTicketUsage.GetFailed(responseTicketUsageName)))
		return
	}

	updateRefund, err := h.repository.UpdateOne(context.Background(), uint(ticketUsageId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.UpdateFailed(responseTicketUsageName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseTicketUsage.UpdateSuccessfully(responseTicketUsageName), updateRefund))

}

func (h *TicketUsageHanlder) DeleteOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseTicketUsage.IdFailed(responseTicketUsageName)))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(ticketUsageId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseTicketUsage.DeleteFailed(responseTicketUsageName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseTicketUsage.DeleteSuccessfully(responseTicketUsageName), nil))
}
