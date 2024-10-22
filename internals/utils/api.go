package utils

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ApiError struct {
	Code        int    `json:"-"`
	Message     string `json:"message"`
	Errors      any    `json:"errors,omitempty"`
	InternalErr string `json:"-"`
}

func (e ApiError) Error() string {
	if e.InternalErr == "" {
		return e.Message
	}
	return e.InternalErr
}

func InvalidJsonRequestError() ApiError {
	return ApiError{
		Code:    fiber.StatusBadRequest,
		Message: "failed to parse request body",
	}
}

func ValidationError(errs any) ApiError {
	return ApiError{
		Code:    fiber.StatusBadRequest,
		Message: "invalid request data",
		Errors:  errs,
	}
}

func ConflictError(msg string) ApiError {
	return ApiError{
		Code:    fiber.StatusConflict,
		Message: msg,
	}
}

func NotFoundError(msg string) ApiError {
	return ApiError{
		Code:    fiber.StatusNotFound,
		Message: msg,
	}
}

func InternalServerError(err error) ApiError {
	return ApiError{
		Code:        fiber.StatusInternalServerError,
		Message:     "internal server error",
		InternalErr: err.Error(),
	}
}

func UnauthorizedError() ApiError {
	return ApiError{
		Code:    fiber.StatusUnauthorized,
		Message: "unauthorized",
	}
}
