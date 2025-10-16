package handler

import (
	postPayload "auth_service/api/payload/post"
	"auth_service/api/presenter"
	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/post"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context, postService post.UseCase) {
	// Get payload
	var payload postPayload.PostCreate
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	postEntity := convertPayloadToPostEntity(payload)
	err = postService.CreatePost(postEntity)
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

func generateNewsfeed(ctx *gin.Context, postService post.UseCase) {
	userId := ctx.Query("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		util.HandlerException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	result, err := postService.GenerateNewsfeed(userIdInt)
	if err != nil {
		util.HandlerException(ctx, http.StatusInternalServerError, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: config.SUCCESS,
		Results: result,
	}

	ctx.JSON(http.StatusOK, response)
}
