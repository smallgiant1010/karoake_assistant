package main

import (
	"fmt"
	"net/http"	
	"os"
	"karaoke_assistant/backend/internal/http/handlers"
	"karaoke_assistant/backend/internal/services"
	"karaoke_assistant/backend/internal/repository"
	"karaoke_assistant/backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found, relying on environment variables\n")
	}

	client := &http.Client{}

	aiRepo := repository.NewAIAPIRepository(
		client,
		config.AI_API,
		config.MODEL,
		false,
		config.SYSTEM_PROMPT,
	)

	songService := services.NewSongService(
		aiRepo,
	)

	songHandler := handlers.NewSongHandler( 
		songService,
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/songs", songHandler.PostSong)

	server := &http.Server {
		Addr: config.PORT,
		Handler: mux,
	}

	fmt.Printf("starting server on %v\n", config.PORT)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server %v\n", err)
		os.Exit(1)
	}
}	
