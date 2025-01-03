package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	repository entity.ProviderRepository
}

func NewProviderHandler(repo entity.ProviderRepository) *ProviderHandler {
	return &ProviderHandler{
		repository: repo,
	}
}

func (h *ProviderHandler) GetMany(ctx *gin.Context) {
	providers, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch providers"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data providers successfully", providers))
}

func (h *ProviderHandler) GetOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid provider ID Not Found"))
		return
	}

	provider, err := h.repository.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch provider"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Fetch data provider successfully", provider))
}

func (h *ProviderHandler) CreateOne(ctx *gin.Context) {
	var provider entity.Provider
	if err := ctx.ShouldBindJSON(&provider); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}
	createProvider, err := h.repository.CreateOne(context.Background(), &provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create provider"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Create provider successfully", createProvider))
}

func (h *ProviderHandler) UpdateOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid provider ID"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid request payload"))
		return
	}
	updateProvider, err := h.repository.UpdateOne(context.Background(), uint(providerId), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update provider"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Update provider successfully", updateProvider))
}

func (h *ProviderHandler) DeleteOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid provider ID"))
		return
	}

	err = h.repository.DeleteOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete provider"))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse("Delete provider successfully", nil))
}
