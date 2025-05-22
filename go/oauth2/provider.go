package oauth2

import (
	"context"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type Provider interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	GetProfile(
		ctx context.Context,
		code string,
		idToken *string,
		extra map[string]any,
	) (oidc.Profile, error)
}
