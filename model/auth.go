package model

import (
	"github.com/dgrijalva/jwt-go"
)

type HealthCheck struct {
	Status string `json:"status"`
}

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
}

type AccessTokenClaims struct {
	SessionId string `json:"session_id"`
	Id        int    `json:"id"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	SessionId string `json:"session_id"`
	jwt.StandardClaims
}

type TokenClaims interface {
	Valid() error
}
