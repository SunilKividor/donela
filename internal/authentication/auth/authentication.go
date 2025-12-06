package auth

import (
	"context"
)

type Authentication interface {
	SignUp(ctx context.Context, name, username, password, role string) (*AuthTokens, error)
	Login(ctx context.Context, username, password string) (*AuthTokens, error)
	Refresh(ctx context.Context, refreshToken string) (*AuthTokens, error)
	ValidateAccessToken(ctx context.Context, token string) (*AuthUser, error)
	Logout(ctx context.Context, id string) error
}
