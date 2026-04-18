package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"karoake_assistant/backend/internal/auth"
)

type contextKey string

const jwtClaimsKey contextKey = "jwt:claims"

func JWTContextMiddleware(jwtService *auth.JWTService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "no token detected"}`, http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, `{"error": "invalid bearer string provided"}`, http.StatusUnauthorized)
			return
		}

		claims, err := jwtService.ValidateToken(parts[1])
		if err != nil {
			http.Error(w, `{"error": "could not validate token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwtClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetJWTClaimsFromContext(ctx context.Context) (*auth.JWTClaims, bool) {
	claims, ok := ctx.Value(jwtClaimsKey).(*auth.JWTClaims)
	return claims, ok
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
