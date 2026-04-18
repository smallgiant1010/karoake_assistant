package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type JWTClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	GenerateCount int32 `json:"generateCount"`
	jwt.RegisteredClaims
}

func NewJWTClaims(userID int32, username string, generateCount_ int32, expiresAt time.Time) *JWTClaims {
	return &JWTClaims{
		UserID:   strconv.Itoa(int(userID)),
		Username: username,
		GenerateCount: generateCount_,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (c *JWTClaims) GetUserID() string { return c.UserID }
func (c *JWTClaims) GetUsername() string { return c.Username }
func (c *JWTClaims) GetGenerateCount() int32 { return c.GenerateCount }
