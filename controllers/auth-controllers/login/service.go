package loginAuth


import (
	 "go-crud/models"
)

type Service interface {
	LoginService(input *LoginInput) (*models.UserEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *LoginInput)(*models.UserEntity, string) {
	// TODO: concrete implementation
	return nil,""
}