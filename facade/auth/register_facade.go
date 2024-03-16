package auth

import (
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	authRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/google/uuid"
)

type RegisterFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *authRepository.AuthRepository
}

func NewRegisterFacade(config *app.Config, loger util.Loger) *RegisterFacade {
	repository := authRepository.AuthRepository{}
	return &RegisterFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewAuthRepository(config),
	}
}

func (srv *RegisterFacade) RegisterLogic(request model.RegisterRequest) (response model.ResponseWithoutPayload) {
	errorResponse := srv.validateRegisterRequest(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}

	errorResponse = srv.createUser(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = errorResponse
		return
	}

	response = util.CreateResponseWithoutPayload(
		201,
		"00000",
		"Registered successfully",
	)
	return
}

func (srv *RegisterFacade) validateRegisterRequest(request model.RegisterRequest) (errorResponse model.ResponseWithoutPayload) {
	errorResponse = srv.validateEmail(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}
	errorResponse = srv.validatePassword(request)
	return
}

func (srv *RegisterFacade) validateEmail(request model.RegisterRequest) (errorResponse model.ResponseWithoutPayload) {
	if util.IsEmptyOrBlank(request.Email) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidEmailEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidEmailEmpty.Code,
			srv.config.ErrorMessage.AuthError.InvalidEmailEmpty.Message,
		)
		return
	}
	if util.IsInvalidEmail(request.Email) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidEmailFormat), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidEmailFormat.Code,
			srv.config.ErrorMessage.AuthError.InvalidEmailFormat.Message,
		)
		return
	}
	userCredential, err := srv.repository.CheckDuplicateEmail(srv.config.DataBase.Connection, request.Email)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	if (userCredential != entity.UserCredential{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidDuplicateEmail), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidDuplicateEmail.Code,
			srv.config.ErrorMessage.AuthError.InvalidDuplicateEmail.Message,
		)
		return
	}
	return
}

func (srv *RegisterFacade) validatePassword(request model.RegisterRequest) (errorResponse model.ResponseWithoutPayload) {
	if util.IsEmptyOrBlank(request.Password) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidPasswordEmpty), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidPasswordEmpty.Code,
			srv.config.ErrorMessage.AuthError.InvalidPasswordEmpty.Message,
		)
		return
	}
	if len(request.Password) < 8 {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidPasswordLength), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidPasswordLength.Code,
			srv.config.ErrorMessage.AuthError.InvalidPasswordLength.Message,
		)
		return
	}
	if util.IsPasswordNotContainSpecialCharacter(request.Password) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidPasswordSpecialCharacter), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidPasswordSpecialCharacter.Code,
			srv.config.ErrorMessage.AuthError.InvalidPasswordSpecialCharacter.Message,
		)
		return
	}
	if util.IsPasswordNotContainUpperAndLowerCase(request.Password) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidPasswordUpperLowerCase), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidPasswordUpperLowerCase.Code,
			srv.config.ErrorMessage.AuthError.InvalidPasswordUpperLowerCase.Message,
		)
		return
	}
	if request.Password != request.RepeatPassword {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidPasswordNotMatch), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			400,
			srv.config.ErrorMessage.AuthError.InvalidPasswordNotMatch.Code,
			srv.config.ErrorMessage.AuthError.InvalidPasswordNotMatch.Message,
		)
		return
	}
	return
}

func (srv *RegisterFacade) createUser(request model.RegisterRequest) (errorResponse model.ResponseWithoutPayload) {
	userUuid := uuid.New().String()
	for {
		userCredential, err := srv.repository.CheckDuplicateUUID(srv.config.DataBase.Connection, userUuid)
		if err != nil {
			srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
			errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
			return
		}
		if (userCredential == entity.UserCredential{}) {
			break
		}
		userUuid = uuid.New().String()
	}

	hashedPassword := util.HashPassword(util.CombindPasswordAndSalt(request.Password, userUuid))
	time := time.Now()

	newUserCredential := entity.UserCredential{
		Uuid:       userUuid,
		Email:      request.Email,
		Password:   hashedPassword,
		IsBacklist: false,
		IsVerify:   false,
		CreatedAt:  time,
		CreatedBy:  "server",
		UpdatedAt:  time,
		UpdatedBy:  "server",
	}

	transaction := srv.config.DataBase.Connection.Begin()
	err := srv.repository.CreateUserCredential(transaction, newUserCredential)
	if err != nil {
		transaction.Rollback()
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	userCredential, err := srv.repository.GetUserCredentialByEmail(transaction, newUserCredential.Email)
	if err != nil {
		transaction.Rollback()
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	newUserProfile := entity.UserProfile{
		UserId:    userCredential.Id,
		CreatedAt: time,
		CreatedBy: "server",
		UpdatedAt: time,
		UpdatedBy: "server",
	}

	err = srv.repository.CreateUserProfile(transaction, newUserProfile)
	if err != nil {
		transaction.Rollback()
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	transaction.Commit()
	return
}
