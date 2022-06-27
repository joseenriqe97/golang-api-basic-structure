package router

import (
	"github.com/joseenriqe97/golang-init/controller"
	"github.com/labstack/echo"
)

func OrganizationRouter(e *echo.Group, c controller.AppController) {
	e.GET("/organizations", func(context echo.Context) error { return c.User.GetUser(context) })
}
