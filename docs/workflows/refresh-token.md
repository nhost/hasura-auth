# Refresh tokens

```mermaid
sequenceDiagram
	autonumber
	actor U as User
	participant A as Hasura Auth
	U->>+A: HTTP POST /token
	A->>A: Create new refresh token
	A->>A: Delete previous refresh token
	A->>-U: HTTP OK response
	Note left of A: Refresh token + access token
```
