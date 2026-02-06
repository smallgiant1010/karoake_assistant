package handlers

import (
	"encoding/json"
	"fmt"
	"karoake_assistant/backend/internal/http/transport"
	"net/http"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request transport.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "body could not be read", http.StatusBadRequest)
		return
	}

	user, err := h.authService.CreateUser(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	response := transport.CreateUserResponse{
		Username: user.Username,
		Password: user.Password,
		UserID:   user.UserID,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request transport.AuthenticateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "body could not be read", http.StatusBadRequest)
		return
	}

	user, err := h.authService.AuthenticateUser(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	response := transport.AuthenticateUserResponse{
		UserID:        user.UserID,
		Username:      user.Username,
		GenerateCount: user.GenerateCount,
	}

	json.NewEncoder(w).Encode(response)
}
