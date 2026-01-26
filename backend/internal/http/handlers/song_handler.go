package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"karaoke_assistant/backend/internal/http/transport"
	"karaoke_assistant/backend/internal/domains"
	"karaoke_assistant/backend/internal/services"
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
		Title: song.Title,
		Language: song.Language,
		Romanticization: song.Lyrics,
	}
	json.NewEncoder(w).Encode(response)
}


