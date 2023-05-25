package loginAuth

import (
	"go-crud/models"
	"go-crud/utils"
	"net/http"

	"github.com/jinzhu/gorm"
)

// The `Repository` interface defines a contract for the repository responsible for handling login-related operations
// it has only one function `LoginRepository` which takes the UserEntity and returns the UserEntity and int
type Repository interface {
	LoginRepository(input *models.UserEntity) (*models.UserEntity, int)
}

// The `repository` struct is the concrete implementation of the `Repository` interface.
type repository struct {
	db *gorm.DB
}

// the `NewRepositoryLogin` function is the constructor of the `repository` struct that
// creates the new instance of the `repository` struct
func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

// The LoginRepository method of the repository struct implements the LoginRepository method
// from the Repository interface.
func (r *repository) LoginRepository(input *models.UserEntity) (*models.UserEntity, int) {
	//  check if the user exist
	var users models.UserEntity
	db := r.db.Model(&users)
	checkAccount := db.Select("*").Where("email=?", input.Email).Find(&users)

	if checkAccount.RowsAffected==0{
		return nil, http.StatusNotFound
	}
	// check if the password matches
	comparePassword := utils.ComparePassword(users.Password, input.Password)
	
	if comparePassword !=nil{
		return nil, http.StatusUnauthorized
	}

	return &users,http.StatusAccepted
}
