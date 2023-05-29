package filehandlers

import (
	"fmt"
	"go-crud/models"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAllFilesHandler(context *gin.Context) {
	jwtData, _ := context.Get("user")

	var user models.UserEntity
	// convert header to user enitity
	errors := utils.StringToEntity(jwtData, &user)

	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, http.MethodPost, nil)
		return
	}
	userId := user.ID
	fmt.Println(userId)
	fileResponse, statusCode := h.service.GetAllFiles(userId)

	fmt.Println(fileResponse)
	fmt.Println(statusCode)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Received files", http.StatusOK, http.MethodPost, fileResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, http.MethodPost, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusConflict, http.MethodPost, nil)
		return
	}
}
