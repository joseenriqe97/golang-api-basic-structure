package controller

import (
	"fmt"

	"github.com/joseenriqe97/golang-init/service"
	"github.com/labstack/echo"
)

//APLICAR ESTO EN LOS SERVICE
/* type projectController struct {
	projectInteractor project.Interactor
} */

type userService struct {
	organizationInteractor service.InteractorUser
}

type UserController interface {
	GetUser(c echo.Context) error
}

var UserOrg UserController = &userService{organizationInteractor: service.New()}

func (orgController *userService) GetUser(c echo.Context) error {
	orgController.organizationInteractor.GetUser("2")
	fmt.Println("Estamos en el controlador")
	return nil
}
