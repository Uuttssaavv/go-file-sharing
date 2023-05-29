package register

import (
	"go-crud/models"
)

type Service interface {
	RegisterService(*RegisterInput) (*models.UserEntity, int)
}

type service struct {
	repository Repository
}

func NewRegisterService(repository *repository) *service {
	return &service{repository: repository}
}

func (service *service) RegisterService(input *RegisterInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
		Image:    input.Image,
	}
	return service.repository.RegisterRepository(&user)
}
