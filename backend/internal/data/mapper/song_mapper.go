package mapper

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"karoake_assistant/backend/internal/data/sqlc"
	"karoake_assistant/backend/internal/domains"
)

func SongModelToDomain(model *sqlc.Song) *domains.Song {
	song, err := domains.NewSong(
		int32(model.Songid),
		model.Language,
		model.Title,
		model.Isgenerated.Bool,
		model.Lyrics.String,
	)

	if err != nil {
		return nil
	}

	return song
}

func SongDomainToModel(domain *domains.Song) *sqlc.Song {
	var isGenerated pgtype.Bool
	if err := isGenerated.Scan(domain.IsGenerated); err != nil {
		fmt.Printf("error occured casting IsGenerated as pgtype: %v", err)
		return nil
	}

	var lyrics pgtype.Text
	if err := lyrics.Scan(domain.Lyrics); err != nil {
		fmt.Printf("error occured casting lyrics as pgtype: %v", err)
		return nil
	}

	return &sqlc.Song{
		Songid:      int64(domain.SongID),
		Language:    domain.Language,
		Title:       domain.Title,
		Isgenerated: isGenerated,
		Lyrics:      lyrics,
	}
}
