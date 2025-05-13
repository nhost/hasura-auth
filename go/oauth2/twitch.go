package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

var (
	errNoUserDataFound      = errors.New("no user data found")
	errUnexpectedStatusCode = errors.New("unexpected status code")
)

type Twitch struct {
	*oauth2.Config
	ClientID string // This is required for Twitch Helix API requests.
}

func NewTwitchProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *Twitch {
	return &Twitch{
		ClientID: clientID,
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

type twitchUser struct {
	ID              string `json:"id"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ProfileImageURL string `json:"profile_image_url"`
}

type twitchUserResponse struct {
	Data []twitchUser `json:"data"`
}

func (t *Twitch) GetProfile(
	ctx context.Context,
	accessToken string,
	_ *string,
	_ map[string]any,
) (oidc.Profile, error) {
	var response twitchUserResponse

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.twitch.tv/helix/users",
		nil,
	)
	if err != nil {
		return oidc.Profile{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Twitch Helix requires the ClientID to be set.
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", t.ClientID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return oidc.Profile{}, fmt.Errorf("request error: %w", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return oidc.Profile{}, fmt.Errorf("%w: %d", errUnexpectedStatusCode, resp.StatusCode)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return oidc.Profile{}, fmt.Errorf("JSON unmarshal error: %w", err)
	}

	if len(response.Data) == 0 {
		return oidc.Profile{}, errNoUserDataFound
	}

	userProfile := response.Data[0]

	return oidc.Profile{
		ProviderUserID: userProfile.ID,
		Name:           userProfile.DisplayName,
		Email:          userProfile.Email,
		EmailVerified:  userProfile.Email != "",
		Picture:        userProfile.ProfileImageURL,
	}, nil
}
