package filehandlers

import (
	"fmt"
	filecontrollers "go-crud/controllers/file-controllers"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateHandler(context *gin.Context) {
	file, header, _ := context.Request.FormFile("file")

	result, err := utils.UploadFile(file, header.Header.Get("Content-Type"))

	if err != nil {
		fmt.Println(err)
	}
	
	fileInput := filecontrollers.FileInput{
		ID:   result.PublicID,
		Type: result.Format,
		Name: header.Filename,
	}

	fileResponse, statusCode := h.service.CreateFile(&fileInput)

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Uploaded the file successfully.", http.StatusCreated, http.MethodPost, fileResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, http.MethodPost, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusCreated, http.MethodPost, nil)
		return
	}

}
