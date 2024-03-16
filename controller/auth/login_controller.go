package auth

import (
	"net/http"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/auth"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/gin-gonic/gin"
)

type LoginEndpoint struct {
	config *app.Config
}

func NewLoginEndpoint(config *app.Config) *LoginEndpoint {
	return &LoginEndpoint{
		config: config,
	}
}

func (ep *LoginEndpoint) Login(c *gin.Context) {
	loger := util.NewLog("0", "LOG")
	var request model.LoginRequest
	if err := util.IsInvalidFormat(
		c, &request,
		ep.config.ErrorMessage.CommonError.BindJsonError.Code,
		ep.config.ErrorMessage.CommonError.BindJsonError.Message); err != (model.ResponseWithoutPayload{}) {
		loger.LogErrorf("ENDPOINT: %+v", util.ConvertObjectToString(ep.config.ErrorMessage.CommonError.BindJsonError), "0")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	loger.LogInfof("ENDPOINT REQUEST: %+v", request, "0")
	facade := auth.NewLoginFacade(ep.config, loger)
	response := facade.LoginLogic(request)
	loger.LogInfof("ENDPOINT RESPONSE: %+v", response, "0")
	c.JSON(response.Status.HttpStatus, response)
}
