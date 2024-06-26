// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: sessions.sql

package sqlcs

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSessions = `-- name: CreateSessions :one
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at) VALUES ($1 , $2, $3, $4, $5, $6, $7) RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionsParams struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
}

func (q *Queries) CreateSessions(ctx context.Context, arg CreateSessionsParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSessions,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSessions = `-- name: GetSessions :one
SELECT username, hashedpassword, full_name, email, password_changed_at, created_at FROM users WHERE 'id' = $1 LIMIT 1
`

func (q *Queries) GetSessions(ctx context.Context, dollar_1 interface{}) (User, error) {
	row := q.db.QueryRowContext(ctx, getSessions, dollar_1)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Hashedpassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
