-- start a transaction
BEGIN;
INSERT INTO auth.providers (id)
    VALUES ('oidc')
ON CONFLICT
    DO NOTHING;
COMMIT;
