package auth

import (
	"auth_service/config"
	"auth_service/entity"
	"auth_service/infrastucture/repository/define"
	"auth_service/util"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Service struct {
	userRepo   UserRepository
	statusRepo StatusRepository
	roleRepo   RoleRepository
	verifier   Verifier
	emailRepo  EmailRepository
}

func NewService(userRepo UserRepository,
	statusRepo StatusRepository,
	roleRepo RoleRepository,
	verifier Verifier,
	emailRepo EmailRepository) *Service {
	return &Service{
		userRepo:   userRepo,
		statusRepo: statusRepo,
		roleRepo:   roleRepo,
		verifier:   verifier,
		emailRepo:  emailRepo,
	}
}

func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.userRepo = s.userRepo.WithTrx(trxHandle)
	s.statusRepo = s.statusRepo.WithTrx(trxHandle)
	s.roleRepo = s.roleRepo.WithTrx(trxHandle)

	return s
}

func (s Service) WithRedisTrx(trxHandle redis.Pipeliner) Service {
	s.verifier = s.verifier.WithTrx(trxHandle)

	return s
}

func (s Service) Register(user *entity.User, otpCode string) error {
	// verify information
	isEmailExist, err := s.userRepo.VerifyEmailExist(user.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return entity.ErrEmailAlreadyExist
	}

	isUserNameExist, err := s.userRepo.VerifyUserNameExist(user.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return entity.ErrUserNameAlreadyExist
	}

	// Get user OTP code
	userOtpCode, err := s.verifier.GetOTPRegisterCode(user.Email)
	if err != nil {
		return entity.InvalidOTPCode
	}

	if otpCode != userOtpCode {
		return entity.InvalidOTPCode
	}

	// Invalidate OTP code of user
	err = s.verifier.InvalidateOTPRegisterCode(user.Email, config.GetInt("jwt.accessMaxAge"))
	if err != nil {
		return err
	}

	err = user.HashPassword()
	if err != nil {
		return err
	}

	// Get default role
	defaultRole, err := s.roleRepo.GetRoleByCode(define.USER)
	if err != nil {
		return err
	}
	user.RoleId = defaultRole.Id

	// Get default status
	defaultStatus, err := s.statusRepo.GetStatusByCode(define.USER_ACTIVE)
	if err != nil {
		return err
	}
	user.StatusId = defaultStatus.Id

	// Register user
	err = s.userRepo.Register(user)
	if err != nil {
		return entity.ErrInternalServerError
	}

	return nil
}

func (s Service) SendOTPRegister(email string) error {
	// verify information
	isEmailExist, err := s.userRepo.VerifyEmailExist(email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return entity.ErrEmailAlreadyExist
	}

	otpCode, err := util.GenerateOTPCode(config.GetInt("jwt.otpLength"))
	if err != nil {
		return err
	}

	// Cache user OTP code
	err = s.verifier.CacheRegisterOTPCode(email, otpCode, config.GetInt("jwt.otpMaxAge"))
	if err != nil {
		return err
	}

	// Send email
	err = s.emailRepo.SendOTPForRegister(email, otpCode)
	if err != nil {
		return err
	}

	return nil
}
