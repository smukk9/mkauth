package handlers

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"github.com/smukk9/mkauth/server/internal/db"
	"github.com/smukk9/mkauth/server/internal/errors"
	"github.com/smukk9/mkauth/server/internal/logger"
	"github.com/smukk9/mkauth/server/internal/validators"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request)

func AuthorizeEndpoint() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := validators.ValidateAuthorizeEndpoint(r); err != nil {
			logger.LogError("validation failed:", "error", err)
			errors.HandleError(w, err)
			return
		}

		// generate code and redirect
		code, err := generateAuthCode(16)

		if err != nil {
			errors.HandleError(w, err)
		}

		logger.LogInfo("Code Geenrated: ", code)
		redirect_uri := r.URL.Query().Get("redirect_uri")

		redirectUrl, parseErr := url.Parse(redirect_uri)
		if parseErr != nil {
			errors.HandleError(w, errors.NewAppError(http.StatusBadRequest, "Unable to parse the redirect_uri"))
			return
		}
		logger.LogInfo("Code Geenrated: ", code)

		query := redirectUrl.Query()
		query.Add("code", code)
		redirectUrl.RawQuery = query.Encode()

		dberr := db.Init()
		if dberr != nil {
			fmt.Println("Error initializing database:", dberr)
			return
		}

		dberr = db.InsertClient("Alice", "alice@example.com")
		if err != nil {
			fmt.Println("Error inserting client:", dberr)
			return
		}

		clients, dbqueryErr := db.QueryClients()
		if dbqueryErr != nil {
			fmt.Println("Error querying clients:", dbqueryErr)
			return
		}
		fmt.Println(clients)

		http.Redirect(w, r, redirectUrl.String(), http.StatusFound)

	}
}

// there is no restriction on this code in RFC however recomends short expiry and no collosion
func generateAuthCode(n int) (string, *errors.AppError) {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.NewAppError(http.StatusInternalServerError, "unable to generate random string")
	}

	// Filter out non-alphanumeric characters
	var result string
	for _, char := range b {
		index := int(char) % len(charset)
		result += string(charset[index])
	}

	return result[:n], nil
}
