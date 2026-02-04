package handlers

import (
	"karoake_assistant/backend/internal/ai"
	"karoake_assistant/backend/internal/platform/config"
	"karoake_assistant/backend/internal/data/sqlc"
	"karoake_assistant/backend/internal/services"
)

type Handler struct {
	queries     *sqlc.Queries
	authService *services.AuthService
	songService *services.SongService
}

func NewHandler(queries_ *sqlc.Queries, cfg *config.Config, aiClient *ai.AIClient) *Handler {
	return &Handler{
		queries:     queries_,
		authService: services.NewAuthService(queries_),
		songService: services.NewSongService(queries_, cfg, aiClient),
	}
}
