package main

import (
	"fmt"
	"net/http"	
	"os"
	"karoake_assistant/backend/internal/app"
	"karoake_assistant/backend/internal/http/middleware"
	"karoake_assistant/backend/internal/platform/config"
)

func main() {
	cfg := config.NewConfig()

	application := app.NewApp(cfg)
	defer application.Close()

	handler := middleware.CORSMiddleware(application.Mux)

	server := &http.Server {
		Addr: ":" + cfg.ServerPort,
		Handler: handler,
	}

	fmt.Printf("starting server on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server %v\n", err)
		os.Exit(1)
	}
}	
