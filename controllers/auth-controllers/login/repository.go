package loginAuth

import (
	"go-crud/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	LoginRepository(input *models.UserEntity) (*models.UserEntity, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *models.UserEntity) (*models.UserEntity, string) {
	// TODO: concrete implementation
	return nil,""
}