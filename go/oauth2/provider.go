package oauth2

import (
	"context"
	"time"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

const fetchProfileTimeout = 10 * time.Second

type Provider interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	GetProfile(ctx context.Context, accessToken string) (oidc.Profile, error)
}
