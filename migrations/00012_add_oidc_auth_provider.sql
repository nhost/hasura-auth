-- start a transaction
BEGIN;

INSERT INTO auth.providers (id) VALUES ('oidc');

-- commit the change (or roll it back later)
COMMIT;

