package services

import (
	"context"
	"karaoke_assistant/backend/internal/domains"
	"karaoke_assistant/backend/internal/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(authRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo: authRepo,
	}
}

func (u *UserRequest) CreateUser(ctx context.Context, user *domains.User) (*domains.User, err) {
	return u.repo.CreateUserRow(ctx, user)
}

func (u *UserRequest) AuthenticateUser(ctx context.Context, user *domains.User) (*domains.User, err) {
	return u.repo.GetUserRow(ctx, user)
}

