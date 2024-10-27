package routes

import (
	"net/http"

	"github.com/smukk9/mkauth/server/internal/handlers"
)

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {

	router := &Router{
		ServeMux: http.NewServeMux(),
	}
	router.setupRoutes()
	return router
}

func (r *Router) setupRoutes() {

	//Auth Code Flow routes

	//https://pkg.go.dev/net/http#ServeMux
	r.Handle("GET /authorize", http.HandlerFunc(handlers.AuthorizeEndpoint()))

}
