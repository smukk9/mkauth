package client

import (
	"net/http"

	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
)

func RegisterRoutes(mux *http.ServeMux, db *db.Database, cfg *config.Config) {
	handler := NewHandler(db, cfg)

	mux.HandleFunc("GET /client", handler.GetClients)
	mux.HandleFunc("POST /client", handler.StoreClient)
}
