---
'hasura-auth': patch
---

Store refresh tokens in SHA256

The existing refresh tokens remains stored in plain text to avoid breaking the ongoing connections
