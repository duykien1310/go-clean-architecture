package middleware

import (
	"net/http"

	"auth_service/entity"
	"auth_service/infrastucture/repository/define"
	"auth_service/infrastucture/repository/util"

	"github.com/gin-gonic/gin"
)

func JWTVerifyMiddleware(verifier util.Verifier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetToken(ctx)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		tokenVerified, userData, err := verifier.Verify(token)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		if userData.Status != string(define.USER_ACTIVE) {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}

		if tokenVerified {
			ctx.Next()
			return
		} else {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}
	}
}
