package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/db"
	"github.com/smukk9/mkauth/internal/health"
)

type Server struct {
	cfg    *config.Config
	db     *db.Database
	mux    *http.ServeMux
	server *http.Server
}

func New(cfg *config.Config, db *db.Database) (*Server, error) {
	mux := http.NewServeMux()

	s := &Server{
		cfg: cfg,
		db:  db,
		mux: mux,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
			Handler: mux,
		},
	}

	// Register all routes
	s.registerRoutes()

	return s, nil
}

func (s *Server) registerRoutes() {
	// Register health routes
	health.RegisterRoutes(s.mux, s.db, s.cfg)

	// Future: Register user routes
	// user.RegisterRoutes(s.mux, s.db, s.cfg)

	// Future: Register oauth routes
	// oauth.RegisterRoutes(s.mux, s.db, s.cfg)
}

func (s *Server) Start() error {
	log.Printf("Starting %s %s on %s", s.cfg.Server.Service, s.cfg.Server.Version, s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	log.Println("Shutting down server...")
	return s.db.Close()
}
