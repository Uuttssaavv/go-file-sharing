package loginAuth

import (
	"github.com/jinzhu/gorm"
	"go-crud/models"
)

// The `Repository` interface defines a contract for the repository responsible for handling login-related operations
// it has only one function `LoginRepository` which takes the UserEntity and returns the UserEntity and string
type Repository interface {
	LoginRepository(input *models.UserEntity) (*models.UserEntity, string)
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
func (r *repository) LoginRepository(input *models.UserEntity) (*models.UserEntity, string) {
	// TODO: concrete implementation
	return nil, ""
}
