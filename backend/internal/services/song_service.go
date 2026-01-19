package services	

import (
	"context"
	"karaoke_assistant/backend/internal/domains"
	"karaoke_assistant/backend/internal/repository"
)

type SongService struct {
	repo repository.AISongRepository
}

func NewSongService(repo_ repository.AISongRepository) *SongService {
	return &SongService{
		repo: repo_,
	}
}

func (s *SongService) RomanticizeSong(ctx context.Context, song *domains.Song) (string, error) {
	return s.repo.CallToAI(ctx, song)
}



