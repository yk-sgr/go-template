package db

import (
	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

func (u *User) ToUser() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Verified: u.Verified,
	}
}
