package main

import (
	"log"
	"net/http"

	"github.com/smukk9/mkauth/server/internal/logger"
	"github.com/smukk9/mkauth/server/internal/routes"
)

func main() {

	logger.LogInfo("Hello")

	router := routes.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router, // Choose your preferred router
	}

	log.Fatal(server.ListenAndServe())

}
