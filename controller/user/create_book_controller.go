package user

import (
	"net/http"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/auth"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/user"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/gin-gonic/gin"
)

type CreateBookEndpoint struct {
	config *app.Config
}

func NewCreateBookEndpoint(config *app.Config) *CreateBookEndpoint {
	return &CreateBookEndpoint{
		config: config,
	}
}

func (ep *CreateBookEndpoint) CreateBook(c *gin.Context) {
	loger := util.NewLog("0", "LOG")
	var request model.CreateBookRequest
	if err := util.IsInvalidFormat(
		c, &request,
		ep.config.ErrorMessage.CommonError.BindJsonError.Code,
		ep.config.ErrorMessage.CommonError.BindJsonError.Message); err != (model.ResponseWithoutPayload{}) {
		loger.LogErrorf("ENDPOINT: %+v", util.ConvertObjectToString(ep.config.ErrorMessage.CommonError.BindJsonError), "0")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	authFacade := auth.NewLoginFacade(ep.config, loger)
	authFacade.GetClaimCurrent(c)
	userId := authFacade.GetUserIdInClaim()

	loger.LogInfof("ENDPOINT REQUEST: %+v", request, string(rune(userId)))
	facade := user.NewCreateBookFacade(ep.config, loger)
	response := facade.CreateBookLogic(request, userId)
	loger.LogInfof("ENDPOINT RESPONSE: %+v", response, string(rune(userId)))
	c.JSON(response.Status.HttpStatus, response)
}
