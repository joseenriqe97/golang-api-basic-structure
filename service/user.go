package service

import (
	"fmt"

	"github.com/joseenriqe97/golang-init/repository"
)

//Name may be changed
type UtilUserService struct {
	DBRepository repository.UserRepository
}

type InteractorUser interface {
	GetUser(ID string) error
}

func New() *UtilUserService {
	return &UtilUserService{
		DBRepository: repository.NewUserRepository(),
	}
}

func (util *UtilUserService) GetUser(idUser string) error {
	res, _ := util.DBRepository.GetUser("1")
	fmt.Print(res)
	return nil
}
