package domains

import (
	"errors"
)

type Song struct {
	SongID      int32
	Language    string
	Title       string
	IsGenerated bool
	Lyrics      string
}

func NewSong(songID int32, language string, title string, isGenerated bool, lyrics string) (*Song, error) {
	if lyrics == "" {
		return nil, errors.New("lyrics cannot be empty")
	}

	return &Song{
		SongID:      songID,
		Language:    language,
		Title:       title,
		IsGenerated: isGenerated,
		Lyrics:      lyrics,
	}, nil
}
