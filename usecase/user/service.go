package user

import "auth_service/entity"

type Service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s Service) GetUserDetail(userId int) (*entity.User, error) {
	return s.userRepo.FindById(userId)
}
