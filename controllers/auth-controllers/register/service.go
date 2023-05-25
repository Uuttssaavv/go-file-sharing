package register

import (
	"go-crud/models"
)

type Service interface {
	RegisterService(*RegisterInput) (*models.UserEntity, int)
}

type service struct {
	r Repository
}

func NewRegisterService(repository *repository) *service {
	return &service{r: repository}
}

func (s *service) RegisterService(input *RegisterInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
	}
	return s.r.RegisterRepository(&user)
}
