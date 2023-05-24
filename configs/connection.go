package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"go-crud/utils"
)

func Connection() *gorm.DB {
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

func databaseMigrations(db *gorm.DB){
	// 
	
	logrus.Info("Database migrations")

}