// handler/test_helpers.go
package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/smukk9/mkauth/internal/config"
)

// setupTestRouter creates a router for testing
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		Server: config.Server{
			Port:    ":8080",
			Mode:    "test",
			Version: "1.0.0",
			Service: "mkauth",
		},
	}
	router := NewRouter(cfg)
	return router.Setup()
}

// makeRequest is a helper to make test HTTP requests
func makeRequest(router *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// parseResponse unmarshals JSON response
func parseResponse(w *httptest.ResponseRecorder, v interface{}) error {
	return json.Unmarshal(w.Body.Bytes(), v)
}
