package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpError struct {
	Message string
	Status  int
}

func NewHttpError(message string, status int) *HttpError {
	return &HttpError{
		Message: message,
		Status:  status,
	}
}

func NewJsonError(c echo.Context, statusCode int, message string) error {
	if message == "" {
		message = "Internal Server Error"
	}

	return c.JSON(http.StatusInternalServerError, map[string]any{
		"message":    message,
		"error":      true,
		"statusCode": statusCode,
	})
}
