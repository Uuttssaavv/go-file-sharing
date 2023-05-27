package filecontrollers

import "go-crud/models"

type Repository interface {
	CreateFile(file *models.FileModel) (models.FileModel, int)

	GetAllFiles() ([]models.FileModel, int)

	DeleteFile(fileID uint) int
}
