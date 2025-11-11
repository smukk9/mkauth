package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smukk9/mkauth/internal/config"
)

type Router struct {
	config        *config.Config
	healthHandler *Healthhandler
}

func NewRouter(cfg *config.Config) *Router {

	return &Router{
		config:        cfg,
		healthHandler: NewHealthHandler(cfg.Server.Version, cfg.Server.Service),
	}
}

func (r *Router) Setup() *gin.Engine {

	router := gin.Default()

	r.setupHealthRoutes(router)

	return router

}

func (r *Router) setupHealthRoutes(router *gin.Engine) {
	router.GET("/health", r.healthHandler.Handle)
}
