package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `db:"id"`
	DisplayName string    `db:"display_name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Version     int       `db:"version"`
}

type CreateUserRequest struct {
	DisplayName string
}

type UpdateUserRequest struct {
	ID          uuid.UUID
	DisplayName string
}
