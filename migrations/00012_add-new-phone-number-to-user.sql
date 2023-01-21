-- start a transaction
BEGIN;
ALTER TABLE auth.users
ADD COLUMN new_phone_number text;
COMMIT;