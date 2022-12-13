---
'hasura-auth': minor
---

Store refresh tokens in SHA256

The existing refresh tokens remains stored in plain text to avoid breaking the ongoing connections
Hashed refresh tokens are stored in a separate `auth.hashed_refresh_tokens` table. Every new token will be automatically stored in this table.

The use of the `auth.refresh_tokens` table is deprecated, and the table will be removed in the future.
