---
'hasura-auth': patch
---

Added support for OpenID Connect auth provider

Tested with [Keycloak](http://keycloak.org/) but other OIDC providers should be working as well. It uses Authorization Code Flow.
In addition you can enable PKCE (Proof Key for Code Exchange) via the env variable `AUTH_PROVIDER_OIDC_PKCE`.

