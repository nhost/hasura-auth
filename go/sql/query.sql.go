// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const countSecurityKeysUser = `-- name: CountSecurityKeysUser :one
SELECT COUNT(*) FROM auth.user_security_keys
WHERE user_id = $1
`

func (q *Queries) CountSecurityKeysUser(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countSecurityKeysUser, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteRefreshTokens = `-- name: DeleteRefreshTokens :exec
DELETE FROM auth.refresh_tokens
WHERE user_id = $1
`

func (q *Queries) DeleteRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteRefreshTokens, userID)
	return err
}

const deleteUserRoles = `-- name: DeleteUserRoles :exec
DELETE FROM auth.user_roles
WHERE user_id = $1
`

func (q *Queries) DeleteUserRoles(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserRoles, userID)
	return err
}

const findUserProviderByProviderId = `-- name: FindUserProviderByProviderId :one
SELECT id, created_at, updated_at, user_id, access_token, refresh_token, provider_id, provider_user_id FROM auth.user_providers
WHERE provider_user_id = $1 AND provider_id = $2
`

type FindUserProviderByProviderIdParams struct {
	ProviderUserID string
	ProviderID     string
}

func (q *Queries) FindUserProviderByProviderId(ctx context.Context, arg FindUserProviderByProviderIdParams) (AuthUserProvider, error) {
	row := q.db.QueryRow(ctx, findUserProviderByProviderId, arg.ProviderUserID, arg.ProviderID)
	var i AuthUserProvider
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.ProviderID,
		&i.ProviderUserID,
	)
	return i, err
}

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

const getUserByProviderID = `-- name: GetUserByProviderID :one
WITH user_providers AS (
    SELECT id, created_at, updated_at, user_id, access_token, refresh_token, provider_id, provider_user_id FROM auth.user_providers
    WHERE provider_user_id = $1
    AND provider_id = $2
    LIMIT 1
)
SELECT id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge FROM auth.users
WHERE id = (SELECT user_id FROM user_providers) LIMIT 1
`

type GetUserByProviderIDParams struct {
	ProviderUserID string
	ProviderID     string
}

func (q *Queries) GetUserByProviderID(ctx context.Context, arg GetUserByProviderIDParams) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByProviderID, arg.ProviderUserID, arg.ProviderID)
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

const getUserByRefreshTokenHash = `-- name: GetUserByRefreshTokenHash :one
WITH refresh_token AS (
    SELECT id, created_at, expires_at, user_id, metadata, type, refresh_token_hash FROM auth.refresh_tokens
    WHERE refresh_token_hash = $1 AND type = $2 AND expires_at > now()
    LIMIT 1
)
SELECT id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge FROM auth.users
WHERE id = (SELECT user_id FROM refresh_token) LIMIT 1
`

type GetUserByRefreshTokenHashParams struct {
	RefreshTokenHash pgtype.Text
	Type             RefreshTokenType
}

func (q *Queries) GetUserByRefreshTokenHash(ctx context.Context, arg GetUserByRefreshTokenHashParams) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByRefreshTokenHash, arg.RefreshTokenHash, arg.Type)
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

const getUserByTicket = `-- name: GetUserByTicket :one
WITH selected_user AS (
    SELECT id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge FROM auth.users
    WHERE ticket = $1  AND ticket_expires_at > now()
    LIMIT 1
)
UPDATE auth.users
SET ticket = NULL, ticket_expires_at = now()
WHERE id = (SELECT id FROM selected_user)
RETURNING id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge
`

func (q *Queries) GetUserByTicket(ctx context.Context, dollar_1 pgtype.Text) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByTicket, dollar_1)
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
        id,
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
      $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
    )
    RETURNING id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($13::TEXT[]) AS roles(role)
RETURNING user_id, (SELECT created_at FROM inserted_user WHERE id = user_id)
`

type InsertUserParams struct {
	ID              uuid.UUID
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
		arg.ID,
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
    RETURNING id
), inserted_refresh_token AS (
    INSERT INTO auth.refresh_tokens (user_id, refresh_token_hash, expires_at)
        SELECT inserted_user.id, $12, $13
        FROM inserted_user
    RETURNING id, user_id
), inserted_user_role AS (
    INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($14::TEXT[]) AS roles(role)
)
SELECT u.id, u.created_at, u.updated_at, u.last_seen, u.disabled, u.display_name, u.avatar_url, u.locale, u.email, u.phone_number, u.password_hash, u.email_verified, u.phone_number_verified, u.new_email, u.otp_method_last_used, u.otp_hash, u.otp_hash_expires_at, u.default_role, u.is_anonymous, u.totp_secret, u.active_mfa_type, u.ticket, u.ticket_expires_at, u.metadata, u.webauthn_current_challenge, r.id AS refresh_token_id
FROM auth.users u
JOIN inserted_refresh_token r
ON u.id = r.user_id
WHERE u.id = (SELECT id FROM inserted_user)
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
	RefreshTokenHash      pgtype.Text
	RefreshTokenExpiresAt pgtype.Timestamptz
	Roles                 []string
}

type InsertUserWithRefreshTokenRow struct {
	AuthUser       AuthUser
	RefreshTokenID uuid.UUID
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
		arg.RefreshTokenHash,
		arg.RefreshTokenExpiresAt,
		arg.Roles,
	)
	var i InsertUserWithRefreshTokenRow
	err := row.Scan(
		&i.AuthUser.ID,
		&i.AuthUser.CreatedAt,
		&i.AuthUser.UpdatedAt,
		&i.AuthUser.LastSeen,
		&i.AuthUser.Disabled,
		&i.AuthUser.DisplayName,
		&i.AuthUser.AvatarUrl,
		&i.AuthUser.Locale,
		&i.AuthUser.Email,
		&i.AuthUser.PhoneNumber,
		&i.AuthUser.PasswordHash,
		&i.AuthUser.EmailVerified,
		&i.AuthUser.PhoneNumberVerified,
		&i.AuthUser.NewEmail,
		&i.AuthUser.OtpMethodLastUsed,
		&i.AuthUser.OtpHash,
		&i.AuthUser.OtpHashExpiresAt,
		&i.AuthUser.DefaultRole,
		&i.AuthUser.IsAnonymous,
		&i.AuthUser.TotpSecret,
		&i.AuthUser.ActiveMfaType,
		&i.AuthUser.Ticket,
		&i.AuthUser.TicketExpiresAt,
		&i.AuthUser.Metadata,
		&i.AuthUser.WebauthnCurrentChallenge,
		&i.RefreshTokenID,
	)
	return i, err
}

const insertUserWithSecurityKey = `-- name: InsertUserWithSecurityKey :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        id,
        disabled,
        display_name,
        avatar_url,
        email,
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
    RETURNING id
), inserted_security_key AS (
    INSERT INTO auth.user_security_keys
        (user_id, credential_id, credential_public_key, nickname)
    VALUES
        ($1, $13, $14, $15)
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($12::TEXT[]) AS roles(role)
RETURNING user_id
`

type InsertUserWithSecurityKeyParams struct {
	ID                  uuid.UUID
	Disabled            bool
	DisplayName         string
	AvatarUrl           string
	Email               pgtype.Text
	Ticket              pgtype.Text
	TicketExpiresAt     pgtype.Timestamptz
	EmailVerified       bool
	Locale              string
	DefaultRole         string
	Metadata            []byte
	Roles               []string
	CredentialID        string
	CredentialPublicKey []byte
	Nickname            pgtype.Text
}

func (q *Queries) InsertUserWithSecurityKey(ctx context.Context, arg InsertUserWithSecurityKeyParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertUserWithSecurityKey,
		arg.ID,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.Roles,
		arg.CredentialID,
		arg.CredentialPublicKey,
		arg.Nickname,
	)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const insertUserWithSecurityKeyAndRefreshToken = `-- name: InsertUserWithSecurityKeyAndRefreshToken :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        id,
        disabled,
        display_name,
        avatar_url,
        email,
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
    RETURNING id
), inserted_refresh_token AS (
    INSERT INTO auth.refresh_tokens
        (user_id, refresh_token_hash, expires_at)
    VALUES
        ($1, $12, $13)
    RETURNING id, user_id
), inserted_security_key AS (
    INSERT INTO auth.user_security_keys
        (user_id, credential_id, credential_public_key, nickname)
    VALUES
        ($1, $14, $15, $16)
), inserted_user_role AS (
    INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($17::TEXT[]) AS roles(role)
)
SELECT u.id, u.created_at, u.updated_at, u.last_seen, u.disabled, u.display_name, u.avatar_url, u.locale, u.email, u.phone_number, u.password_hash, u.email_verified, u.phone_number_verified, u.new_email, u.otp_method_last_used, u.otp_hash, u.otp_hash_expires_at, u.default_role, u.is_anonymous, u.totp_secret, u.active_mfa_type, u.ticket, u.ticket_expires_at, u.metadata, u.webauthn_current_challenge, r.id AS refresh_token_id
FROM auth.users u
JOIN inserted_refresh_token r
ON u.id = r.user_id
WHERE u.id = (SELECT id FROM inserted_user)
`

type InsertUserWithSecurityKeyAndRefreshTokenParams struct {
	ID                    uuid.UUID
	Disabled              bool
	DisplayName           string
	AvatarUrl             string
	Email                 pgtype.Text
	Ticket                pgtype.Text
	TicketExpiresAt       pgtype.Timestamptz
	EmailVerified         bool
	Locale                string
	DefaultRole           string
	Metadata              []byte
	RefreshTokenHash      pgtype.Text
	RefreshTokenExpiresAt pgtype.Timestamptz
	CredentialID          string
	CredentialPublicKey   []byte
	Nickname              pgtype.Text
	Roles                 []string
}

type InsertUserWithSecurityKeyAndRefreshTokenRow struct {
	AuthUser       AuthUser
	RefreshTokenID uuid.UUID
}

func (q *Queries) InsertUserWithSecurityKeyAndRefreshToken(ctx context.Context, arg InsertUserWithSecurityKeyAndRefreshTokenParams) (InsertUserWithSecurityKeyAndRefreshTokenRow, error) {
	row := q.db.QueryRow(ctx, insertUserWithSecurityKeyAndRefreshToken,
		arg.ID,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.RefreshTokenHash,
		arg.RefreshTokenExpiresAt,
		arg.CredentialID,
		arg.CredentialPublicKey,
		arg.Nickname,
		arg.Roles,
	)
	var i InsertUserWithSecurityKeyAndRefreshTokenRow
	err := row.Scan(
		&i.AuthUser.ID,
		&i.AuthUser.CreatedAt,
		&i.AuthUser.UpdatedAt,
		&i.AuthUser.LastSeen,
		&i.AuthUser.Disabled,
		&i.AuthUser.DisplayName,
		&i.AuthUser.AvatarUrl,
		&i.AuthUser.Locale,
		&i.AuthUser.Email,
		&i.AuthUser.PhoneNumber,
		&i.AuthUser.PasswordHash,
		&i.AuthUser.EmailVerified,
		&i.AuthUser.PhoneNumberVerified,
		&i.AuthUser.NewEmail,
		&i.AuthUser.OtpMethodLastUsed,
		&i.AuthUser.OtpHash,
		&i.AuthUser.OtpHashExpiresAt,
		&i.AuthUser.DefaultRole,
		&i.AuthUser.IsAnonymous,
		&i.AuthUser.TotpSecret,
		&i.AuthUser.ActiveMfaType,
		&i.AuthUser.Ticket,
		&i.AuthUser.TicketExpiresAt,
		&i.AuthUser.Metadata,
		&i.AuthUser.WebauthnCurrentChallenge,
		&i.RefreshTokenID,
	)
	return i, err
}

const insertUserWithUserProvider = `-- name: InsertUserWithUserProvider :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        id,
        disabled,
        display_name,
        avatar_url,
        email,
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
    RETURNING id
), inserted_user_provider AS (
    INSERT INTO auth.user_providers
        (user_id, provider_id, provider_user_id)
    VALUES
        ($1, $13, $14)
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($12::TEXT[]) AS roles(role)
RETURNING user_id
`

type InsertUserWithUserProviderParams struct {
	ID              uuid.UUID
	Disabled        bool
	DisplayName     string
	AvatarUrl       string
	Email           pgtype.Text
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
	EmailVerified   bool
	Locale          string
	DefaultRole     string
	Metadata        []byte
	Roles           []string
	ProviderID      string
	ProviderUserID  string
}

func (q *Queries) InsertUserWithUserProvider(ctx context.Context, arg InsertUserWithUserProviderParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertUserWithUserProvider,
		arg.ID,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.Roles,
		arg.ProviderID,
		arg.ProviderUserID,
	)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const insertUserWithUserProviderAndRefreshToken = `-- name: InsertUserWithUserProviderAndRefreshToken :one
WITH inserted_user AS (
    INSERT INTO auth.users (
        id,
        disabled,
        display_name,
        avatar_url,
        email,
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
    RETURNING id
), inserted_refresh_token AS (
    INSERT INTO auth.refresh_tokens
        (user_id, refresh_token_hash, expires_at)
    VALUES
        ($1, $12, $13)
    RETURNING id , user_id
), inserted_user_provider AS (
    INSERT INTO auth.user_providers
        (user_id, provider_id, provider_user_id)
    VALUES
        ($1, $14, $15)
), inserted_user_role AS (
    INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($16::TEXT[]) AS roles(role)
)
SELECT u.id, u.created_at, u.updated_at, u.last_seen, u.disabled, u.display_name, u.avatar_url, u.locale, u.email, u.phone_number, u.password_hash, u.email_verified, u.phone_number_verified, u.new_email, u.otp_method_last_used, u.otp_hash, u.otp_hash_expires_at, u.default_role, u.is_anonymous, u.totp_secret, u.active_mfa_type, u.ticket, u.ticket_expires_at, u.metadata, u.webauthn_current_challenge, r.id AS refresh_token_id
FROM auth.users u
JOIN inserted_refresh_token r
ON u.id = r.user_id
WHERE u.id = (SELECT id FROM inserted_user)
`

type InsertUserWithUserProviderAndRefreshTokenParams struct {
	ID                    uuid.UUID
	Disabled              bool
	DisplayName           string
	AvatarUrl             string
	Email                 pgtype.Text
	Ticket                pgtype.Text
	TicketExpiresAt       pgtype.Timestamptz
	EmailVerified         bool
	Locale                string
	DefaultRole           string
	Metadata              []byte
	RefreshTokenHash      pgtype.Text
	RefreshTokenExpiresAt pgtype.Timestamptz
	ProviderID            string
	ProviderUserID        string
	Roles                 []string
}

type InsertUserWithUserProviderAndRefreshTokenRow struct {
	AuthUser       AuthUser
	RefreshTokenID uuid.UUID
}

func (q *Queries) InsertUserWithUserProviderAndRefreshToken(ctx context.Context, arg InsertUserWithUserProviderAndRefreshTokenParams) (InsertUserWithUserProviderAndRefreshTokenRow, error) {
	row := q.db.QueryRow(ctx, insertUserWithUserProviderAndRefreshToken,
		arg.ID,
		arg.Disabled,
		arg.DisplayName,
		arg.AvatarUrl,
		arg.Email,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.EmailVerified,
		arg.Locale,
		arg.DefaultRole,
		arg.Metadata,
		arg.RefreshTokenHash,
		arg.RefreshTokenExpiresAt,
		arg.ProviderID,
		arg.ProviderUserID,
		arg.Roles,
	)
	var i InsertUserWithUserProviderAndRefreshTokenRow
	err := row.Scan(
		&i.AuthUser.ID,
		&i.AuthUser.CreatedAt,
		&i.AuthUser.UpdatedAt,
		&i.AuthUser.LastSeen,
		&i.AuthUser.Disabled,
		&i.AuthUser.DisplayName,
		&i.AuthUser.AvatarUrl,
		&i.AuthUser.Locale,
		&i.AuthUser.Email,
		&i.AuthUser.PhoneNumber,
		&i.AuthUser.PasswordHash,
		&i.AuthUser.EmailVerified,
		&i.AuthUser.PhoneNumberVerified,
		&i.AuthUser.NewEmail,
		&i.AuthUser.OtpMethodLastUsed,
		&i.AuthUser.OtpHash,
		&i.AuthUser.OtpHashExpiresAt,
		&i.AuthUser.DefaultRole,
		&i.AuthUser.IsAnonymous,
		&i.AuthUser.TotpSecret,
		&i.AuthUser.ActiveMfaType,
		&i.AuthUser.Ticket,
		&i.AuthUser.TicketExpiresAt,
		&i.AuthUser.Metadata,
		&i.AuthUser.WebauthnCurrentChallenge,
		&i.RefreshTokenID,
	)
	return i, err
}

const refreshTokenAndGetUserRoles = `-- name: RefreshTokenAndGetUserRoles :many
WITH refreshed_token AS (
    UPDATE auth.refresh_tokens
    SET expires_at = $2
    WHERE refresh_token_hash = $1
    RETURNING id AS refresh_token_id, user_id
),
updated_user AS (
    UPDATE auth.users
    SET last_seen = now()
    FROM refreshed_token
    WHERE auth.users.id = refreshed_token.user_id
)
SELECT refreshed_token.refresh_token_id, role FROM auth.user_roles
RIGHT JOIN refreshed_token ON auth.user_roles.user_id = refreshed_token.user_id
`

type RefreshTokenAndGetUserRolesParams struct {
	RefreshTokenHash pgtype.Text
	ExpiresAt        pgtype.Timestamptz
}

type RefreshTokenAndGetUserRolesRow struct {
	RefreshTokenID uuid.UUID
	Role           pgtype.Text
}

func (q *Queries) RefreshTokenAndGetUserRoles(ctx context.Context, arg RefreshTokenAndGetUserRolesParams) ([]RefreshTokenAndGetUserRolesRow, error) {
	rows, err := q.db.Query(ctx, refreshTokenAndGetUserRoles, arg.RefreshTokenHash, arg.ExpiresAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RefreshTokenAndGetUserRolesRow
	for rows.Next() {
		var i RefreshTokenAndGetUserRolesRow
		if err := rows.Scan(&i.RefreshTokenID, &i.Role); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserChangeEmail = `-- name: UpdateUserChangeEmail :one
UPDATE auth.users
SET (ticket, ticket_expires_at, new_email) = ($2, $3, $4)
WHERE id = $1
RETURNING id, created_at, updated_at, last_seen, disabled, display_name, avatar_url, locale, email, phone_number, password_hash, email_verified, phone_number_verified, new_email, otp_method_last_used, otp_hash, otp_hash_expires_at, default_role, is_anonymous, totp_secret, active_mfa_type, ticket, ticket_expires_at, metadata, webauthn_current_challenge
`

type UpdateUserChangeEmailParams struct {
	ID              uuid.UUID
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
	NewEmail        pgtype.Text
}

func (q *Queries) UpdateUserChangeEmail(ctx context.Context, arg UpdateUserChangeEmailParams) (AuthUser, error) {
	row := q.db.QueryRow(ctx, updateUserChangeEmail,
		arg.ID,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.NewEmail,
	)
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

const updateUserChangePassword = `-- name: UpdateUserChangePassword :one
UPDATE auth.users
SET password_hash = $2
WHERE id = $1
RETURNING id
`

type UpdateUserChangePasswordParams struct {
	ID           uuid.UUID
	PasswordHash pgtype.Text
}

func (q *Queries) UpdateUserChangePassword(ctx context.Context, arg UpdateUserChangePasswordParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateUserChangePassword, arg.ID, arg.PasswordHash)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const updateUserDeanonymize = `-- name: UpdateUserDeanonymize :exec
WITH inserted_user AS (
    UPDATE auth.users
    SET
        is_anonymous = false,
        email = $2,
        default_role = $3,
        display_name = $4,
        locale = $5,
        metadata = $6,
        password_hash = $7,
        ticket = $8,
        ticket_expires_at = $9
    WHERE id = $10
    RETURNING id
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest($1::TEXT[]) AS roles(role)
`

type UpdateUserDeanonymizeParams struct {
	Roles           []string
	Email           interface{}
	DefaultRole     pgtype.Text
	DisplayName     pgtype.Text
	Locale          pgtype.Text
	Metadata        []byte
	PasswordHash    pgtype.Text
	Ticket          pgtype.Text
	TicketExpiresAt pgtype.Timestamptz
	ID              pgtype.UUID
}

func (q *Queries) UpdateUserDeanonymize(ctx context.Context, arg UpdateUserDeanonymizeParams) error {
	_, err := q.db.Exec(ctx, updateUserDeanonymize,
		arg.Roles,
		arg.Email,
		arg.DefaultRole,
		arg.DisplayName,
		arg.Locale,
		arg.Metadata,
		arg.PasswordHash,
		arg.Ticket,
		arg.TicketExpiresAt,
		arg.ID,
	)
	return err
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
