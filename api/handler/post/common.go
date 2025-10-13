package handler

import (
	postPayload "auth_service/api/payload/post"
	"auth_service/api/presenter"
	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/util"
	"auth_service/usecase/post"
	"fmt"

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
