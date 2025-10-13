package auth

import (
	"auth_service/entity"
)

type Service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (s Service) Register(user *entity.User) error {
	isUserNameExist, err := s.userRepo.VerifyUserNameExist(user.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return entity.ErrUserNameAlreadyExist
	}

	err = user.HashPassword()
	if err != nil {
		return err
	}

	// Register user
	err = s.userRepo.Register(user)
	if err != nil {
		return entity.ErrInternalServerError
	}

	return nil
}
