package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func startsWithLetter(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z]`).MatchString(username)
}

func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}

func imgEncoding(fl validator.FieldLevel) bool {
	encoding := fl.Field().String()
	return regexp.MustCompile(`^image/(png|jpg|jpeg)$`).MatchString(encoding)
}

func ValidateRequest(req any) map[string]string {
	Validator.RegisterValidation("startsWithLetter", startsWithLetter)
	Validator.RegisterValidation("notBlank", notBlank)
	Validator.RegisterValidation("imgEncoding", imgEncoding)
	if err := Validator.Struct(req); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("failed on '%s' tag", err.Tag())
		}
		return errors
	}
	return nil
}

func CheckEncodingMatchesContent(enc, cont string) bool {
	// already validated above
	data, _ := base64.StdEncoding.DecodeString(cont)
	buf := bytes.NewBuffer(data)
	// garanteed to be one of these encodings
	switch enc {
	case "image/png":
		if _, err := png.Decode(buf); err != nil {
			return false
		}
	case "image/jpeg", "image/jpg":
		if _, err := jpeg.Decode(buf); err != nil {
			return false
		}
	}
	return true
}
