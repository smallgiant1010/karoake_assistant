package domains

import (
	"errors"
)

type Song struct {
	Lyrics string
}

func NewSong(lyrics string) (*Song, error) {
	if lyrics == "" {
		return nil, errors.New("lyrics cannot be empty")
	}

	return &Song{Lyrics: lyrics}, nil
}
