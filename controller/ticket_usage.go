package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type TicketUsageHanlder struct {
	repository entity.TicketUsageRepository
}

func NewTicketUsageHandler(repo entity.TicketUsageRepository) *TicketUsageHanlder {
	return &TicketUsageHanlder{
		repository: repo,
	}
}

func (h *TicketUsageHanlder) GetMany(ctx *gin.Context) {
	ticketUsages, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch ticket usages"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data ticket usages successfully", ticketUsages))
}

func (h *TicketUsageHanlder) GetOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch ticket usage"))
		return
	}

	ticketUsage, err := h.repository.GetOne(context.Background(), uint(ticketUsageId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch ticket usage"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data ticket usage succesfully", ticketUsage))
}

func (h *TicketUsageHanlder) CreateOne(ctx *gin.Context) {
	var ticketUsage entity.TicketUsage
	if err := ctx.ShouldBindJSON(&ticketUsage); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	createTicketUsage, err := h.repository.CreateOne(context.Background(), &ticketUsage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create ticket usage"))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Craate ticket usage successfully", createTicketUsage))
}

func (h *TicketUsageHanlder) UpdateOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid ticket usage ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}

	updateRefund, err := h.repository.UpdateOne(context.Background(), uint(ticketUsageId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update ticket usage"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update ticket usage successfully", updateRefund))

}

func (h *TicketUsageHanlder) DeleteOne(ctx *gin.Context) {
	ticketUsageId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid ticket usage ID"))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(ticketUsageId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete ticket usage"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete ticket usage successfully", nil))
}
