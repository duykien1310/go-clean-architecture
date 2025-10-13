package handler

import (
	payload "auth_service/api/payload/auth"
	"auth_service/entity"
)

func convertRegisterPayloadToUserEntity(payload payload.Register) *entity.User {
	result := entity.User{
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  payload.Password,
	}

	return &result
}
