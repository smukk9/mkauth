package main

import (
	"log"

	"github.com/smukk9/mkauth/internal/config"
	"github.com/smukk9/mkauth/internal/handler"
)

func main() {

	cfg, err := config.Load()

	//create new router with all depedencey
	r := handler.NewRouter(cfg)
	engine := r.Setup()

	if err != nil {

		log.Fatalf("failed to load config: %v", err)
	}
	log.Println("Server Starting...")
	if err := engine.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
