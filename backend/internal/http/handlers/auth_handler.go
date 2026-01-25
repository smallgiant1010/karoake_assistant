package handlers

import (
	"net/http"
	"encoding/json"
	"karaoke_assistant/backend/internal/http/transport"
	"karaoke_assistant/backend/internal/domains"
	"karaoke_assistant/backend/internal/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service_ *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service_,
	}
}

func (h* AuthHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request transport.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "body could not be read", http.StatusBadRequest)
		return
	}

	user, err := domains.NewUser(request.Username, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Finish After Creating Repository
}
