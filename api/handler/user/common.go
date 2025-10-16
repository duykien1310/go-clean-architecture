package handler

import (
	userPresenter "auth_service/api/presenter/user"
	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/user"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserDetail(ctx *gin.Context, userService user.UseCase) {
	userId := ctx.Query("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}
	result, err := userService.GetUserDetail(userIdInt)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := userPresenter.UserDetailResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: convertUserDetailToPresenter(result),
	}

	ctx.JSON(http.StatusOK, response)
}
