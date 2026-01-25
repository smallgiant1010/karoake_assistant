package repository

import (
	"karaoke_assistant/backend/internal/services"
	"karaoke_assistant/backend/internal/domains"
	"context"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUserRow(ctx context.Context, username string, password string) (*domains.User, error)
	GetUserRow(context.Context, username string, password string) (*domains.User, error)
}

type AuthRepository struct {
	database: *pgx.Conn
}

func NewAuthRepository(database_ *pgx.Conn) *AuthRepository {
	return &AuthRepository{
		database: database_,
	}
}

func (u *AuthRepository) CreateUserRow(ctx context.Context, username string, password string) (*domains.User, error) {
}

func (u *AuthRepository) GetUserRow(ctx context.Context, username string, password string) (*domains.User, error) {
}
