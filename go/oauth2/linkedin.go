package oauth2

import (
	"context"
	"fmt"

	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type LinkedIn struct {
	*oauth2.Config
}

func NewLinkedInProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *LinkedIn {
	return &LinkedIn{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/linkedin/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
				TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
			},
		},
	}
}

type linkedInUserProfile struct {
	ID        string `json:"id"`
	FirstName string `json:"localizedFirstName"`
	LastName  string `json:"localizedLastName"`
}

type linkedInUserEmail struct {
	Elements []struct {
		HandleTilde struct {
			EmailAddress string `json:"emailAddress"`
		} `json:"handle~"`
		Primary bool `json:"primary"`
	} `json:"elements"`
}

func (l *LinkedIn) GetProfile(ctx context.Context, accessToken string) (oidc.Profile, error) {
	// Fetch basic profile data
	var userProfile linkedInUserProfile
	if err := fetchOAuthProfile(
		ctx,
		"https://api.linkedin.com/v2/me",
		accessToken,
		&userProfile,
		fetchProfileTimeout,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("LinkedIn API error (profile): %w", err)
	}

	// Fetch email data
	var userEmail linkedInUserEmail
	if err := fetchOAuthProfile(
		ctx,
		"https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))",
		accessToken,
		&userEmail,
		fetchProfileTimeout,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("LinkedIn API error (email): %w", err)
	}

	// Extract email from the response
	email := ""
	emailVerified := false
	if len(userEmail.Elements) > 0 {
		email = userEmail.Elements[0].HandleTilde.EmailAddress
		emailVerified = email != ""
	}

	// Construct the full name
	name := userProfile.FirstName
	if userProfile.LastName != "" {
		if name != "" {
			name += " "
		}
		name += userProfile.LastName
	}

	return oidc.Profile{
		ProviderUserID: userProfile.ID,
		Email:          email,
		EmailVerified:  emailVerified,
		Name:           name,
		Picture:        "", // LinkedIn doesn't provide picture URL in the basic profile
	}, nil
}
