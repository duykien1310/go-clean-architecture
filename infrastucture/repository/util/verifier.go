package util

import (
	"auth_service/entity"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserData struct {
	Status string
}

type Verifier interface {
	Verify(token string) (bool, *UserData, error)
	GetUserData(userId int) (*UserData, error)
	InvalidateToken(token string) error
	CacheUserData(user *entity.User, expiresAt int) error
}

type JWTVerifier struct {
	redis     *redis.Client
	pipeliner redis.Pipeliner
}

func NewVerifier(redis *redis.Client, pipeliner redis.Pipeliner) *JWTVerifier {
	return &JWTVerifier{
		redis:     redis,
		pipeliner: pipeliner,
	}
}

func (v *JWTVerifier) WithTrx(trxHandle redis.Pipeliner) *JWTVerifier {
	if trxHandle == nil {
		return v
	}

	v.pipeliner = trxHandle
	return v
}

func (v *JWTVerifier) Verify(token string) (bool, *UserData, error) {
	claims, err := ParseAccessToken(token)
	if err != nil {
		return false, nil, entity.ErrUnauthorized
	}

	value, err := v.redis.Get(context.TODO(), claims.Jti).Result()
	if err != redis.Nil || value == token {
		return false, nil, entity.ErrUnauthorized
	}

	userData, err := v.GetUserData(claims.UserId)
	if err != nil {
		return false, nil, entity.ErrUnauthorized
	}

	return true, userData, nil
}

func (v *JWTVerifier) GetUserData(userId int) (*UserData, error) {
	result, err := v.redis.Get(context.TODO(), fmt.Sprint(userId)).Result()
	if err != nil {
		return nil, entity.ErrUnauthorized
	}

	userData := UserData{}
	err = json.Unmarshal([]byte(result), &userData)
	if err != nil {
		return nil, entity.ErrLock
	}

	return &userData, nil
}

func (v *JWTVerifier) InvalidateToken(token string) error {
	claims, err := ParseAccessToken(token)
	if err != nil {
		return err
	}

	err = v.pipeliner.Set(context.TODO(), claims.Jti, token, time.Duration(time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()))).Err()
	if err != nil {
		return err
	}

	return nil
}

func (v *JWTVerifier) CacheUserData(user *entity.User, expiresAt int) error {
	userData := UserData{
		Status: user.Status.Code,
	}

	userDataJSON, err := json.Marshal(&userData)
	if err != nil {
		return err
	}

	err = v.pipeliner.Set(context.TODO(), fmt.Sprint(user.Id), userDataJSON, time.Duration(expiresAt)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (v *JWTVerifier) CacheRegisterOTPCode(email string, otpCode string, expiresAt int) error {
	err := v.pipeliner.Set(context.TODO(), fmt.Sprintf("otp_%v", email), otpCode, time.Duration(expiresAt)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (v *JWTVerifier) GetOTPRegisterCode(email string) (string, error) {
	otpCode, err := v.redis.Get(context.TODO(), fmt.Sprintf("otp_%v", email)).Result()
	if err != nil {
		return "", err
	}

	return otpCode, nil
}

func (v *JWTVerifier) InvalidateOTPRegisterCode(email string, expiresAt int) error {
	err := v.pipeliner.Set(context.TODO(), fmt.Sprintf("otp_%v", email), "", time.Duration(time.Unix(int64(expiresAt), 0).Sub(time.Now()))).Err()
	if err != nil {
		return err
	}

	return nil
}
