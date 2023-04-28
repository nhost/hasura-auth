BEGIN;
ALTER TABLE auth.refresh_tokens
RENAME COLUMN refresh_token TO id;
COMMENT ON COLUMN auth.refresh_tokens.id IS '';
ALTER TABLE auth.refresh_tokens ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE auth.refresh_tokens ALTER COLUMN refresh_token_hash DROP EXPRESSION;
COMMIT;