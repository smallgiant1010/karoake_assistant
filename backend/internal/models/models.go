package domain

type Artist struct {
	ArtistID  string `json:"artist_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Song struct {
	SongID  string   `json:"song_id"`
	Artists []Artist `json:"artists"`
	Genre   string   `json:"genre"`
	Lyrics  []string `json:"lyrics"`
}
