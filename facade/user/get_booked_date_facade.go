package user

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	userRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
)

type GetBookedDateFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *userRepository.UserRepository
}

func NewGetBookedDateFacade(config *app.Config, loger util.Loger) *GetBookedDateFacade {
	repository := userRepository.UserRepository{}
	return &GetBookedDateFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewUserRepository(config),
	}
}

func (srv *GetBookedDateFacade) GetBookedDateLogic() (response model.ResponseWithPayload) {
	bookedDate, err := srv.repository.GetBookedDate(srv.config.DataBase.Connection)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse := util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	response = util.CreateResponseWithPayload(
		200,
		"00000",
		"Get all booked date successfully",
		bookedDate,
	)
	return
}
