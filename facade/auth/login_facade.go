package auth

import (
	"crypto/rsa"
	"os"
	"time"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/entity"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	authRepository "github.com/V-Develop/funch-fullstack-assignment-backend/repository"
	util "github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type LoginFacade struct {
	config     *app.Config
	loger      util.Loger
	repository *authRepository.AuthRepository
	atclaims   *model.AccessTokenClaims
	rtclaims   *model.RefreshTokenClaims
}

func NewLoginFacade(config *app.Config, loger util.Loger) *LoginFacade {
	repository := authRepository.AuthRepository{}
	return &LoginFacade{
		config:     config,
		loger:      loger,
		repository: repository.NewAuthRepository(config),
	}
}

func (srv *LoginFacade) LoginLogic(request model.LoginRequest) (response model.ResponseWithPayload) {
	userCredential, errorResponse := srv.validateLoginRequest(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	transaction := srv.config.DataBase.Connection.Begin()

	accessTokenClaims, refreshTokenClaims, errorResponse := srv.updateLoginSession(transaction, userCredential, request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		transaction.Rollback()
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	accessToken, errorResponse := srv.signToken(accessTokenClaims, userCredential.Id)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	refreshToken, errorResponse := srv.signToken(refreshTokenClaims, userCredential.Id)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	errorResponse = srv.updateRefreshToken(transaction, userCredential.Email, refreshToken)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		transaction.Rollback()
		response = model.ResponseWithPayload{
			Status:  errorResponse.Status,
			Payload: nil,
		}
		return
	}

	transaction.Commit()

	response = util.CreateResponseWithPayload(
		200,
		"00000",
		"Logged in successfully",
		model.LoginResponse{
			AccessToken:        accessToken,
			RefreshToken:       refreshToken,
			AccessTokenExpire:  accessTokenClaims.StandardClaims.ExpiresAt,
			RefreshTokenExpire: refreshTokenClaims.StandardClaims.ExpiresAt,
		},
	)
	return
}

func (srv *LoginFacade) validateLoginRequest(request model.LoginRequest) (userCredential entity.UserCredential, errorResponse model.ResponseWithoutPayload) {
	userCredential, errorResponse = srv.validateEmail(request)
	if (errorResponse != model.ResponseWithoutPayload{}) {
		return
	}
	errorResponse = srv.validatePassword(userCredential, request)
	return
}

func (srv *LoginFacade) validateEmail(request model.LoginRequest) (userCredential entity.UserCredential, errorResponse model.ResponseWithoutPayload) {
	userCredential, err := srv.repository.CheckDuplicateEmail(srv.config.DataBase.Connection, request.Email)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	if (userCredential == entity.UserCredential{}) {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			401,
			srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword.Code,
			srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword.Message,
		)
		return
	}
	return
}

func (srv *LoginFacade) validatePassword(userCredential entity.UserCredential, request model.LoginRequest) (errorResponse model.ResponseWithoutPayload) {
	requestUserPassword := util.HashPassword(util.CombindPasswordAndSalt(request.Password, userCredential.Uuid))
	if requestUserPassword != userCredential.Password {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword), "0")
		errorResponse = util.CreateResponseWithoutPayload(
			401,
			srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword.Code,
			srv.config.ErrorMessage.AuthError.InvalidEmailOrPassword.Message,
		)
		return
	}
	return
}

func (srv *LoginFacade) updateLoginSession(transaction *gorm.DB, userCredential entity.UserCredential, request model.LoginRequest) (accessTokenClaims model.AccessTokenClaims, refreshTokenClaims model.RefreshTokenClaims, errorResponse model.ResponseWithoutPayload) {
	sessionId := uuid.New().String()
	lastLoginTime := time.Now()
	for {
		userCredential, err := srv.repository.CheckDuplicateSessionId(srv.config.DataBase.Connection, sessionId)
		if err != nil {
			srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
			errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
			return
		}
		if (userCredential == entity.UserCredential{}) {
			break
		}
		sessionId = uuid.New().String()
	}

	err := srv.repository.UpdateLoginSession(transaction, request.Email, lastLoginTime, sessionId)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	tokenExpireTime := lastLoginTime.Add(time.Hour * 3)
	accessTokenClaims = model.AccessTokenClaims{
		SessionId: sessionId,
		Id:        userCredential.Id,
		Role:      "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpireTime.Unix(),
		},
	}

	tokenExpireTime = lastLoginTime.Add(time.Hour * 3)
	refreshTokenClaims = model.RefreshTokenClaims{
		SessionId: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpireTime.Unix(),
		},
	}
	return
}

func (srv *LoginFacade) signToken(claims model.TokenClaims, userId int) (token string, errorResponse model.ResponseWithoutPayload) {
	var signKey *rsa.PrivateKey
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	signBytes, err := os.ReadFile(envs["PRIVATE_KEY_PATH"])
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err = newToken.SignedString(signKey)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}

	return
}

func (srv *LoginFacade) updateRefreshToken(transaction *gorm.DB, email string, refreshToken string) (errorResponse model.ResponseWithoutPayload) {
	err := srv.repository.UpdateRefreshToken(transaction, email, time.Now(), refreshToken)
	if err != nil {
		srv.loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse = util.CreateResponseWithoutPayload(500, "00000", srv.config.ErrorMessage.CommonError.InternalServerError.Message)
		return
	}
	return
}

func (srv *LoginFacade) GetClaimCurrent(c *gin.Context) {
	if claims, ok := c.Get("claim"); ok {
		srv.atclaims = claims.(*model.AccessTokenClaims)
	}
}

func (srv *LoginFacade) GetUserIdInClaim() int {
	return srv.atclaims.Id
}

func (srv *LoginFacade) GetRoleInClaim() string {
	return srv.atclaims.Role
}
