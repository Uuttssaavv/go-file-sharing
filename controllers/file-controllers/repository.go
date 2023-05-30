package filecontrollers

import (
	"fmt"
	"go-crud/models"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateFile(input *models.FileModel) (*models.FileModel, int)

	GetAllFiles(userId uint) ([]models.FileModel, int)

	DeleteFile(fileID uint) int
}

type repository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateFile(input *models.FileModel) (*models.FileModel, int) {

	db := repo.db

	var file models.FileModel

	checkIfFileExists := db.Select("*").Where("ID=?", input.ID).Find(&file)

	db.Select("*").Where("ID=?", input.UserID).Find(&input.User)

	if checkIfFileExists.RowsAffected > 0 {
		return nil, http.StatusConflict
	}
	db.NewRecord(input)
	createFile := db.Create(&input)

	if createFile.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}

func (repo *repository) GetAllFiles(userId uint) ([]models.FileModel, int) {
	var files []models.FileModel
	db := repo.db

	checkIfFileExists := db.Select("*").Where("user_id=?", userId).Find(&files)
	
	if checkIfFileExists.Error != nil {
		return nil, http.StatusNotFound
	}
	fmt.Printf("%+v",files)
	return files, http.StatusOK
}

func (repo *repository) DeleteFile(fileID uint) int {

	return 0
}
