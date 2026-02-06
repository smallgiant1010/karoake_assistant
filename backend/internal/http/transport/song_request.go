package transport

type CreateSongRequest struct {
	Title       string `json:"title"`
	Language    string `json:"language"`
	Lyrics      string `json:"lyrics"`
	IsGenerated bool   `json:"isGenerated"`
}
