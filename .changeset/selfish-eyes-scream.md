---
'hasura-auth': patch
---

Option to resend verification code

On phone number change user can request new verification code on `/user/phone-number/resend-verification`.
For this purpose there must be an active session. On sign in when user need to request new verification,
`/signin/passwordless/sms/resend` must be used.
