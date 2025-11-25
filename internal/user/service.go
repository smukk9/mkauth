package user

import (
	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
)

type Service struct {
	db *db.Database
}

func NewService(db *db.Database) *Service {
	return &Service{
		db: db,
	}
}

type Handler struct {
	service *Service
}

func NewHandler(db *db.Database, cfg *config.Config) *Handler {
	return &Handler{
		service: NewService(db),
	}
}
