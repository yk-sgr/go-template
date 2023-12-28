package domain

import (
	"context"
	"time"
)

// SignUpRequest is a request for sign up.
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpResponse is a response for sign up.
type SignUpResponse struct {
}

// SignInRequest is a request for sign in.
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInResponse is a response for sign in.
type SignInResponse struct {
	User        User      `json:"user"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// AuthService is a service for managing authentication.
type AuthService interface {
	SignUp(ctx context.Context, dto *SignUpRequest) (*SignUpResponse, error)
	SignIn(ctx context.Context, dto *SignInRequest) (*SignInResponse, error)
	ValidateToken(ctx context.Context, token string) (*User, error)
}
