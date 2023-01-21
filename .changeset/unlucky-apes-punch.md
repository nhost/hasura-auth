---
'hasura-auth': minor
---

Ability to change and verify user phone number

Similar to changing email, user can trigger change of phone number action. The new phone number should be verified on the `verify/otp` endpoint using received via SMS one-time password code. **The verification, does not update or change user session, as changing phone number requires active session.**
