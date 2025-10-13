package handler

import (
	"auth_service/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, userService user.UseCase) {
	masterGroup := app.Group("/api/user")
	{
		masterGroup.GET("/detail", func(ctx *gin.Context) {
			getUserDetail(ctx, userService)
		})
	}
}
