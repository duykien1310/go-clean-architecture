package handler

import (
	"auth_service/usecase/auth"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, authService auth.UseCase) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.GET("/healthCheck", func(ctx *gin.Context) {
			ctx.JSON(200, nil)
		})

		authGroup.POST("/register", func(ctx *gin.Context) {
			register(ctx, authService)
		})
	}
}
