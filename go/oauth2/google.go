package oauth2

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

type Google struct {
	*oauth2.Config
}

func NewGoogleProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *Google {
	return &Google{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/google/callback",
			Scopes:       scopes,
			Endpoint:     endpoints.Google,
		},
	}
}

func (g *Google) Oauth2() *oauth2.Config {
	return g.Config
}

type googleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func (g *Google) GetProfile(ctx context.Context, accessToken string) (Profile, error) {
	user, err := fetchOAuthProfile[googleUser](
		ctx,
		"https://www.googleapis.com/oauth2/v2/userinfo",
		accessToken,
		fetchProfileTimeout,
	)
	if err != nil {
		return Profile{}, fmt.Errorf("Google API error: %w", err)
	}

	return Profile{
		ID:            user.ID,
		Email:         user.Email,
		EmailVerified: user.VerifiedEmail,
		Name:          user.Name,
		AvatarURL:     user.Picture,
	}, nil
}
