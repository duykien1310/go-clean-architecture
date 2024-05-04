package util

import (
	"auth_service/config"
	"auth_service/entity"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func keyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New(config.UNAUTHORIZED)
		}
		return []byte(key), nil
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader(config.GetString("jwt.header"))
	if len(authHeader) <= len(config.GetString("jwt.schema"))+1 {
		return "", errors.New(config.UNAUTHORIZED)
	}

	tokenString := authHeader[len(config.GetString("jwt.schema"))+1:]

	return tokenString, nil
}

func ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, keyFunc(config.GetString("jwt.key")))
}

func ParseAccessToken(token string) (*entity.TokenClaims, error) {
	claims := entity.TokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc(config.GetString("jwt.key")))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func ParseUnverifiedAccessToken(accessToken string) (*entity.TokenClaims, error) {
	claims := entity.TokenClaims{}

	parser := jwt.NewParser()
	_, _, err := parser.ParseUnverified(accessToken, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
func GenerateAccessToken(user *entity.User) (string, error) {
	// Generate random
	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// Generate claims
	claims := entity.TokenClaims{
		UserId: user.Id,
		Email:  user.Email,
		Jti:    random.String(),
		Role:   user.Role.Code,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetInt("jwt.accessMaxAge")))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(user *entity.User) (string, error) {
	claims := entity.RefreshTokenClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetInt("jwt.refreshMaxAge")))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedRefreshToken, err := refreshToken.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}

	return signedRefreshToken, nil
}
