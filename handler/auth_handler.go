package handler

import (
	"library_api/contract"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service contract.AuthService
}

func (c *authController) getPrefix() string {
	return "/auth"
}

func (c *authController) initService(service *contract.Service) {
	c.service = service.Auth
}

func (c *authController) initRoute(app *gin.RouterGroup) {

}
