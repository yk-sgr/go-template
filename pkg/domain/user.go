package domain

import (
	"context"

	"github.com/google/uuid"
)

const (
	UserMinEmailLength    = 3
	UserMaxEmailLength    = 320
	UserMinPasswordLength = 8
	UserMaxPasswordLength = 128
	UserMinNameLength     = 1
	UserMaxNameLength     = 64
)

// User is a Nexthook user.
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Verified bool      `json:"verified"`
}

type GetUserByIDResponse struct {
	User *User `json:"user"`
}

// UserService is a service for managing users.
type UserService interface {
	GetUserByID(ctx context.Context, uid uuid.UUID, id uuid.UUID) (*User, error)
}
