package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/configs"
	"go-crud/routes"
)

func main() {
	SetupAppRouter()
}

func SetupAppRouter() *gin.Engine {

	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	api := router.Group("api/v1")

	routes.InitAuthRoutes(db, api)

	return router
}
