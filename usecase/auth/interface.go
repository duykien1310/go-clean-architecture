package auth

import (
	"auth_service/entity"
	"auth_service/infrastucture/repository"
	"auth_service/infrastucture/repository/util"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Verifier interface {
	WithTrx(redis.Pipeliner) *util.JWTVerifier

	CacheRegisterOTPCode(email string, otpCode string, expiresAt int) error
	GetOTPRegisterCode(email string) (string, error)
	InvalidateOTPRegisterCode(email string, expiresAt int) error
}

type EmailRepository interface {
	SendOTPForRegister(email string, otpCode string) error
}

type RoleRepository interface {
	WithTrx(*gorm.DB) repository.RoleRepository

	GetRoleByCode(code string) (*entity.Role, error)
}

type StatusRepository interface {
	WithTrx(*gorm.DB) repository.StatusRepository

	GetStatusByCode(code string) (*entity.Status, error)
}

type UserRepository interface {
	WithTrx(*gorm.DB) repository.UserRepository

	Register(user *entity.User) error
	VerifyEmailExist(email string) (bool, error)
	VerifyUserNameExist(userName string) (bool, error)
}

type UseCase interface {
	WithTrx(*gorm.DB) Service
	WithRedisTrx(redis.Pipeliner) Service

	SendOTPRegister(accessToken string) error
	Register(user *entity.User, otpCode string) error
}
