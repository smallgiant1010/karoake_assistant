package handlers

import (
	"encoding/json"
	"fmt"
	"karoake_assistant/backend/internal/http/middleware"
	"karoake_assistant/backend/internal/http/transport"
	"net/http"
	"strconv"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var request transport.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "body could not be read"}`, http.StatusBadRequest)
		return
	}

	user, err := h.authService.CreateUser(r.Context(), &request)
	if err != nil {
		http.Error(w, `{"error": "failed to create user"}`, http.StatusBadRequest)
		return
	}

	token, err := h.jwtService.GenerateToken(user)
	if err != nil {
		http.Error(w, `{"error": "failed to generate token"}`, http.StatusUnauthorized)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	response := transport.CreateUserResponse{
		Username: user.Username,
		Password: user.Password,
		UserID:   user.UserID,
		Token:    token,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var request transport.AuthenticateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "body could not be read"}`, http.StatusBadRequest)
		return
	}

	user, err := h.authService.AuthenticateUser(r.Context(), &request)
	if err != nil {
		http.Error(w, `{"error": "invalid credentials"}`, http.StatusBadRequest)
		return
	}

	token, err := h.jwtService.GenerateToken(user)
	if err != nil {
		http.Error(w, `{"error": "failed to generate token"}`, http.StatusUnauthorized)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodPost)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)

	response := transport.AuthenticateUserResponse{
		UserID:        user.UserID,
		Username:      user.Username,
		GenerateCount: user.GenerateCount,
		Token:         token,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetJWTClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, `{"error": "unable to obtain token"}`, http.StatusUnauthorized)
		return
	}

	userID, err := strconv.ParseInt(claims.UserID, 10, 32)
	if err != nil {
		http.Error(w, `{"error": "invalid user"}`, http.StatusBadRequest)
		return
	}

	fmt.Printf("%s / HTTP/1.1\n", http.MethodGet)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := transport.UserProfileResponse{
		UserID:        int32(userID),
		Username:      claims.Username,
		GenerateCount: claims.GenerateCount,
	}

	json.NewEncoder(w).Encode(response)
}
