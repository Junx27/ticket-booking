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

var responseProviderName = "provider"
var responseProvider helper.ResponseMessage

type ProviderHandler struct {
	repositoryProvider    entity.ProviderRepository
	repositoryActivityLog entity.ActivityLogRepository
}

func NewProviderHandler(repositoryProvider entity.ProviderRepository, repositoryActivityLog entity.ActivityLogRepository) *ProviderHandler {
	return &ProviderHandler{
		repositoryProvider:    repositoryProvider,
		repositoryActivityLog: repositoryActivityLog,
	}
}

func (h *ProviderHandler) GetMany(ctx *gin.Context) {
	providers, err := h.repositoryProvider.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.GetFailed(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.GetSuccessfully(responseProviderName), providers))
}

func (h *ProviderHandler) GetOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.GetFailed(responseProviderName)))
		return
	}

	provider, err := h.repositoryProvider.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseProvider.NotFound(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.GetSuccessfully(responseProviderName), provider))
}

func (h *ProviderHandler) CreateOne(ctx *gin.Context) {
	var provider entity.Provider
	if err := ctx.ShouldBindJSON(&provider); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.RequestFailed(responseProviderName)))
		return
	}
	createProvider, err := h.repositoryProvider.CreateOne(context.Background(), &provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.CreateFailed(responseProviderName)))
		return
	}
	activityLog := entity.ActivityLog{
		UserID:      provider.ID,
		Description: fmt.Sprintf("Provider create by user id %d successfully", provider.ID),
	}

	_, err = h.repositoryActivityLog.CreateOne(context.Background(), &activityLog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseBooking.CreateFailed("activity log")))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.CreateSuccessfully(responseProviderName), createProvider))
}

func (h *ProviderHandler) UpdateOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.IdFailed(responseProviderName)))
		return
	}

	provider, err := h.repositoryProvider.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseProvider.NotFound(responseProviderName)))
		return
	}

	var updateData entity.Provider
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.RequestFailed(responseProviderName)))
		return
	}
	updateFields := map[string]interface{}{
		"id":           provider.ID,
		"name":         updateData.Name,
		"description":  updateData.Description,
		"contact_info": updateData.ContactInfo,
	}
	updateProvider, err := h.repositoryProvider.UpdateOne(context.Background(), uint(providerId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.UpdateFailed(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.UpdateSuccessfully(responseProviderName), updateProvider))
}

func (h *ProviderHandler) DeleteOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.IdFailed(responseProviderName)))
		return
	}

	_, err = h.repositoryProvider.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse(responseProvider.NotFound(responseProviderName)))
		return
	}

	err = h.repositoryProvider.DeleteOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.DeleteFailed(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.DeleteSuccessfully(responseProviderName), nil))
}
