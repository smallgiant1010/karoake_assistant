package app

import (
	"github.com/jackc/pgx/v5"
	"karoake_assistant/backend/internal/ai"
	"karoake_assistant/backend/internal/http/handlers"
	"karoake_assistant/backend/internal/platform/config"
	"karoake_assistant/backend/internal/platform/db"
	"context"
	"net/http"
	"fmt"
)

type App struct {
	Mux      *http.ServeMux
	Database *pgx.Conn
}

func NewApp(cfg *config.Config) *App {
	mux := http.NewServeMux()
	client := &http.Client{}
	conn, queries, err := db.NewDatabaseConnection(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("error occured with database: %v\n", err)
		return nil
	}

	// AI Client
	aiClient := ai.NewAIClient(
		client,
		false,
	)

	// Handlers
	handler := handlers.NewHandler(queries, cfg, aiClient)

	// Routing
	InitializeSongRoutes(mux, handler)
	InitializeAuthRoutes(mux, handler)

	return &App{
		Mux:      mux,
		Database: conn,
	}
}

func (a *App) Close() {
	a.Database.Close(context.Background())
}

func InitializeSongRoutes(m *http.ServeMux, songHandler *handlers.Handler) {
	m.HandleFunc("/songs/query", songHandler.Romanticize)
}

func InitializeAuthRoutes(m *http.ServeMux, authHandler *handlers.Handler) {
	m.HandleFunc("/auth/add", authHandler.Signup)
	m.HandleFunc("/auth/login", authHandler.GetUser)
}
