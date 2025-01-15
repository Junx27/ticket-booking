package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/Junx27/ticket-booking/entity"
	"github.com/Junx27/ticket-booking/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	service entity.AuthService
}

func NewAuthHandler(
	service entity.AuthService,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	creds := &entity.AuthCredentials{}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.LoginFailed(helper.ResponseMessage{}))
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.LoginFailedEntity(helper.ResponseMessage{}))
		return
	}

	token, _, err := h.service.Login(ctxTimeout, creds)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.LoginFailed(helper.ResponseMessage{}))
		return
	}
	ctx.SetCookie("token", token, 3600*24*1, "/", "", false, true)
	ctx.JSON(http.StatusOK, helper.AuthResponse(helper.ResponseMessage.LoginSuccessfully(helper.ResponseMessage{}), token))
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	creds := &entity.User{}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.RegisterFailed(helper.ResponseMessage{}))
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.RegisterFailedEntity(helper.ResponseMessage{}))
		return
	}

	token, _, err := h.service.Register(ctxTimeout, creds)
	ctx.SetCookie("token", token, 3600*24*1, "/", "", false, true)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseMessage.RegisterFailed(helper.ResponseMessage{}))
		return
	}

	ctx.JSON(http.StatusCreated, helper.AuthResponse(helper.ResponseMessage.RegisterSuccessfully(helper.ResponseMessage{}), token))
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, helper.SuccessResponse("Logout successfully", nil))
}
