package handler

import (
	"errors"
	"library_api/contract"
	"library_api/pkg/errs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller interface {
	getPrefix() string
	initService(service *contract.Service)
	initRoute(app *gin.RouterGroup)
}

func New(app *gin.Engine, service *contract.Service) {
	allController := []controller{
		&authController{},
		&bookController{},
		&lendingController{},
	}

	for _, c := range allController {
		c.initService(service)
		group := app.Group(c.getPrefix())
		c.initRoute(group)
		log.Printf("initiate route %s\n", c.getPrefix())
	}
}

func handlerError(ctx *gin.Context, err error) {
	var messageErr errs.MessageErr
	if errors.As(err, &messageErr) {
		ctx.JSON(messageErr.Status(), messageErr)
	} else {
		ctx.Error(err).SetType(gin.ErrorTypePrivate)
		ctx.JSON(http.StatusInternalServerError, errs.ErrServer)
	}
}
