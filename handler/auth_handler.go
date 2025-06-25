package handler

import (
	"library_api/contract"
	"library_api/dto"
	"library_api/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service contract.AuthService
}

func (a *authController) getPrefix() string {
	return "/auth"
}

func (a *authController) initService(service *contract.Service) {
	a.service = service.Auth
}

func (a *authController) initRoute(app *gin.RouterGroup) {
	app.POST("/register", a.Register)
	app.POST("/login", a.Login)
}

func (a *authController) Register(ctx *gin.Context) {
	var payload dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrRequestBody)
		return
	}

	result, err := a.service.Register(ctx, &payload)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (a *authController) Login(ctx *gin.Context) {
	var payload dto.LoginRequest

	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.ErrRequestBody)
		return
	}

	result, err := a.service.Login(ctx, &payload)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
