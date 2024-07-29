-- name: GetUser :one
SELECT * FROM auth.users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM auth.users
WHERE email = $1 LIMIT 1;

-- name: GetUserRolesAndOrgs :many
SELECT ur.role, om.organization_id, om.role as organization_role FROM auth.user_roles as ur
LEFT JOIN auth.organizations_members om ON om.user_id = ur.user_id
WHERE ur.user_id = $1;

-- name: GetUserByRefreshTokenHash :one
WITH refresh_token AS (
    SELECT * FROM auth.refresh_tokens
    WHERE refresh_token_hash = $1 AND type = $2 AND expires_at > now()
    LIMIT 1
)
SELECT * FROM auth.users
WHERE id = (SELECT user_id FROM refresh_token) LIMIT 1;

-- name: InsertUser :one
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
    RETURNING *
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest(@roles::TEXT[]) AS roles(role)
RETURNING user_id, (SELECT created_at FROM inserted_user WHERE id = user_id);

-- name: InsertUserWithSecurityKeyAndRefreshToken :one
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
        ($1, @refresh_token_hash, @refresh_token_expires_at)
    RETURNING id AS refresh_token_id
), inserted_security_key AS (
    INSERT INTO auth.user_security_keys
        (user_id, credential_id, credential_public_key, nickname)
    VALUES
        ($1, @credential_id, @credential_public_key, @nickname)
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest(@roles::TEXT[]) AS roles(role)
RETURNING (SELECT refresh_token_id FROM inserted_refresh_token), user_id;

-- name: InsertUserWithSecurityKey :one
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
        ($1, @credential_id, @credential_public_key, @nickname)
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest(@roles::TEXT[]) AS roles(role)
RETURNING user_id;

-- name: InsertUserWithRefreshToken :one
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
        SELECT inserted_user.id, @refresh_token_hash, @refresh_token_expires_at
        FROM inserted_user
    RETURNING id AS refresh_token_id
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest(@roles::TEXT[]) AS roles(role)
RETURNING (SELECT refresh_token_id FROM inserted_refresh_token), user_id;

-- name: InsertRefreshtoken :one
INSERT INTO auth.refresh_tokens (user_id, refresh_token_hash, expires_at, type, metadata)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: RefreshTokenAndGetUserRoles :many

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
),
organizations_members AS (
    SELECT organization_id, user_id, role
    FROM auth.organizations_members
    WHERE user_id IN (SELECT user_id FROM refreshed_token)
)
SELECT r.refresh_token_id, ur.role, om.organization_id, om.role as organization_role
FROM refreshed_token r
LEFT JOIN auth.user_roles ur ON ur.user_id = r.user_id
LEFT JOIN organizations_members om ON ur.user_id = om.user_id;

-- name: UpdateUserLastSeen :one
UPDATE auth.users
SET last_seen = now()
WHERE id = $1
RETURNING last_seen;

-- name: UpdateUserTicket :one
UPDATE auth.users
SET (ticket, ticket_expires_at) = ($2, $3)
WHERE id = $1
RETURNING id;

-- name: UpdateUserChangeEmail :one
UPDATE auth.users
SET (ticket, ticket_expires_at, new_email) = ($2, $3, $4)
WHERE id = $1
RETURNING *;

-- name: CountSecurityKeysUser :one
SELECT COUNT(*) FROM auth.user_security_keys
WHERE user_id = $1;

-- name: UpdateUserDeanonymize :exec
WITH inserted_user AS (
    UPDATE auth.users
    SET
        is_anonymous = false,
        email = @email,
        default_role = @default_role,
        display_name = @display_name,
        locale = @locale,
        metadata = @metadata,
        password_hash = @password_hash,
        ticket = @ticket,
        ticket_expires_at = @ticket_expires_at
    WHERE id = @id
    RETURNING id
)
INSERT INTO auth.user_roles (user_id, role)
    SELECT inserted_user.id, roles.role
    FROM inserted_user, unnest(@roles::TEXT[]) AS roles(role);

-- name: DeleteRefreshTokens :exec
DELETE FROM auth.refresh_tokens
WHERE user_id = $1;

-- name: DeleteUserRoles :exec
DELETE FROM auth.user_roles
WHERE user_id = $1;

-- name: GetOrganizations :many
SELECT * FROM auth.organizations
WHERE id IN (
    SELECT organization_id FROM auth.organizations_members
    WHERE user_id = $1 AND role = 'admin'
);

-- name: InsertOrganization :one
WITH new_organization AS (
    INSERT INTO auth.organizations (name)
    VALUES ($1)
    RETURNING id, name
)
INSERT INTO auth.organizations_members (organization_id, user_id, role)
VALUES ((SELECT id FROM new_organization), $2, 'admin')
RETURNING (SELECT name FROM new_organization), organization_id;

-- name: DeleteOrganization :one
DELETE FROM auth.organizations
WHERE id IN (
    SELECT organization_id FROM auth.organizations_members
    WHERE organization_id = $1 AND user_id = $2 AND role = 'admin'
) RETURNING *;
