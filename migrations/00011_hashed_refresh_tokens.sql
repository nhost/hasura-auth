CREATE TABLE "auth"."hashed_refresh_tokens" (
    hashed_refresh_token text NOT NULL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    user_id uuid NOT NULL
);

ALTER TABLE auth.hashed_refresh_tokens
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth.users (id) ON UPDATE CASCADE ON DELETE CASCADE;

COMMENT ON TABLE auth.refresh_tokens IS 'Hashed user refresh tokens. Hasura auth uses them to rotate new access tokens as long as the refresh token is not expired. Don''t modify its structure as Hasura Auth relies on it to function properly.';

COMMENT ON TABLE auth.hashed_refresh_tokens IS 'Deprecated: user refresh tokens. This table is replaced by hashed_refresh_tokens and will be removed in the next major release.';

