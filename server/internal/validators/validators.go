package validators

import (
	"net/http"

	"github.com/smukk9/mkauth/server/internal/errors"
)

func ValidateAuthorizeEndpoint(r *http.Request) *errors.AppError {

	query := r.URL.Query()
	clientId := query.Get("client_id")
	responseType := query.Get("response_type")
	// state := query.Get("state")
	// redirectURI := query.Get("redirect_uri")

	if len(query) < 4 {
		return errors.NewAppError(http.StatusBadRequest, "missing required query parameters")
	}
	if responseType != "code" {
		return errors.NewAppError(http.StatusBadRequest, "invalid repsone_type")
	}
	if clientId == "" {
		return errors.NewAppError(http.StatusBadRequest, "client_id can not be empty")
	}

	if clientId != "client1" {
		return errors.NewAppError(http.StatusBadRequest, "invalid client")
	}
	// if responseType != "code" {
	// 	return errors.NewAppError(http.StatusBadRequest, "invalid grant_type")
	// }

	return nil
}
