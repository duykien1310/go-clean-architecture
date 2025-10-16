package handler

import (
	"auth_service/usecase/following"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, followingService following.UseCase) {
	masterGroup := app.Group("/api/following")
	{
		masterGroup.POST("/create", func(ctx *gin.Context) {
			followUser(ctx, followingService)
		})
	}
}
