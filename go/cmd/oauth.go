package cmd

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/urfave/cli/v2"
)

func getDefaultScopes(provider string) []string {
	switch provider {
	case "google":
		return oauth2.DefaultGoogleScopes
	case "github":
		return oauth2.DefaultGithubScopes
	case "apple":
		return oauth2.DefaultAppleScopes
	case "linkedin":
		return oauth2.DefaultLinkedInScopes
	default:
		return []string{}
	}
}

func getScopes(provider string, scopes []string) []string {
	// clean the scopes in case of empty string
	var cleanedScopes []string
	for _, scope := range scopes {
		if scope != "" {
			cleanedScopes = append(cleanedScopes, scope)
		}
	}

	if len(cleanedScopes) > 0 {
		return cleanedScopes
	}
	return getDefaultScopes(provider)
}

func getOauth2Providers(cCtx *cli.Context) (*oauth2.Providers, error) {
	providers := make(map[string]oauth2.Provider)

	if cCtx.Bool(flagGoogleEnabled) {
		providers["google"] = oauth2.NewGoogleProvider(
			cCtx.String(flagGoogleClientID),
			cCtx.String(flagGoogleClientSecret),
			cCtx.String(flagServerURL),
			getScopes("google", cCtx.StringSlice(flagGoogleScope)),
		)
	}

	if cCtx.Bool(flagGithubEnabled) {
		providers["github"] = oauth2.NewGithubProvider(
			cCtx.String(flagGithubClientID),
			cCtx.String(flagGithubClientSecret),
			cCtx.String(flagServerURL),
			cCtx.String(flagGithubAuthorizationURL),
			cCtx.String(flagGithubTokenURL),
			cCtx.String(flagGithubUserProfileURL),
			getScopes("github", cCtx.StringSlice(flagGithubScope)),
		)
	}

	if cCtx.Bool(flagAppleEnabled) {
		clientSecret, err := oauth2.GenerateClientSecret(
			cCtx.String(flagAppleTeamID),
			cCtx.String(flagAppleKeyID),
			cCtx.String(flagAppleClientID),
			cCtx.String(flagApplePrivateKey),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to generate Apple client secret: %w", err)
		}

		providers["apple"], err = oauth2.NewAppleProvider(
			cCtx.Context,
			cCtx.String(flagAppleClientID),
			clientSecret,
			cCtx.String(flagServerURL),
			getScopes("apple", cCtx.StringSlice(flagAppleScope)),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create Apple provider: %w", err)
		}
	}

	if cCtx.Bool(flagLinkedInEnabled) {
		providers["linkedin"] = oauth2.NewLinkedInProvider(
			cCtx.String(flagLinkedInClientID),
			cCtx.String(flagLinkedInClientSecret),
			cCtx.String(flagServerURL),
			getScopes("linkedin", cCtx.StringSlice(flagLinkedInScope)),
		)
	}

	return oauth2.NewProviders(providers), nil
}
