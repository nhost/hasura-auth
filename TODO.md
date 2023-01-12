- only import webauthn when it's activated
- only import nodemailer and **email-templates** when needed
  - even better: get rid of `email-templates`
- reduce the number of SQL queries
- async emails

Other things, with less priority (higher cost/benefit)

- Joi
- replace Winston by something else
- replace qrcode by? Or import it only when needed
- only when needed
  - qrcode
  - lib phone numbers
  - hibp (mini impact)
  - grant
- replace express?
- do we need express helmet?
-
