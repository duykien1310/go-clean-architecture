package auth

import (
	"auth_service/entity"
)

type UserRepository interface {
	Register(user *entity.User) error
	VerifyUserNameExist(userName string) (bool, error)
}

type UseCase interface {
	Register(user *entity.User) error
}
