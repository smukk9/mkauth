package handlers

import (
	"net/http"

	"github.com/smukk9/mkauth/internal/web/templates"
)

type HomeHandler struct {
	*Handler // Embeds base handler (gets cfg, db)
}

func NewHomeHandler(base *Handler) *HomeHandler {
	return &HomeHandler{Handler: base}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Render home page template
	templates.Home(
		h.cfg.Server.Service,
		h.cfg.Server.Version,
	).Render(r.Context(), w)
}
