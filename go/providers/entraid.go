package providers

import (
	"context"
	"fmt"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type EntraID struct {
	*oauth2.Config

	ProfileURL string
}

func NewEnrtaIDProvider(
	clientID, clientSecret, authServerURL, tenant string,
	scopes []string,
) *Provider {
	baseURL := "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0"

	entraid := &AzureAD{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/entraid/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{ //nolint:exhaustruct
				AuthURL:  baseURL + "/authorize",
				TokenURL: baseURL + "/token",
			},
		},
		ProfileURL: formatAzureADURL(tenant, "/openid/userinfo"),
	}

	return NewOauth2Provider(entraid)
}

type entraidUser struct {
	OID    string `json:"oid"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	UPN    string `json:"upn"`
	Prefer string `json:"preferred_username"`
}

func (a *EntraID) GetProfile(
	ctx context.Context,
	accessToken string,
	_ *string,
	_ map[string]any,
) (oidc.Profile, error) {
	var userProfile entraidUser
	if err := fetchOAuthProfile(
		ctx,
		a.ProfileURL,
		accessToken,
		&userProfile,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("EntraID API error: %w", err)
	}

	email := userProfile.Email
	if email == "" {
		email = userProfile.Prefer
	}

	if email == "" {
		email = userProfile.UPN
	}

	return oidc.Profile{
		ProviderUserID: userProfile.OID,
		Email:          email,
		EmailVerified:  email != "",
		Name:           userProfile.Name,
		Picture:        "",
	}, nil
}
