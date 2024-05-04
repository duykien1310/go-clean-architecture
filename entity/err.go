package entity

import (
	"auth_service/config"
	"errors"

	"github.com/lib/pq"
)

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

var ErrLock = errors.New(config.LOCK)
var ErrUnauthorized = errors.New(config.UNAUTHORIZED)
var ErrInternalServerError = errors.New(config.INTERNAL_SERVER_ERROR)
var ErrBadRequest = errors.New(config.BAD_REQUEST)
var ErrAlreadyApprove = errors.New(config.BAD_REQUEST)
var ErrUsernameNotExists = errors.New(config.USERNAME_NOT_EXISTS)
var ErrUserNotExists = errors.New(config.USER_NOT_EXISTS)
var ErrUsernameOrEmailNotExists = errors.New(config.USERNAME_OR_EMAIL_NOT_EXISTS)
var ErrInvalidPassword = errors.New(config.INVALID_PASSWORD)
var ErrForbidden = errors.New(config.FORBIDDEN)
var ErrEmailNotExists = errors.New(config.EMAIL_NOT_EXISTS)
var InvalidOTPCode = errors.New(config.INVALID_OTP_CODE)
var ErrInvalidDateFormat = errors.New(config.INVALID_DATE_FORMAT)
var ErrDuplicateUsername = errors.New(config.DUPLICATE_USERNAME)
var ErrInvalidRole = errors.New(config.INVALID_ROLE)
var ErrUserDuplicate = errors.New(config.USER_DUPLICATE)
var ErrEmailAlreadyExist = errors.New(config.USER_EXISTS)
var ErrUserNameAlreadyExist = errors.New(config.USER_NAME_EXISTS)
var ErrSameOldNewPassword = errors.New(config.SAME_OLD_NEW_PASSWORD)
var ErrInvalidEmail = errors.New(config.INVALID_EMAIL)
