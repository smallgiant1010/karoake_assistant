package transport

type CreateSongRequest struct {
	Title       string `json:"title"`
	Language    string `json:"langauge"`
	Lyrics      string `json:"lyrics"`
	IsGenerated bool   `json:"bool"`
}
