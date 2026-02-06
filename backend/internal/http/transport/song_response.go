package transport

type CreateSongResponse struct {
	SongID          int32  `json:"songID"`
	Title           string `json:"title"`
	Langauge        string `json:"language"`
	Romanticization string `json:"romanticization"`
}
