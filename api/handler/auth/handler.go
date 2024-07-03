package handler

import (
	"auth_service/api/middleware"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/auth"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, authService auth.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.GET("/healthCheck", func(ctx *gin.Context) {
			ctx.JSON(200, 'Server A')
		})

		authGroup.POST("/register/otp", tx.RedisTransactionMiddleware(), func(ctx *gin.Context) {
			sendOTPRegister(ctx, authService)
		})

		authGroup.POST("/register", tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			register(ctx, authService)
		})
	}
}
