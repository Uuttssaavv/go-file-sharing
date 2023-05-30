package configs

import (
	"go-crud/models"
	"go-crud/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)
type DBConnection interface {
	Connection() *gorm.DB
}

type service struct{}

func NewDBService() *service{
	return &service{}
}

func (s *service) Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	databaseURI <- utils.GodotEnv("DATABASE_URL_DEV")

	db, err := gorm.Open("postgres", <-databaseURI)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	} else {
		logrus.Info("Connection to Database Successfully")
	}

	//  setup db migrations here
	databaseMigrations(db)

	return db
}

func databaseMigrations(db *gorm.DB) {
	//
	db.AutoMigrate(&models.UserEntity{}, &models.FileModel{})
	db.Model(&models.FileModel{}).AddForeignKey("user_id", "user_entities(id)", "RESTRICT", "RESTRICT")

	logrus.Info("Database migrations")

}


// Proxy function that can be replaced for testing
var Connection = func() *gorm.DB {
	svc := &service{}
	return svc.Connection()
}