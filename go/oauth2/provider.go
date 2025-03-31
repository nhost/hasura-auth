package oauth2

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

const fetchProfileTimeout = 10 * time.Second

type Provider interface {
	GetProfile(ctx context.Context, accessToken string) (Profile, error)
	Oauth2() *oauth2.Config
}

type Profile struct {
	ID            string
	Email         string
	EmailVerified bool
	Name          string
	AvatarURL     string
}
