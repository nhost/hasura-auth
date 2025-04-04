package cmd

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/urfave/cli/v2"
)

func getOauth2Providers(cCtx *cli.Context) (*oauth2.Providers, error) {
	providers := make(map[string]oauth2.Provider)

	if cCtx.Bool(flagGoogleEnabled) {
		providers["google"] = oauth2.NewGoogleProvider(
			cCtx.String(flagGoogleClientID),
			cCtx.String(flagGoogleClientSecret),
			cCtx.String(flagServerURL),
			cCtx.StringSlice(flagGoogleScope),
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
			cCtx.StringSlice(flagGithubScope),
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

		providers["apple"] = oauth2.NewAppleProvider(
			cCtx.String(flagAppleClientID),
			clientSecret,
			cCtx.String(flagServerURL),
			cCtx.StringSlice(flagAppleScope),
		)
	}

	if cCtx.Bool(flagLinkedInEnabled) {
		providers["linkedin"] = oauth2.NewLinkedInProvider(
			cCtx.String(flagLinkedInClientID),
			cCtx.String(flagLinkedInClientSecret),
			cCtx.String(flagServerURL),
			cCtx.StringSlice(flagLinkedInScope),
		)
	}

	return oauth2.NewProviders(providers), nil
}
