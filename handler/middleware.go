package handler

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Middleware struct {
	config *app.Config
}

func NewMiddleware(config *app.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}

func (m *Middleware) ValidateRequestHeader(c *gin.Context) (err error) {
	loger := util.NewLog("0", "LOG")
	apiKey := c.Request.Header.Get("Api-Key")
	if util.IsEmptyOrBlank(apiKey) {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(m.config.ErrorMessage.AuthError.InvalidAccess), "0")
		errorResponse := util.CreateResponseWithoutPayload(
			401,
			m.config.ErrorMessage.AuthError.InvalidAccess.Code,
			m.config.ErrorMessage.AuthError.InvalidAccess.Message,
		)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	}
	m.validateApiKey(apiKey, c, loger)
	err = m.authWithClaims(c, loger)
	return
}

func (m *Middleware) validateApiKey(apiKey string, c *gin.Context, loger util.Loger) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse := util.CreateResponseWithoutPayload(500, "00000", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	if apiKey != envs["API_KEY"] {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(m.config.ErrorMessage.AuthError.InvalidAccess), "0")
		errorResponse := util.CreateResponseWithoutPayload(
			401,
			m.config.ErrorMessage.AuthError.InvalidAccess.Code,
			m.config.ErrorMessage.AuthError.InvalidAccess.Message,
		)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	}
}

func (m *Middleware) authWithClaims(c *gin.Context, loger util.Loger) (err error) {
	var envs map[string]string
	envs, err = godotenv.Read(".env")
	if err != nil {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse := util.CreateResponseWithoutPayload(500, "00000", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	token, ok := c.Request.Header["Authorization"]
	if len(token) == 0 || !ok {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(m.config.ErrorMessage.AuthError.InvalidAccess), "0")
		errorResponse := util.CreateResponseWithoutPayload(
			401,
			m.config.ErrorMessage.AuthError.InvalidAccess.Code,
			m.config.ErrorMessage.AuthError.InvalidAccess.Message,
		)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	accessTokenClaims := &model.AccessTokenClaims{}
	var verifyKey *rsa.PublicKey

	verifyBytes, err := os.ReadFile(envs["PUBLIC_KEY_PATH"])
	if err != nil {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse := util.CreateResponseWithoutPayload(500, "00000", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		loger.LogErrorf("LOGIC: %+v", util.ConvertObjectToString(err.Error()), "0")
		errorResponse := util.CreateResponseWithoutPayload(500, "00000", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
		return
	}
	tkn, err := jwt.ParseWithClaims(token[0], accessTokenClaims, func(token *jwt.Token) (interface{}, error) {
		err = fmt.Errorf("invalid token")
		return verifyKey, nil
	})

	if tkn == nil || !tkn.Valid {
		return
	}

	c.Set("claim", accessTokenClaims)
	return

}
