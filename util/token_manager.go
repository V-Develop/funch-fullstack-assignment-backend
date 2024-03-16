package util

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"github.com/gin-gonic/gin"
)

type TokenManager struct {
	accestokenClaims   *model.AccessTokenClaims
	refreshTokenClaims   *model.RefreshTokenClaims
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

func (tokenManager *TokenManager) GetAccessTokenClaims(c *gin.Context) {
	if claims, ok := c.Get("claim"); ok {
		tokenManager.accestokenClaims = claims.(*model.AccessTokenClaims)
	}
}

func (tokenManager *TokenManager) GetUserIdInClaim() int {
	return tokenManager.accestokenClaims.Id
}

func (tokenManager *TokenManager) GetRoleInClaim() string {
	return tokenManager.accestokenClaims.Role
}

func (tokenManager *TokenManager) GetSessionIdInClaim() string {
	return tokenManager.refreshTokenClaims.SessionId
}
