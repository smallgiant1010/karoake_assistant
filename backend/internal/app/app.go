package app

import (
	"context"
	"fmt"
	"karoake_assistant/backend/internal/ai"
	"karoake_assistant/backend/internal/auth"
	"karoake_assistant/backend/internal/http/handlers"
	"karoake_assistant/backend/internal/http/middleware"
	"karoake_assistant/backend/internal/platform/config"
	"karoake_assistant/backend/internal/platform/db"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type App struct {
	Mux        *http.ServeMux
	Database   *pgx.Conn
	JWTService *auth.JWTService
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
	// consider just passing the config straight to this client instead of drilling through handlers
	aiClient := ai.NewAIClient(
		client,
		false,
	)

	// JWT Service
	jwtService := auth.NewJWTService(cfg.JWT_SECRET, 24)

	// Handlers
	handler := handlers.NewHandler(queries, cfg, aiClient, jwtService)

	// Routing
	InitializeSongRoutes(mux, handler)
	InitializeAuthRoutes(mux, handler, jwtService)

	return &App{
		Mux:        mux,
		Database:   conn,
		JWTService: jwtService,
	}
}

func (a *App) Close() {
	a.Database.Close(context.Background())
}

func InitializeSongRoutes(m *http.ServeMux, songHandler *handlers.Handler) {
	m.HandleFunc("/songs/query", songHandler.Romanticize)
}

func InitializeAuthRoutes(m *http.ServeMux, authHandler *handlers.Handler, jwtService *auth.JWTService) {
	m.HandleFunc("/auth/add", authHandler.Signup)
	m.HandleFunc("/auth/login", authHandler.Login)
	protectedHandler := middleware.JWTContextMiddleware(jwtService, http.HandlerFunc(authHandler.Profile))
	m.Handle("/auth/profile", protectedHandler)
}
