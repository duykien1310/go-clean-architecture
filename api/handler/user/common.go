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

// @Summary Get user detail
// @Schemes
// @Description Get user detail
// @Tags user
// @Param userId query int true "User Id"
// @Success 200 {object} userPresenter.UserDetailResponse
// @Failure 400 {object} presenter.Error400Response
// @Failure 404 {object} presenter.Error404Response
// @Failure 500 {object} presenter.Error500Response
// @Router /user/detail [get]
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
