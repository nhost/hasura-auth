package oauth2

import (
	"context"
	"time"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

const fetchProfileTimeout = 10 * time.Second

type Provider interface {
	GetProfile(ctx context.Context, accessToken string) (oidc.Profile, error)
	Oauth2() *oauth2.Config
}
