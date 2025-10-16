package following

import "auth_service/entity"

type FollowingRepository interface {
	FollowUser(following *entity.Following) error
}

type UseCase interface {
	FollowUser(userId, followUserId int) error
}
