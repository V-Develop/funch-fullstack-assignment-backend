package util

import (
	"net/mail"
	"regexp"
	"strings"
	"unicode"

	"github.com/V-Develop/funch-fullstack-assignment-backend/model"
	"github.com/gin-gonic/gin"
)

func IsEmptyOrBlank(input string) bool {
	trimmedInput := strings.TrimSpace(input)
	return trimmedInput == ""
}

func IsInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func IsPasswordNotContainSpecialCharacter(input string) bool {
	pattern := "^[a-zA-Z0-9]*$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(input)
}

func IsPasswordNotContainUpperAndLowerCase(password string) bool {
	hasUpper := false
	hasLower := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
	}

	return !(hasUpper && hasLower)
}

func IsPasswordNotContainUpperOrLowerCase(password string) bool {
	specialCharRegex := regexp.MustCompile(`[~!@#$%^&*()-_+={}[\]|;:',.<>?/\\]`)
	return !specialCharRegex.MatchString(password)
}

func IsInvalidFormat(c *gin.Context, request interface{}, code string, message string) (response model.ResponseWithoutPayload) {
	if err := c.BindJSON(&request); err != nil {
		response = CreateResponseWithoutPayload(
			400,
			code,
			message,
		)
		return
	}
	return
}

func IsInvalidEnglishString(input string) bool {
	pattern := "^[a-zA-Z]+$"
	reg := regexp.MustCompile(pattern)
	return !reg.MatchString(input)
}

func IsInvalidPhoneNumber(input string) bool {
	pattern := `^\d{3}-\d{3}-\d{4}$`
	reg := regexp.MustCompile(pattern)
	return !reg.MatchString(input)
}
