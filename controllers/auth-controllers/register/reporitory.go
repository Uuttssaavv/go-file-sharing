package register

import (
	"go-crud/models"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	RegisterRepository(input *models.UserEntity) (*models.UserEntity, int)
}

type repository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *models.UserEntity) (*models.UserEntity, int) {

	var user models.UserEntity
	db := r.db

	//  check if user exists
	checkUserAccount := db.Select("*").Where("Email=?", input.Email).Find(&user)

	if checkUserAccount.RowsAffected > 0 {

		return nil, http.StatusConflict
	}

	//  if not then create the user into db

	db.NewRecord(input)
	createUser := db.Create(&input)
	
	if createUser.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}
