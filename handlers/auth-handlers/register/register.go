package register

import (
	"go-crud/controllers/auth-controllers/register"
	"go-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.RegisterInput
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Username",
				Message: "Username is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Username",
				Message: "Username must be using lowercase",
			},
			{
				Tag:     "required",
				Field:   "Email",
				Message: "Email is required on body",
			},
			{
				Tag:     "email",
				Field:   "Email",
				Message: "Email format is not valid",
			},
			{
				Tag:     "required",
				Field:   "Password",
				Message: "Password is required on body",
			},
			{
				Tag:     "gte",
				Field:   "Password",
				Message: "Password minimum must be 8 character",
			},
		},
	}

	errorResponse, errCount := utils.GoValidator(&input, config.Options)
	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errorResponse)
		return
	}
	registerResult, errorCode := h.service.RegisterService(&input)

	switch errorCode {
	case http.StatusCreated:
		accessTokenData := map[string]interface{}{"id": registerResult.ID, "email": registerResult.Email}
		accessToken, errToken := utils.Sign(accessTokenData, utils.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}
		//  parse the UserEntity into response json
		var data register.RegisterResponse

		utils.ObjectToJson(registerResult, &data)

		data.Token = accessToken

		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, data)
		return

	case http.StatusConflict:
		utils.APIResponse(ctx, "Email already taken", http.StatusConflict, http.MethodPost, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", http.StatusExpectationFailed, http.MethodPost, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodPost, nil)
	}

}
