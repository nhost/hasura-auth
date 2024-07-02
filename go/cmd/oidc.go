package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zitadel/oidc/v3/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v3/pkg/http"
)

func getOIDCProviders(ctx context.Context) (map[string]rp.RelyingParty, error) {
	const clientID = "hxIfXvta0LnGQwSQEXNqsueZ0Fenf5pMtsoCkcMy"
	const clientSecret = "8a5GcUzf3j54JwcfIt1BW5Xwd5vIytcSaa030XLszaRSnDYhtzE6lZjR9bUJcmMTkVcCZxG96vYmY68KV7x879y2YGxOZAc3UiuqWXuOgLW6Za0I51YCRSNOcpSAz6Nw"
	const issuer = "http://localhost:9000/application/o/test/"

	key := []byte("test1234test1234")
	scopes := []string{"email openid profile"}
	redirectURI := "http://localhost:4000/signin/oidc/goauthentik/callback"

	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())
	client := &http.Client{}

	options := []rp.Option{
		// rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
		rp.WithHTTPClient(client),
		// rp.WithLogger(logger),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler)) // TODO: Fix this, most likely encryption breaks the callback
	}

	provider, err := rp.NewRelyingPartyOIDC(ctx, issuer, clientID, clientSecret, redirectURI, scopes, options...)
	if err != nil {
		return nil, fmt.Errorf("error creating provider %w", err)
	}

	return map[string]rp.RelyingParty{
		"goauthentik": provider,
	}, nil
}
