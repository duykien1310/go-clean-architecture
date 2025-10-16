package following

import "auth_service/entity"

type Service struct {
	followingRepo FollowingRepository
}

func NewService(followingRepo FollowingRepository) *Service {
	return &Service{
		followingRepo: followingRepo,
	}
}

func (s Service) FollowUser(userId, followUserId int) error {
	following := &entity.Following{
		UserId:       userId,
		FollowUserId: followUserId,
	}

	return s.followingRepo.FollowUser(following)
}
