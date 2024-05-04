package entity

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Jti    string `json:"jti"`
}

type TokenPair struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"userId"`
	Jti    string `json:"jti"`
}

type ForgotPasswordTokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"userId"`
	Email  string `json:"email"`
	Jti    string `json:"jti"`
}
