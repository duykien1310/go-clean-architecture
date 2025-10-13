package handler

import (
	payload "auth_service/api/payload/post"
	"auth_service/entity"
)

func convertPayloadToPostEntity(payload payload.PostCreate) *entity.Post {
	result := entity.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserId:  payload.UserId,
	}

	return &result
}
