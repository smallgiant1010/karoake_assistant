package auth

import (
	"fmt"
	"karoake_assistant/backend/internal/domains"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
	issuer string
	expiry time.Duration
}

func NewJWTService(secret_ string, expiryhrs int16) *JWTService {
	return &JWTService{
		secret: []byte(secret_),
		issuer: "karaoke_assistant",
		expiry: time.Hour * time.Duration(expiryhrs),
	}
}

func (s *JWTService) GenerateToken(user *domains.User) (string, error) {
	now := time.Now()
	expiresAt := now.Add(s.expiry)
	user_id := strconv.Itoa(int(user.UserID))

	claims := &JWTClaims{
		UserID:   user_id,
		Username: user.Username,
		GenerateCount: user.GenerateCount,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    s.issuer,
			Subject:   user_id,
		},
	}

	// creates a new token using sha256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("error occured signing token: %w", err)
	}

	return tokenStr, nil
}

func (s *JWTService) ValidateToken(tokenStr string) (*JWTClaims, error) {
	// attempt to parse the token
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		// verify signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method incorrect: %v", t.Header["alg"])
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// validate the token
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	// check who issued the token
	if claims.Issuer != s.issuer {
		return nil, fmt.Errorf("invalid token issuer")
	}

	return claims, nil
}
