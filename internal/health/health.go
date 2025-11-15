package health

import (
	"encoding/json"
	"net/http"

	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
)

type Handler struct {
	service *Service
	cfg     *config.Config
}

func NewHandler(db *db.Database, cfg *config.Config) *Handler {
	return &Handler{
		service: NewService(db),
		cfg:     cfg,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get database health status
	dbStatus, err := h.service.CheckDatabase()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"service": h.cfg.Server.Service,
			"version": h.cfg.Server.Version,
			"database": map[string]interface{}{
				"connected": false,
				"error":     err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"service": h.cfg.Server.Service,
		"version": h.cfg.Server.Version,
		"database": map[string]interface{}{
			"connected":  true,
			"last_check": dbStatus.LastCheck,
		},
	})
}
