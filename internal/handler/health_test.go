package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	router := setupTestRouter()
	w := makeRequest(router, "GET", "/health")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	parseResponse(w, &response)

	assert.Equal(t, "ok", response["status"])
	assert.Equal(t, "mkauth", response["service"])
	assert.Equal(t, "1.0.0", response["version"])
}
