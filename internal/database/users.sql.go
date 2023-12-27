// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(Email,passwd,id,red,created_at) VALUES($1,$2,$3,$4,$5)
RETURNING email, passwd, id, red, created_at
`

type CreateUserParams struct {
	Email     string
	Passwd    []byte
	ID        uuid.UUID
	Red       sql.NullBool
	CreatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.Passwd,
		arg.ID,
		arg.Red,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.Email,
		&i.Passwd,
		&i.ID,
		&i.Red,
		&i.CreatedAt,
	)
	return i, err
}