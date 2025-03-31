package oauth2

import (
	"context"
	"fmt"
	"strconv"

	"golang.org/x/oauth2"
)

type Github struct {
	profileURL string
	*oauth2.Config
}

func NewGithubProvider(
	clientID, clientSecret, authServerURL string,
	authURL, tokenURL, profileURL string,
	scopes []string,
) *Github {
	return &Github{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/github/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{ //nolint:exhaustruct
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
		profileURL: profileURL,
	}
}

func (g *Github) Oauth2() *oauth2.Config {
	return g.Config
}

type gitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func (g *Github) GetProfile(ctx context.Context, accessToken string) (Profile, error) {
	user, err := fetchOAuthProfile[gitHubUser](
		ctx,
		g.profileURL,
		accessToken,
		fetchProfileTimeout,
	)
	if err != nil {
		return Profile{}, fmt.Errorf("GitHub API error: %w", err)
	}

	return Profile{
		ID:            strconv.Itoa(user.ID),
		Email:         user.Email,
		EmailVerified: user.Email != "",
		Name:          user.Name,
		AvatarURL:     user.AvatarURL,
	}, nil
}
