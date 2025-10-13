package user

import "auth_service/entity"

type UserRepository interface {
	FindById(id int) (*entity.User, error)
}

type UseCase interface {
	GetUserDetail(userId int) (*entity.User, error)
}
