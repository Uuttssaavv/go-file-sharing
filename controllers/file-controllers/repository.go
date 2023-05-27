package filecontrollers

import (
	"go-crud/models"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateFile(file *models.FileModel) (*models.FileModel, int)

	GetAllFiles() ([]models.FileModel, int)

	DeleteFile(fileID uint) int
}

type repository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateFile(file *models.FileModel) (*models.FileModel, int) {

	return nil, 0
}

func (repo *repository) GetAllFiles() ([]models.FileModel, int) {

	return nil, 0
}

func (repo *repository) DeleteFile(fileID uint) int {

	return 0
}
