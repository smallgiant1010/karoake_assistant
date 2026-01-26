package handlers

import (
	"karaoke_assistant/backend/internal/data/sqlc"
	"karaoke_assitant/backend/internal/services"
	"karaoke_assistant/backend/internal/config"
	"karaoke_assistant/backend/internal/ai"
)

type Handler struct {
	queries *sqlc.Queries
	authService *services.AuthService
	songService *services.SongService
}

func NewHandler(queries_ *sqlc.Queries, cfg *config.Config, aiClient *ai.AIClient) *Handler {
	return &Handler{
		queries: queries_,
		authService: services.NewAuthService(queries_),
		songService: services.NewSongService(aiClient, cfg, queries_),
	}
}
