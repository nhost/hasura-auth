---
'hasura-auth': patch
---

Improve the handling of Grant errors

When Grant fails to handle an Oauth query, it now logs and sends the error description.
