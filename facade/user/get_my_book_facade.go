package user

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	userRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
)

type GetMyBookFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *userRepository.UserRepository
}

func NewGetMyBookFacade(config *app.Config, loger util.Loger) *GetMyBookFacade {
	repository := userRepository.UserRepository{}
	return &GetMyBookFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewUserRepository(config),
	}
}

func (srv *GetMyBookFacade) GetMyBookLogic(userId int) (response model.ResponseWithPayload) {
	userBookdate, err := srv.repository.GetBookedDateById(srv.config.DataBase.Connection, userId)
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
		"Get user booked date successfully",
		userBookdate,
	)
	return
}
