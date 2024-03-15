// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge FROM auth.users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastSeen,
		&i.Disabled,
		&i.DisplayName,
		&i.AvatarUrl,
		&i.Locale,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.EmailVerified,
		&i.PhoneNumberVerified,
		&i.NewEmail,
		&i.OtpMethodLastUsed,
		&i.OtpHash,
		&i.OtpHashExpiresAt,
		&i.DefaultRole,
		&i.IsAnonymous,
		&i.TotpSecret,
		&i.ActiveMfaType,
		&i.Ticket,
		&i.TicketExpiresAt,
		&i.Metadata,
		&i.WebauthnCurrentChallenge,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge FROM auth.users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastSeen,
		&i.Disabled,
		&i.DisplayName,
		&i.AvatarUrl,
		&i.Locale,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.EmailVerified,
		&i.PhoneNumberVerified,
		&i.NewEmail,
		&i.OtpMethodLastUsed,
		&i.OtpHash,
		&i.OtpHashExpiresAt,
		&i.DefaultRole,
		&i.IsAnonymous,
		&i.TotpSecret,
		&i.ActiveMfaType,
		&i.Ticket,
		&i.TicketExpiresAt,
		&i.Metadata,
		&i.WebauthnCurrentChallenge,
	)
	return i, err
}

const getUserRoles = `-- name: GetUserRoles :many
SELECT id, created_at, user_id, role FROM auth.user_roles
WHERE user_id = $1
`

func (q *Queries) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]AuthUserRole, error) {
	rows, err := q.db.Query(ctx, getUserRoles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AuthUserRole
	for rows.Next() {
		var i AuthUserRole
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UserID,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertRefreshtoken = `-- name: InsertRefreshtoken :one
INSERT INTO auth.refresh_tokens (user_id, refresh_token_hash, expires_at, type, metadata)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type InsertRefreshtokenParams struct {
	UserID           uuid.UUID
	RefreshTokenHash pgtype.Text
	ExpiresAt        pgtype.Timestamptz
	Type             RefreshTokenType
	Metadata         []byte
}

func (q *Queries) InsertRefreshtoken(ctx context.Context, arg InsertRefreshtokenParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRefreshtoken,
		arg.UserID,
		arg.RefreshTokenHash,
		arg.ExpiresAt,
		arg.Type,
		arg.Metadata,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertUser = `-- name: InsertUser :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        disabled,
        display_name,
        avatar_url,
        email,
        password_hash,
        ticket,
        ticket_expires_at,
        email_verified,
        locale,
        default_role,
        metadata
    ) VALUES (
      $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
    )
    RETURNING id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($12::TEXT[]) AS roles(role)
RETURNING user_id, (SELECT created_at FROM inserted_user WHERE id = user_id)
`

type InsertUserParams struct {
	Disabled        bool
	DisplayName     string
	AvatarUrl       string
	Email           pgtype.Text
	PasswordHash    pgtype.Text
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
	EmailVerified   bool
	Locale          string
	DefaultRole     string
	Metadata        []byte
	Roles           []string
}

type InsertUserRow struct {
	UserID    uuid.UUID
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error) {
	row := q.db.QueryRow(ctx, insertUser,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.PasswordHash,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.Roles,
	)
	var i InsertUserRow
	err := row.Scan(&i.UserID, &i.CreatedAt)
	return i, err
}

const insertUserWithRefreshToken = `-- name: InsertUserWithRefreshToken :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        disabled,
        display_name,
        avatar_url,
        email,
        password_hash,
        ticket,
        ticket_expires_at,
        email_verified,
        locale,
        default_role,
        metadata,
        last_seen
    ) VALUES (
      $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now()
    )
    RETURNING id, created_at
), inserted_refresh_token AS (
    INSERT INTO auth.refresh_tokens (user_id, refresh_token_hash, expires_at)
        SELECT inserted_user.id, $13, $14
        FROM inserted_user
    RETURNING id AS refresh_token_id
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($12::TEXT[]) AS roles(role)
RETURNING user_id, (SELECT created_at FROM inserted_user WHERE id = user_id)
`

type InsertUserWithRefreshTokenParams struct {
	Disabled              bool
	DisplayName           string
	AvatarUrl             string
	Email                 pgtype.Text
	PasswordHash          pgtype.Text
	Ticket                pgtype.Text
	TicketExpiresAt       pgtype.Timestamptz
	EmailVerified         bool
	Locale                string
	DefaultRole           string
	Metadata              []byte
	Roles                 []string
	RefreshTokenHash      pgtype.Text
	RefreshTokenExpiresAt pgtype.Timestamptz
}

type InsertUserWithRefreshTokenRow struct {
	UserID    uuid.UUID
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) InsertUserWithRefreshToken(ctx context.Context, arg InsertUserWithRefreshTokenParams) (InsertUserWithRefreshTokenRow, error) {
	row := q.db.QueryRow(ctx, insertUserWithRefreshToken,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.PasswordHash,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.Roles,
		arg.RefreshTokenHash,
		arg.RefreshTokenExpiresAt,
	)
	var i InsertUserWithRefreshTokenRow
	err := row.Scan(&i.UserID, &i.CreatedAt)
	return i, err
}

const updateUserChangeEmail = `-- name: UpdateUserChangeEmail :one
UPDATE auth.users
SET (ticket, ticket_expires_at, new_email) = ($2, $3, $4)
WHERE id = $1
RETURNING id, locale, display_name, email
`

type UpdateUserChangeEmailParams struct {
	ID              uuid.UUID
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
	NewEmail        pgtype.Text
}

type UpdateUserChangeEmailRow struct {
	ID          uuid.UUID
	Locale      string
	DisplayName string
	Email       pgtype.Text
}

func (q *Queries) UpdateUserChangeEmail(ctx context.Context, arg UpdateUserChangeEmailParams) (UpdateUserChangeEmailRow, error) {
	row := q.db.QueryRow(ctx, updateUserChangeEmail,
		arg.ID,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.NewEmail,
	)
	var i UpdateUserChangeEmailRow
	err := row.Scan(
		&i.ID,
		&i.Locale,
		&i.DisplayName,
		&i.Email,
	)
	return i, err
}

const updateUserLastSeen = `-- name: UpdateUserLastSeen :one
UPDATE auth.users
SET last_seen = now()
WHERE id = $1
RETURNING last_seen
`

func (q *Queries) UpdateUserLastSeen(ctx context.Context, id uuid.UUID) (pgtype.Timestamptz, error) {
	row := q.db.QueryRow(ctx, updateUserLastSeen, id)
	var last_seen pgtype.Timestamptz
	err := row.Scan(&last_seen)
	return last_seen, err
}

const updateUserTicket = `-- name: UpdateUserTicket :one
UPDATE auth.users
SET (ticket, ticket_expires_at) = ($2, $3)
WHERE id = $1
RETURNING id
`

type UpdateUserTicketParams struct {
	ID              uuid.UUID
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
}

func (q *Queries) UpdateUserTicket(ctx context.Context, arg UpdateUserTicketParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateUserTicket, arg.ID, arg.Ticket, arg.TicketExpiresAt)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
