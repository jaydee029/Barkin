// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Email     string
	Passwd    []byte
	ID        uuid.UUID
	Red       sql.NullBool
	CreatedAt time.Time
}