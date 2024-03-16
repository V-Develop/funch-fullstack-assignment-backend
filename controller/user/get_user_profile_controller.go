package user

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/auth"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/user"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/gin-gonic/gin"
)

type GetUserProfileEndpoint struct {
	config *app.Config
}

func NewGetUserProfileEndpoint(config *app.Config) *GetUserProfileEndpoint {
	return &GetUserProfileEndpoint{
		config: config,
	}
}

func (ep *GetUserProfileEndpoint) GetUserProfile(c *gin.Context) {
	loger := util.NewLog("0", "LOG")

	authFacade := auth.NewLoginFacade(ep.config, loger)
	authFacade.GetClaimCurrent(c)
	userId := authFacade.GetUserIdInClaim()

	loger.LogInfof("ENDPOINT REQUEST: %+v", userId, string(rune(userId)))
	facade := user.NewGetUserProfileFacade(ep.config, loger)
	response := facade.GetUserProfileLogic(userId)
	loger.LogInfof("ENDPOINT RESPONSE: %+v", response, string(rune(userId)))
	c.JSON(response.Status.HttpStatus, response)
}
