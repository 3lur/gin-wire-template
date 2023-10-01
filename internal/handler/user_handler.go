package handler

import (
	"gin-mini-mall/internal/pkg/request"
	"gin-mini-mall/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	service service.UserService
}

func (h *userHandler) Register(ctx *gin.Context) {
	var reqBody request.UserCreateRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err := h.service.CreateUser(reqBody.ToUser()); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "success")
}

// NewUserHandler return a new user handler
func NewUserHandler(handler *Handler, service service.UserService) UserHandler {
	return &userHandler{handler, service}
}
