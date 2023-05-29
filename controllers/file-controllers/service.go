package filecontrollers

import "go-crud/models"

type Service interface {
	CreateFile(input *FileInput) (*models.FileModel, int)

	GetAllFiles() ([]models.FileModel, int)

	DeleteFile(fileID uint) int
}

type service struct {
	repository Repository
}

func NewFileService(r Repository) *service {
	return &service{repository: r}
}

func (s *service) CreateFile(input *FileInput) (*models.FileModel, int) {
	fileModel := models.FileModel{
		Type:      input.Type,
		Name:      input.Name,
		Url:       input.Url,
		AccessKey: input.ID,
		UserID:      input.UserId,
	}
	return s.repository.CreateFile(&fileModel)
}

func (s *service) GetAllFiles() ([]models.FileModel, int) {

	return s.repository.GetAllFiles()
}

func (s *service) DeleteFile(fileID uint) int {

	return s.repository.DeleteFile(fileID)
}
