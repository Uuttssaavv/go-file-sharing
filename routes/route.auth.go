package routes

import (
	"go-crud/controllers/auth-controllers/login"
	"go-crud/controllers/auth-controllers/register"
	"go-crud/handlers/auth-handlers/login"
	registerHandler "go-crud/handlers/auth-handlers/register"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := loginHandler.NewHandlerLogin(loginService)

	registerRepository := register.NewRegisterRepository(db)
	registerService := register.NewRegisterService(registerRepository)
	registerHandlers := registerHandler.NewHandlerRegister(registerService)


	route.POST("/login", loginHandler.LoginHandler)

	route.POST("/register", registerHandlers.RegisterHandler)
}
