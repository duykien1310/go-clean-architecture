package handler

import (
	authPayload "auth_service/api/payload/auth"
	"auth_service/api/presenter"
	"fmt"

	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Register
// @Schemes
// @Description register
// @Tags auth
// @Param account body authPayload.Register true "Register"
// @Success 200 {object} presenter.BasicResponse
// @Failure 400 {object} presenter.Error400Response
// @Failure 404 {object} presenter.Error404Response
// @Failure 500 {object} presenter.Error500Response
// @Router /auth/register [post]
func register(ctx *gin.Context, authService auth.UseCase) {
	// Get payload
	var payload authPayload.Register
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	userEntity := convertRegisterPayloadToUserEntity(payload)
	err = authService.Register(userEntity)
	if err != nil {
		switch err {
		case entity.ErrUserNameAlreadyExist:
			util.HandlerExceptionWithCustomStatus(ctx, http.StatusBadRequest, config.INVALID_INFORMATION, err)
			return

		default:
			util.HandlerException(ctx, http.StatusInternalServerError, err)
			return
		}
	}
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: nil,
	}
	ctx.JSON(http.StatusOK, response)
}
