package app

import (
	"net/http"	
	"context"
	"github.com/jackc/pgx/v5"
	"karaoke_assistant/backend/internal/http/handlers"
	"karaoke_assistant/backend/internal/services"
	"karaoke_assistant/backend/internal/repository"
	"karaoke_assistant/backend/internal/config"
)

type App struct {
	Mux *http.ServeMux
	Database *pgx.Conn
}

func NewApp(cfg *config.Config) *App {
	mux := http.NewServeMux()
	client := &http.Client{}
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)

	// Repository
	aiRepo := repository.NewAIAPIRepository(
			client,
			cfg.AIAPIURL,
			cfg.Model
			false,
			cfg.SystemPrompt,
			conn,
	)
	authRepo := repository.NewAuthRepository(
		conn,
	)

	// Services
	songService := services.NewSongService(aiRepo)
	authService := services.NewAuthService(authRepo)

	// Handlers
	songHandler := handlers.NewSongHandler(songService)
	authHandler := handlers.NewAuthHandler(authService)

	// Routing
	InitalizeSongRoutes(mux, songHandler)
	InitializeAuthRoutes(mux, songHandler)

	return &App{
		Mux: mux,
		Database: conn,
	}
}

func (a *App) Close() {
	a.Database.Close()
}

func InitializeSongRoutes(m *http.ServeMux, songHandler *SongHandler) {
	m.HandleFunc("/songs/query", songHandler.PostSong)
}

func InitializeAuthRoutes(m *http.ServeMux, authHandler *AuthHandler) {
	m.HandleFunc("/auth/add", authHandler.PostUser)
	m.HandleFunc("/auth/login", authHandler.GetUser)
}
