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

	"github.com/go-redis/redis/v8"

	ginI18n "github.com/gin-contrib/i18n"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// Get transaction
	txHandle := ctx.MustGet("db_trx").(*gorm.DB)

	userEntity := convertRegisterPayloadToUserEntity(payload)
	err = authService.WithTrx(txHandle).Register(userEntity, payload.OtpCode)
	if err != nil {
		switch err {
		case entity.ErrEmailAlreadyExist:
			util.HandlerExceptionWithCustomStatus(ctx, http.StatusBadRequest, config.INVALID_INFORMATION, err)
			return

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
		Message: ginI18n.MustGetMessage(ctx, config.SUCCESS),
		Results: nil,
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Send OTP for register
// @Schemes
// @Description Send OTP for register
// @Tags auth
// @Success 200 {object} presenter.BasicResponse
// @Failure 400 {object} presenter.Error400Response
// @Failure 404 {object} presenter.Error404Response
// @Failure 500 {object} presenter.Error500Response
// @Router /auth/register/otp [post]
func sendOTPRegister(ctx *gin.Context, authService auth.UseCase) {
	// Get payload
	var payload authPayload.SendMailRegister
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Get redis transaction
	trxRedis := ctx.MustGet("redis_trx").(redis.Pipeliner)

	// Send OTP
	err = authService.WithRedisTrx(trxRedis).SendOTPRegister(payload.Email)
	if err != nil {
		switch err {
		case entity.ErrEmailAlreadyExist:
			util.HandlerExceptionWithCustomStatus(ctx, http.StatusBadRequest, config.INVALID_INFORMATION, err)
			return

		default:
			util.HandlerException(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	// Response in JSON
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(ctx, config.SUCCESS),
	}
	ctx.JSON(http.StatusOK, response)
}
