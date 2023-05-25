package routes

import (
	"go-crud/controllers/auth-controllers/login"
	"go-crud/handlers/auth-handlers/login"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := loginHandler.NewHandlerLogin(loginService)

	route.POST("/login", loginHandler.LoginHandler)
}
