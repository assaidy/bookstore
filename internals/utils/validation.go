package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var Validator = validator.New()

func startsWithLetter(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z]`).MatchString(username)
}

func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func ValidateRequest(req any) map[string]string {
	Validator.RegisterValidation("startsWithLetter", startsWithLetter)
	Validator.RegisterValidation("notBlank", notBlank)
	if err := Validator.Struct(req); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("failed on '%s' tag", err.Tag())
		}
		return errors
	}
	return nil
}
