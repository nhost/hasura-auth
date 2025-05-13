package oauth2

import (
	"context"
	"fmt"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type Twitch struct {
	*oauth2.Config
}

func NewTwitchProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *Twitch {
	return &Twitch{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/twitch/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{ //nolint:exhaustruct
				AuthURL:  "https://id.twitch.tv/oauth2/authorize",
				TokenURL: "https://id.twitch.tv/oauth2/token",
			},
		},
	}
}

type twitchUserProfile struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

func (t *Twitch) GetProfile(
	ctx context.Context,
	accessToken string,
	_ *string,
	_ map[string]any,
) (oidc.Profile, error) {
	var userProfile twitchUserProfile
	if err := fetchOAuthProfile(
		ctx,
		"https://api.twitch.tv/helix/users",
		accessToken,
		&userProfile,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("Twitch API error: %w", err)
	}

	return oidc.Profile{
		ProviderUserID: userProfile.ID,
		Name:           userProfile.DisplayName,
		Email:          userProfile.Email,
		EmailVerified:  userProfile.Email != "",
		Picture:        userProfile.Avatar,
	}, nil
}
