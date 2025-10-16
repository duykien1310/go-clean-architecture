package handler

import (
	followingPayload "auth_service/api/payload/following"
	"auth_service/api/presenter"
	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/following"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func followUser(ctx *gin.Context, followService following.UseCase) {
	// Get payload
	var payload followingPayload.FollowingCreate
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	err = followService.FollowUser(payload.UserId, payload.FollowUserId)
	if err != nil {
		util.HandlerException(ctx, http.StatusInternalServerError, err)
		return
	}
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: nil,
	}
	ctx.JSON(http.StatusOK, response)
}
