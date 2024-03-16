package user

import (
	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	userRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
)

type GetUserProfileFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *userRepository.UserRepository
}

func NewGetUserProfileFacade(config *app.Config, loger util.Loger) *GetUserProfileFacade {
	repository := userRepository.UserRepository{}
	return &GetUserProfileFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewUserRepository(config),
	}
}

func (srv *GetUserProfileFacade) GetUserProfileLogic(userId int) (response model.ResponseWithPayload) {
	userCredential, errorResponse := srv.getUserEmail(userId)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}
	userProfile, errorResponse := srv.getUserProfile(userId)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	response = util.CreateResponseWithPayload(
		200,
		"00000",
		"Get user profile successfully",
		model.GetUserProfileResponse{
			Email:       userCredential.Email,
			Firstname:   userProfile.Firstname,
			Lastname:    userProfile.Lastname,
			PhoneNumber: userProfile.PhoneNumber,
		},
	)
	return
}

func (srv *GetUserProfileFacade) getUserEmail(userId int) (userCredential entity.UserCredential, errorResponse model.ResponseWithoutPayload) {
	userCredential, err := srv.repository.GetUserCredentialById(srv.config.DataBase.Connection, userId)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	if (userCredential == entity.UserCredential{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.CommonError.NotFoundData), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			404,
			srv.config.ErrorMessage.CommonError.NotFoundData.Code,
			srv.config.ErrorMessage.CommonError.NotFoundData.Message,
		)
		return
	}
	return
}

func (srv *GetUserProfileFacade) getUserProfile(userId int) (userProfile entity.UserProfile, errorResponse model.ResponseWithoutPayload) {
	userProfile, err := srv.repository.GetUserProfileById(srv.config.DataBase.Connection, userId)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	if (userProfile == entity.UserProfile{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.CommonError.NotFoundData), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			404,
			srv.config.ErrorMessage.CommonError.NotFoundData.Code,
			srv.config.ErrorMessage.CommonError.NotFoundData.Message,
		)
		return
	}
	return
}
