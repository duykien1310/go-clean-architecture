package util

import (
	"auth_service/api/presenter"
	"auth_service/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerExceptionWithCustomStatus(ctx *gin.Context, statusCode int, status config.CustomStatus, err error) {
	errorMessage := &presenter.BasicResponse{
		Status:  string(status),
		Message: ParseError(ctx, err),
	}

	ctx.AbortWithStatusJSON(statusCode, errorMessage)
}
func HandlerException(ctx *gin.Context, statusCode int, err error) {
	errorMessage := &presenter.BasicResponse{
		Status:  fmt.Sprint(statusCode),
		Message: ParseError(ctx, err),
	}

	ctx.AbortWithStatusJSON(statusCode, errorMessage)
}
