package auth

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	authRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
)

type ActuatorFacade struct {
	config     *app.Config
	repository *authRepository.AuthRepository
}

func NewActuatorFacade(config *app.Config) *ActuatorFacade {
	repository := authRepository.AuthRepository{}
	return &ActuatorFacade{
		config:     config,
		repository: repository.NewAuthRepository(config),
	}
}

func (srv *ActuatorFacade) ActuatorLogic() (response model.HealthCheck) {
	response = model.HealthCheck{
		Status: "UP",
	}
	return
}