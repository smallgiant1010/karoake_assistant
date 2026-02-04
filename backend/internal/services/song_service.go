package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"karoake_assistant/backend/internal/ai"
	"karoake_assistant/backend/internal/data/mapper"
	"karoake_assistant/backend/internal/data/sqlc"
	"karoake_assistant/backend/internal/domains"
	"karoake_assistant/backend/internal/http/transport"
	"karoake_assistant/backend/internal/platform/config"

	"github.com/jackc/pgx/v5/pgtype"
)

type SongService struct {
	queries *sqlc.Queries
	api     ai.AIAPI
	cfg     *config.Config
}

func NewSongService(queries_ *sqlc.Queries, cfg_ *config.Config, api_ ai.AIAPI) *SongService {
	return &SongService{
		queries: queries_,
		cfg:     cfg_,
		api:     api_,
	}
}

func (s *SongService) RomanticizeSong(ctx context.Context, req *transport.CreateSongRequest) (*domains.Song, error) {
	if req.Title == "" {
		return nil, fmt.Errorf("title must not be empty")
	}

	if req.Language == "" {
		return nil, fmt.Errorf("language must not be empty")
	}

	if req.Lyrics == "" {
		return nil, fmt.Errorf("lyrics must not be empty")
	}

	romanticized, err := s.api.CallToAI(ctx, s.cfg, req.Lyrics)
	if err != nil {
		return nil, fmt.Errorf("error occured with ai model: %v", err)
	}

	var titleFormatted pgtype.Text
	if err := titleFormatted.Scan(req.Title); err != nil {
		return nil, fmt.Errorf("error occured converting title: %v", err)
	}

	var isGeneratedFormatted pgtype.Bool
	if err := isGeneratedFormatted.Scan(req.IsGenerated); err != nil {
		return nil, fmt.Errorf("error occured converting isGenerated: %v", err)
	}

	var lyricsFormatted pgtype.Text
	if err := lyricsFormatted.Scan(romanticized); err != nil {
		return nil, fmt.Errorf("error occured converting lyrics: %v", err)
	}

	songs, err := s.queries.GetSongsByTitle(ctx, titleFormatted)
	if errors.Is(err, sql.ErrNoRows) {
		newSong, err := s.queries.CreateSong(ctx, sqlc.CreateSongParams{
			Language:    req.Language,
			Title:       req.Title,
			Isgenerated: isGeneratedFormatted,
			Lyrics:      lyricsFormatted,
		})

		if err != nil {
			return nil, fmt.Errorf("error occured adding song: %v", err)
		}

		return mapper.SongModelToDomain(&newSong), nil
	} else if err != nil {
		return nil, fmt.Errorf("error occured searching for song: %v", err)
	} else {
		firstSong := &sqlc.Song{
			Songid: songs[0].Songid,
			Title: songs[0].Title,
			Language: songs[0].Language,
			Lyrics: songs[0].Lyrics,
			Isgenerated: songs[0].Isgenerated,
		}	
		return mapper.SongModelToDomain(firstSong), nil
	}
}
