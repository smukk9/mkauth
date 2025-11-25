package handlers

import (
	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
)

// Handler holds common dependencies for web UI handlers
type Handler struct {
	cfg *config.Config
	db  *db.Database
}

// New creates a base handler with dependencies
func New(cfg *config.Config, db *db.Database) *Handler {
	return &Handler{
		cfg: cfg,
		db:  db,
	}
}
