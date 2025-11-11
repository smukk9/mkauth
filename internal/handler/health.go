package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthhandler struct {
	version string
	service string
}

func NewHealthHandler(version, service string) *Healthhandler {

	return &Healthhandler{
		version: version,
		service: service,
	}
}

func (h *Healthhandler) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": h.version,
		"service": h.service,
	})
}
