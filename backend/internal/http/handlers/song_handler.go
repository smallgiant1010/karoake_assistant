package handlers

import (
	"encoding/json"
	"fmt"
	"karoake_assistant/backend/internal/http/transport"
	"net/http"
)

func (h *Handler) Romanticize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request transport.CreateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "body could not be read", http.StatusBadRequest)
		return
	}

	song, err := h.songService.RomanticizeSong(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	response := transport.CreateSongResponse{
		SongID:          song.SongID,
		Title:           song.Title,
		Langauge:        song.Language,
		Romanticization: song.Lyrics,
	}
	json.NewEncoder(w).Encode(response)
}
