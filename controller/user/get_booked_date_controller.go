package user

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/user"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/gin-gonic/gin"
)

type GetBookedDateEndpoint struct {
	config *app.Config
}

func NewGetBookedDateEndpoint(config *app.Config) *GetBookedDateEndpoint {
	return &GetBookedDateEndpoint{
		config: config,
	}
}

func (ep *GetBookedDateEndpoint) GetBookedDate(c *gin.Context) {
	loger := util.NewLog("0", "LOG")

	facade := user.NewGetBookedDateFacade(ep.config, loger)
	response := facade.GetBookedDateLogic()
	loger.LogInfof("ENDPOINT RESPONSE: %+v", response, "0")
	c.JSON(response.Status.HttpStatus, response)
}
