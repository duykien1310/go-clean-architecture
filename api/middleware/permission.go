package middleware

import (
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionAllowMiddleware(roleCode []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetToken(ctx)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		claims, err := util.ParseAccessToken(token)
		if err != nil {
			util.HandlerException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
			return
		}

		if util.InArray(claims.Role, roleCode) {
			ctx.Next()
			return
		} else {
			util.HandlerException(ctx, http.StatusForbidden, entity.ErrForbidden)
			return
		}
	}
}
