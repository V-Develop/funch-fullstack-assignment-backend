package app

import (
	log "github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"gopkg.in/yaml.v2"
	"os"
)

type ErrorMessage struct {
	CommonError CommonError `yaml:"common_error"`
	AuthError   AuthError   `yaml:"auth_error"`
	UserError   UserError   `yaml:"user_error"`
}

type CommonError struct {
	BindJsonError       ErrorCode `yaml:"bind_json_error"`
	CannotRead          ErrorCode `yaml:"cannot_read"`
	NotFoundData        ErrorCode `yaml:"not_found_data"`
	InternalServerError ErrorCode `yaml:"internal_server_error"`
}

type AuthError struct {
	InvalidEmailEmpty               ErrorCode `yaml:"invalid_email_empty"`
	InvalidPasswordEmpty            ErrorCode `yaml:"invalid_password_empty"`
	InvalidEmailFormat              ErrorCode `yaml:"invalid_email_format"`
	InvalidPasswordLength           ErrorCode `yaml:"invalid_password_length"`
	InvalidPasswordUpperLowerCase   ErrorCode `yaml:"invalid_password_upper_lower_case"`
	InvalidPasswordSpecialCharacter ErrorCode `yaml:"invalid_password_special_character"`
	InvalidPasswordNotMatch         ErrorCode `yaml:"invalid_password_not_match"`
	InvalidDuplicateEmail           ErrorCode `yaml:"invalid_duplicate_email"`
	InvalidEmailOrPassword          ErrorCode `yaml:"invalid_email_or_password"`
	InvalidAccess                   ErrorCode `yaml:"invalid_access"`
}

type UserError struct {
	InvalidFirstnameOrLastname ErrorCode `yaml:"invalid_firstname_or_lastname"`
	InvalidPhoneNumber         ErrorCode `yaml:"invalid_phone_number"`
	InvalidBookRoomDate        ErrorCode `json:"invalid_book_room_date"`
	InvalidEmailEmpty          ErrorCode `json:"invalid_user_email_empty"`
	InvalidFirstnameEmpty      ErrorCode `json:"invalid_firstname_empty"`
	InvalidLastnameEmpty       ErrorCode `json:"invalid_lastname_empty"`
	InvalidPhoneNumberEmpty    ErrorCode `json:"invalid_phone_number_empty"`
	InvalidEmailFormat         ErrorCode `yaml:"invalid_user_email_format"`
}

type ErrorCode struct {
	Code    string `yaml:"code"`
	Message string `yaml:"message"`
}

func InitErrorMessage(loger log.Loger) *ErrorMessage {
	errorMessage = errorMessage.ReadErrorMessage(loger)
	return errorMessage
}

func (em *ErrorMessage) ReadErrorMessage(loger log.Loger) *ErrorMessage {
	errorYAML, errors := os.ReadFile("./config/error.yaml")
	if errors != nil {
		loger.LogError("SERVER: failed to read YAML file: "+errors.Error(), "0")
		panic(errors.Error())
	}
	errors = yaml.Unmarshal(errorYAML, &em)
	if errors != nil {
		loger.LogError("SERVER: Failed to parse YAML: "+errors.Error(), "0")
		panic(errors.Error())
	}
	return em
}
