package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yk-sgr/nexthook-api/internal/db"
	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) domain.UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, uid uuid.UUID, id uuid.UUID) (*domain.User, error) {
	if uid != id {
		return nil, domain.NewForbiddenError()
	}

	u, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.ToUser(), nil
}
