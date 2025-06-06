package cmd

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/urfave/cli/v2"
)

//nolint:cyclop
func getDefaultScopes(provider string) []string {
	switch provider {
	case "google":
		return oauth2.DefaultGoogleScopes
	case "discord":
		return oauth2.DefaultDiscordScopes
	case "github":
		return oauth2.DefaultGithubScopes
	case "apple":
		return oauth2.DefaultAppleScopes
	case "linkedin":
		return oauth2.DefaultLinkedInScopes
	case "spotify":
		return oauth2.DefaultSpotifyScopes
	case "twitch":
		return oauth2.DefaultTwitchScopes
	case "gitlab":
		return oauth2.DefaultGitlabScopes
	case "bitbucket":
		return oauth2.DefaultBitbucketScopes
	case "workos":
		return oauth2.DefaultWorkOSScopes
	case "azuread":
		return oauth2.DefaultAzureadScopes
	case "facebook":
		return oauth2.DefaultFacebookScopes
	default:
		panic("Unknown OAuth2 provider: " + provider)
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

//nolint:funlen,cyclop
func getOauth2Providers(
	cCtx *cli.Context,
) (*oauth2.Providers, error) {
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

	if cCtx.Bool(flagDiscordEnabled) {
		providers["discord"] = oauth2.NewDiscordProvider(
			cCtx.String(flagDiscordClientID),
			cCtx.String(flagDiscordClientSecret),
			cCtx.String(flagServerURL),
			getScopes("discord", cCtx.StringSlice(flagDiscordScope)),
		)
	}

	if cCtx.Bool(flagSpotifyEnabled) {
		providers["spotify"] = oauth2.NewSpotifyProvider(
			cCtx.String(flagSpotifyClientID),
			cCtx.String(flagSpotifyClientSecret),
			cCtx.String(flagServerURL),
			getScopes("spotify", cCtx.StringSlice(flagSpotifyScope)),
		)
	}

	if cCtx.Bool(flagTwitchEnabled) {
		providers["twitch"] = oauth2.NewTwitchProvider(
			cCtx.String(flagTwitchClientID),
			cCtx.String(flagTwitchClientSecret),
			cCtx.String(flagServerURL),
			getScopes("twitch", cCtx.StringSlice(flagTwitchScope)),
		)
	}

	if cCtx.Bool(flagGitlabEnabled) {
		providers["gitlab"] = oauth2.NewGitlabProvider(
			cCtx.String(flagGitlabClientID),
			cCtx.String(flagGitlabClientSecret),
			cCtx.String(flagServerURL),
			getScopes("gitlab", cCtx.StringSlice(flagGitlabScope)),
		)
	}

	if cCtx.Bool(flagBitbucketEnabled) {
		providers["bitbucket"] = oauth2.NewBitbucketProvider(
			cCtx.String(flagBitbucketClientID),
			cCtx.String(flagBitbucketClientSecret),
			cCtx.String(flagServerURL),
			getScopes("bitbucket", cCtx.StringSlice(flagBitbucketScope)),
		)
	}

	if cCtx.Bool(flagWorkosEnabled) {
		providers["workos"] = oauth2.NewWorkosProvider(
			cCtx.String(flagWorkosClientID),
			cCtx.String(flagWorkosClientSecret),
			cCtx.String(flagServerURL),
			getScopes("workos", cCtx.StringSlice(flagWorkosScope)),
			cCtx.String(flagWorkosDefaultOrganization),
			cCtx.String(flagWorkosDefaultConnection),
			cCtx.String(flagWorkosDefaultDomain),
		)
	}

	if cCtx.Bool(flagAzureadEnabled) {
		providers["azuread"] = oauth2.NewAzureadProvider(
			cCtx.String(flagAzureadClientID),
			cCtx.String(flagAzureadClientSecret),
			cCtx.String(flagServerURL),
			cCtx.String(flagAzureadTenant),
			getScopes("azuread", cCtx.StringSlice(flagAzureadScope)),
		)
	}

	if cCtx.Bool(flagFacebookEnabled) {
		providers["facebook"] = oauth2.NewFacebookProvider(
			cCtx.String(flagFacebookClientID),
			cCtx.String(flagFacebookClientSecret),
			cCtx.String(flagServerURL),
			getScopes("facebook", cCtx.StringSlice(flagFacebookScope)),
		)
	}

	return oauth2.NewProviders(providers), nil
}
