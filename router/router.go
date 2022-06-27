package router

import (
	"github.com/joseenriqe97/golang-init/controller"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	rootEndpoint := e.Group("api")
	OrganizationRouter(rootEndpoint, c)
	return e
}
