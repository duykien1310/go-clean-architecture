package handler

import (
	"auth_service/usecase/post"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, postService post.UseCase) {
	masterGroup := app.Group("/api/post")
	{
		masterGroup.POST("/create", func(ctx *gin.Context) {
			createPost(ctx, postService)
		})

		masterGroup.GET("/newsfeed/generate", func(ctx *gin.Context) {
			generateNewsfeed(ctx, postService)
		})
	}
}
