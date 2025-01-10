// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Shepot struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	UserID    uuid.UUID
}

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}
