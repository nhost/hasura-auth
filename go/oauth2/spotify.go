package oauth2

import (
	"context"
	"fmt"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type Spotify struct {
	profileURL string
	*oauth2.Config
}

func NewSpotifyProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *Spotify {
	redirectURL := authServerURL + "/signin/provider/spotify/callback"

	return &Spotify{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.spotify.com/authorize",
				TokenURL: "https://accounts.spotify.com/api/token",
			},
		},
		profileURL: "https://api.spotify.com/v1/me",
	}
}

type spotifyImage struct {
	URL string `json:"url"`
}

type spotifyUserProfile struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	DisplayName string         `json:"display_name"`
	Images      []spotifyImage `json:"images"`
}

func (s *Spotify) GetProfile(
	ctx context.Context,
	accessToken string,
	_ *string,
	_ map[string]any,
) (oidc.Profile, error) {
	var userProfile spotifyUserProfile

	if err := fetchOAuthProfile(
		ctx,
		s.profileURL,
		accessToken,
		&userProfile,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("Spotify API error: %w", err)
	}

	var avatarURL string
	if len(userProfile.Images) > 0 {
		avatarURL = userProfile.Images[0].URL
	}

	return oidc.Profile{
		ProviderUserID: userProfile.ID,
		Name:           userProfile.DisplayName,
		Email:          userProfile.Email,
		EmailVerified:  userProfile.Email != "",
		Picture:        avatarURL,
	}, nil
}
