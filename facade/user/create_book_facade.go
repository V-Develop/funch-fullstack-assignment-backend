package user

import (
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	userRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/google/uuid"
)

type CreateBookFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *userRepository.UserRepository
}

func NewCreateBookFacade(config *app.Config, loger util.Loger) *CreateBookFacade {
	repository := userRepository.UserRepository{}
	return &CreateBookFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewUserRepository(config),
	}
}

func (srv *CreateBookFacade) CreateBookLogic(request model.CreateBookRequest, userId int) (response model.ResponseWithoutPayload) {
	errorResponse := srv.validateEmail(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}
	errorResponse = srv.validateFirstnameAndLastname(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}
	errorResponse = srv.validatePhoneNumber(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}

	errorResponse = srv.validateBookRoomDate(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}
	errorResponse = srv.createBookRoomDate(request, userId)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}

	response = util.CreateResponseWithoutPayload(
		200,
		"00000",
		"Created book room successfully",
	)
	return
}

func (srv *CreateBookFacade) validateBookRoomDate(request model.CreateBookRequest) (errorResponse model.ResponseWithoutPayload) {
	if (request.CheckinAt == time.Time{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidBookRoomDate), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidBookRoomDate.Code,
			srv.config.ErrorMessage.UserError.InvalidBookRoomDate.Message,
		)
		return
	}
	if (request.CheckoutAt == time.Time{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidBookRoomDate), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidBookRoomDate.Code,
			srv.config.ErrorMessage.UserError.InvalidBookRoomDate.Message,
		)
		return
	}
	return
}

func (srv *CreateBookFacade) validateEmail(request model.CreateBookRequest) (errorResponse model.ResponseWithoutPayload) {
	if util.IsEmptyOrBlank(request.Email) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidEmailEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidEmailEmpty.Code,
			srv.config.ErrorMessage.UserError.InvalidEmailEmpty.Message,
		)
		return
	}
	if util.IsInvalidEmail(request.Email) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidEmailFormat), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidEmailFormat.Code,
			srv.config.ErrorMessage.UserError.InvalidEmailFormat.Message,
		)
		return
	}
	return
}

func (srv *CreateBookFacade) validateFirstnameAndLastname(request model.CreateBookRequest) (errorResponse model.ResponseWithoutPayload) {
	if util.IsEmptyOrBlank(request.Firstname) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty.Code,
			srv.config.ErrorMessage.UserError.InvalidFirstnameEmpty.Message,
		)
		return
	}
	if util.IsEmptyOrBlank(request.Lastname) {
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

func (srv *CreateBookFacade) validatePhoneNumber(request model.CreateBookRequest) (errorResponse model.ResponseWithoutPayload) {
	if util.IsEmptyOrBlank(request.PhoneNumber) {
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

func (srv *CreateBookFacade) createBookRoomDate(request model.CreateBookRequest, userId int) (errorResponse model.ResponseWithoutPayload) {
	bookRoomUuid := uuid.New().String()
	for {
		bookHistory, err := srv.repository.CheckDuplicateBookRoomUUID(srv.config.DataBase.Connection, bookRoomUuid)
		if err != nil {
			srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
			errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
			return
		}
		if (bookHistory == entity.BookHistory{}) {
			break
		}
		bookRoomUuid = uuid.New().String()
	}

	time := time.Now()
	newBookRoomRecord := entity.BookHistory{
		Uuid:        bookRoomUuid,
		BookerId:    userId,
		CheckinAt:   request.CheckinAt,
		CheckoutAt:  request.CheckoutAt,
		Email:       request.Email,
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   time,
		CreatedBy:   "server",
		UpdatedAt:   time,
		UpdatedBy:   "server",
	}
	err := srv.repository.CreateBookRoom(srv.config.DataBase.Connection, newBookRoomRecord)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	return
}
