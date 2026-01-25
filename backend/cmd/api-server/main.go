package main

import (
	"fmt"
	"net/http"	
	"os"
	"karaoke_assistant/backend/internal/app"
	"karaoke_assistant/backend/internal/config"
)

func main() {
	cfg := config.NewConfig()

	application := app.NewApp(cfg)
	defer application.Close()

	server := &http.Server {
		Addr: ":" + cfg.ServerPort,
		Handler: application.Mux,
	}

	fmt.Printf("starting server on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server %v\n", err)
		os.Exit(1)
	}
}	
