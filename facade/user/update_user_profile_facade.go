package user

import (
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	userRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
)

type UpdateUserProfileFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *userRepository.UserRepository
}

func NewUpdateUserProfileFacade(config *app.Config, loger util.Loger) *UpdateUserProfileFacade {
	repository := userRepository.UserRepository{}
	return &UpdateUserProfileFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewUserRepository(config),
	}
}

func (srv *UpdateUserProfileFacade) UpdateUserProfileLogic(request model.UpdateUserProfileRequest, userId int) (response model.ResponseWithoutPayload) {
	errorResponse := srv.validateUpdateUserProfileRequest(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}
	errorResponse = srv.updateUserProfile(request, userId)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}
	response = util.CreateResponseWithoutPayload(
		200,
		"00000",
		"Updated user profile successfully",
	)
	return
}

func (srv *UpdateUserProfileFacade) validateUpdateUserProfileRequest(request model.UpdateUserProfileRequest) (errorResponse model.ResponseWithoutPayload) {
	errorResponse = srv.validateFirstnameAndLastname(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}
	errorResponse = srv.validatePhoneNumber(request)
	return
}

func (srv *UpdateUserProfileFacade) validateFirstnameAndLastname(request model.UpdateUserProfileRequest) (errorResponse model.ResponseWithoutPayload) {
	if(util.IsEmptyOrBlank(request.Firstname)) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty.Code,
			srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty.Message,
		)
		return
	}
	if(util.IsEmptyOrBlank(request.Lastname)) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidLastnameEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidLastnameEmpty.Code,
			srv.config.ErrorMessage.UserError.InvalidLastnameEmpty.Message,
		)
		return
	}
	if util.IsInvalidEnglishString(request.Firstname) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname.Code,
			srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname.Message,
		)
		return
	}
	if util.IsInvalidEnglishString(request.Lastname) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname.Code,
			srv.config.ErrorMessage.UserError.InvalidFirstnameOrLastname.Message,
		)
		return
	}
	return
}

func (srv *UpdateUserProfileFacade) validatePhoneNumber(request model.UpdateUserProfileRequest) (errorResponse model.ResponseWithoutPayload) {
	if(util.IsEmptyOrBlank(request.PhoneNumber)) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidPhoneNumberEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidPhoneNumberEmpty.Code,
			srv.config.ErrorMessage.UserError.InvalidPhoneNumberEmpty.Message,
		)
		return
	}
	if util.IsInvalidPhoneNumber(request.PhoneNumber) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidPhoneNumber), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidPhoneNumber.Code,
			srv.config.ErrorMessage.UserError.InvalidPhoneNumber.Message,
		)
		return
	}
	return
}

func (srv *UpdateUserProfileFacade) updateUserProfile(request model.UpdateUserProfileRequest, userId int) (errorResponse model.ResponseWithoutPayload) {
	err := srv.repository.UpdateUserProfile(srv.config.DataBase.Connection, userId, request.Firstname, request.Lastname, request.PhoneNumber, time.Now())
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	return
}
