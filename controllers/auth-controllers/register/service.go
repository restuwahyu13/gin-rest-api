package register

import (
	"time"

	model "github.com/restuwahyu13/gin-rest-api/models"
	"github.com/restuwahyu13/gin-rest-api/utils"
)

type Service interface {
	RegisterService(input *InputRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(utils.HashPassword(input.Password)),
		CreatedAt: time.Now(),
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}