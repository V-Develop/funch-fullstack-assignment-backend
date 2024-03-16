package handler

import (
	"net/http"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/auth"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/gin-gonic/gin"
)

type Validator struct {
	config *app.Config
}

func NewValidator(config *app.Config) *Validator {
	return &Validator{
		config: config,
	}
}

func (validator *Validator) GetUserPermit(c *gin.Context) {
	loger := util.NewLog("0", "LOG")
	middleware := NewMiddleware(validator.config)
	err := middleware.ValidateRequestHeader(c)
	if err != nil {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(validator.config.ErrorMessage.AuthError.InvalidAccess), "0")
		errorResponse := util.CreateResponseWithoutPayload(
			401,
			validator.config.ErrorMessage.AuthError.InvalidAccess.Code,
			validator.config.ErrorMessage.AuthError.InvalidAccess.Message,
		)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	}
	authService := auth.NewLoginFacade(validator.config, loger)
	authService.GetClaimCurrent(c)
	role := authService.GetRoleInClaim()
	if role != "user" {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(validator.config.ErrorMessage.AuthError.InvalidAccess), "0")
		errorResponse := util.CreateResponseWithoutPayload(
			401,
			validator.config.ErrorMessage.AuthError.InvalidAccess.Code,
			validator.config.ErrorMessage.AuthError.InvalidAccess.Message,
		)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	}
}
