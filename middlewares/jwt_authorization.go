package middlewares

import (
	"net/http"

	"go-crud/utils"

	"github.com/gin-gonic/gin"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {

		var errorResponse UnathorizatedError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Authorization is required for this endpoint"

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		token, err := utils.VerifyTokenHeader(ctx, "JWT_SECRET")

		errorResponse.Status = "Unathorizated"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Invalid Access token"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			ctx.Set("user", token.Claims)
			// return to next method if token is exist
			ctx.Next()
		}
	})
}
