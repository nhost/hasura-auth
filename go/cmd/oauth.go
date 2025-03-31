package cmd

import (
	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/urfave/cli/v2"
)

func getOauth2Providers(cCtx *cli.Context) *oauth2.Providers {
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

	return oauth2.NewProviders(providers)
}
