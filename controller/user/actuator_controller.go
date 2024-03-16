package user

import (
	"net/http"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/facade/user"
	"github.com/gin-gonic/gin"
)

type ActuatorEndpoint struct {
	config *app.Config
}

func NewActuatorEndpoint(config *app.Config) *ActuatorEndpoint {
	return &ActuatorEndpoint{
		config: config,
	}
}

func (ep *ActuatorEndpoint) Actuator(c *gin.Context) {
	facade := user.NewActuatorFacade(ep.config)
	response := facade.ActuatorLogic()
	c.JSON(http.StatusOK, response)
}
