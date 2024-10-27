package errors

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/smukk9/mkauth/server/internal/logger"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// this make AppError type, walks like duck , quacks liek douck, it must be duck
func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func HandleError(w http.ResponseWriter, err error) {

	var appErr *AppError
	if errors.As(err, &appErr) {

		w.Header().Set("content-type", "application/json")
		logger.LogInfo("App Erro matched")
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
	} else {
		//generic error
		w.Header().Set("Content-Type", "application/json")
		logger.LogInfo("Generic Error matched")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewAppError(http.StatusInternalServerError, "Internal Server Error"))

	}
}
