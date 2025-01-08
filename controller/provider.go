package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseProviderName = "provider"
var responseProvider helper.ResponseMessage

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
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.GetFailed(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.GetSuccessfully(responseProviderName), providers))
}

func (h *ProviderHandler) GetOne(ctx *gin.Context) {
	providerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.IdFailed(responseProviderName)))
		return
	}

	provider, err := h.repository.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.GetFailed(responseProviderName)))
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
	createProvider, err := h.repository.CreateOne(context.Background(), &provider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.CreateFailed(responseProviderName)))
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

	_, err = h.repository.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse("Not found"))
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseProvider.RequestFailed(responseProviderName)))
		return
	}
	updateProvider, err := h.repository.UpdateOne(context.Background(), uint(providerId), updateData)
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

	_, err = h.repository.GetOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse("Not found"))
		return
	}

	err = h.repository.DeleteOne(context.Background(), uint(providerId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseProvider.DeleteFailed(responseProviderName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseProvider.DeleteSuccessfully(responseProviderName), nil))
}
