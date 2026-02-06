package services

import (
	"context"
	"github.com/jackc/pgx/v5"
	"errors"
	"fmt"
	"karoake_assistant/backend/internal/data/mapper"
	"karoake_assistant/backend/internal/data/sqlc"
	"karoake_assistant/backend/internal/domains"
	"karoake_assistant/backend/internal/http/transport"
)

type AuthService struct {
	queries *sqlc.Queries
}

func NewAuthService(queries_ *sqlc.Queries) *AuthService {
	return &AuthService{
		queries: queries_,
	}
}

func (a *AuthService) CreateUser(ctx context.Context, req *transport.CreateUserRequest) (*domains.User, error) {
	if req.Username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	if req.Password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}

	user, err := a.queries.GetUser(ctx, req.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("error occured with querying user: %v", err)
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return mapper.UserModelToDomain(&user), nil
	} else {
		newUser, err := a.queries.CreateUser(ctx, sqlc.CreateUserParams{
			Username: req.Username,
			Password: req.Password,
		})

		if err != nil {
			return nil, fmt.Errorf("error occured with adding new user: %v", err)
		}

		return mapper.UserModelToDomain(&newUser), nil
	}
}

func (a *AuthService) AuthenticateUser(ctx context.Context, req *transport.AuthenticateUserRequest) (*domains.User, error) {
	if req.Username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	if req.Password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}

	user, err := a.queries.GetUser(ctx, req.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("error occured with querying user: %v", err)
	} else if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("user could not be found")
	} else {
		if user.Password != req.Password {
			return nil, fmt.Errorf("passwords do not match")
		}
		return mapper.UserModelToDomain(&user), nil
	}
}
