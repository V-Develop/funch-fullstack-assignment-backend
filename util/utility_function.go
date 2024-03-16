package util

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
)

func ConvertObjectToString(object interface{}) string {
	return fmt.Sprintf("%v", object)
}

func CreateResponseWithoutPayload(httpStatus int, code string, message string) model.ResponseWithoutPayload {
	return model.ResponseWithoutPayload{
		Status: model.ResponseHeader{
			HttpStatus: httpStatus,
			Code:      code,
			Message: message,
		},
	}
}

func CreateResponseWithPayload(httpStatus int, code string, message string, payload interface{}) model.ResponseWithPayload {
	return model.ResponseWithPayload{
		Status: model.ResponseHeader{
			HttpStatus: httpStatus,
			Code:      code,
			Message: message,
		},
		Payload: payload,
	}
}

func CombindPasswordAndSalt(password string, salt string) string {
	passwordWithSalt := password + salt[12:24]
	return passwordWithSalt
}

func HashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}