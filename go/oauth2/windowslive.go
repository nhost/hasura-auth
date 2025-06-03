package oauth2

import (
	"context"
	"fmt"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type WindowsLive struct {
	*oauth2.Config
}

func NewWindowsliveProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *WindowsLive {

	return &WindowsLive{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/windowslive/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{ //nolint:exhaustruct
				AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
				TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			},
		},
	}
}

type microsoftGraphProfile struct {
	ID                string `json:"id"`
	DisplayName       string `json:"displayName"`
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
}

func (w *WindowsLive) GetProfile(
	ctx context.Context,
	accessToken string,
	_ *string,
	_ map[string]any,
) (oidc.Profile, error) {
	var profile microsoftGraphProfile
	if err := fetchOAuthProfile(
		ctx,
		"https://graph.microsoft.com/v1.0/me",
		accessToken,
		&profile,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("microsoft graph api error: %w", err)
	}

	email := profile.Mail
	if email == "" {
		email = profile.UserPrincipalName
	}

	return oidc.Profile{
		ProviderUserID: profile.ID,
		Name:           profile.DisplayName,
		Email:          email,
		EmailVerified:  email != "",
		Picture:        "", // No avatar included
	}, nil
}
