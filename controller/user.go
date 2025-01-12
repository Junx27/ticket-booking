package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
)

var responseUserName = "user"
var responseUser helper.ResponseMessage

type UserHandler struct {
	repository entity.UserRepository
}

func NewUserHandler(repo entity.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repo,
	}
}

func (h *UserHandler) GetMany(ctx *gin.Context) {
	users, err := h.repository.GetMany(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.GetFailed(responseUserName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseUser.GetSuccessfully(responseUserName), users))
}

func (h *UserHandler) GetOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseUser.IdFailed(responseUserName)))
		return
	}

	user, err := h.repository.GetOne(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.GetFailed(responseUserName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseUser.GetSuccessfully(responseUserName), user))
}

func (h *UserHandler) UpdateOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseUser.IdFailed(responseUserName)))
		return
	}

	user, err := h.repository.GetOne(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.GetFailed(responseUserName)))
		return
	}

	var updateData entity.User
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseUser.RequestFailed(responseUserName)))
		return
	}

	updateFields := map[string]interface{}{
		"id":           user.ID,
		"email":        updateData.Email,
		"password":     updateData.Password,
		"first_name":   updateData.FirstName,
		"last_name":    updateData.LastName,
		"role":         user.Role,
		"phone_number": updateData.PhoneNumber,
	}

	updatedUser, err := h.repository.UpdateOne(context.Background(), uint(userId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.UpdateFailed(responseUserName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseUser.UpdateSuccessfully(responseUserName), updatedUser))
}

func (h *UserHandler) UpdateOneProvider(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseUser.IdFailed(responseUserName)))
		return
	}

	user, err := h.repository.GetOne(context.Background(), uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.GetFailed(responseUserName)))
		return
	}
	updateFields := map[string]interface{}{
		"id":           user.ID,
		"email":        user.Email,
		"password":     user.Password,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"role":         "provider",
		"phone_number": user.PhoneNumber,
	}

	updatedUser, err := h.repository.UpdateOne(context.Background(), uint(userId), updateFields)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.UpdateFailed(responseUserName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseUser.UpdateSuccessfully(responseUserName), updatedUser))
}

func (h *UserHandler) DeleteOne(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse(responseUser.IdFailed(responseUserName)))
		return
	}

	if err := h.repository.DeleteOne(context.Background(), uint(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse(responseUser.DeleteFailed(responseUserName)))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessResponse(responseUser.DeleteSuccessfully(responseUserName), nil))
}
