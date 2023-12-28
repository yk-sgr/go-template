// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SchemaMigration struct {
	Version string
}

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Password   string
	Verified   bool
	LastSeenAt pgtype.Timestamptz
	CreatedAt  pgtype.Timestamptz
}
