package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-crud/controllers/file-controllers"
	"go-crud/handlers/file-handlers"
	"go-crud/middlewares"
)

func InitFileRoutes(db *gorm.DB, route *gin.RouterGroup) {

	fileRepository := filecontrollers.NewFileRepository(db)
	fileService := filecontrollers.NewFileService(fileRepository)
	fileHanlders := filehandlers.NewCreateHandler(fileService)

	// added auth middlewares
	route.Use(middlewares.Auth())

	route.POST("/create", fileHanlders.CreateHandler)

	route.GET("/", fileHanlders.GetAllFilesHandler)

	route.DELETE("/:fileId", fileHanlders.DeleteHandler)
}
