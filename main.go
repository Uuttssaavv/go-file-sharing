package main

import (
	"go-crud/configs"
	"go-crud/routes"
	"go-crud/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupAppRouter()
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupAppRouter() *gin.Engine {

	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	api := router.Group("api/v1")

	file := api.Group("/file")

	routes.InitAuthRoutes(db, api)

	routes.InitFileRoutes(db, file)

	return router
}
