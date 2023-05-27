package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-crud/controllers/file-controllers"
	"go-crud/handlers/file-handlers"
)

func InitFileRoutes(db *gorm.DB, route *gin.RouterGroup) {

	fileRepository := filecontrollers.NewFileRepository(db)
	fileService := filecontrollers.NewFileService(fileRepository)
	fileHanlders := filehandlers.NewCreateHandler(fileService)

	route.POST("/create", fileHanlders.CreateHandler)
}
