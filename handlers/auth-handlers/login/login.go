package loginHandler

import (
	"go-crud/controllers/auth-controllers/login"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	// TODO: concrete implementation of handler
}