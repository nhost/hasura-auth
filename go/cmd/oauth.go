package cmd

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/urfave/cli/v2"
)

//nolint:cyclop
func getDefaultScopes(provider api.SigninProvider) []string {
	switch provider {
	case api.SigninProviderGoogle:
		return oauth2.DefaultGoogleScopes
	case api.SigninProviderDiscord:
		return oauth2.DefaultDiscordScopes
	case api.SigninProviderGithub:
		return oauth2.DefaultGithubScopes
	case api.SigninProviderApple:
		return oauth2.DefaultAppleScopes
	case api.SigninProviderLinkedin:
		return oauth2.DefaultLinkedInScopes
	case api.SigninProviderSpotify:
		return oauth2.DefaultSpotifyScopes
	case api.SigninProviderTwitch:
		return oauth2.DefaultTwitchScopes
	case api.SigninProviderGitlab:
		return oauth2.DefaultGitlabScopes
	case api.SigninProviderBitbucket:
		return oauth2.DefaultBitbucketScopes
	case api.SigninProviderWorkos:
		return oauth2.DefaultWorkOSScopes
	case api.SigninProviderAzuread:
		return oauth2.DefaultAzureadScopes
	case api.SigninProviderFacebook:
		return oauth2.DefaultFacebookScopes
	case api.SigninProviderWindowslive:
		return oauth2.DefaultWindowsliveScopes
	case api.SigninProviderStrava:
		return oauth2.DefaultStravaScopes
	default:
		panic("Unknown OAuth2 provider: " + provider)
	}
}

func getScopes(provider api.SigninProvider, scopes []string) []string {
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
			getScopes(api.SigninProviderGoogle, cCtx.StringSlice(flagGoogleScope)),
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
			getScopes(api.SigninProviderGithub, cCtx.StringSlice(flagGithubScope)),
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
			getScopes(api.SigninProviderApple, cCtx.StringSlice(flagAppleScope)),
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
			getScopes(api.SigninProviderLinkedin, cCtx.StringSlice(flagLinkedInScope)),
		)
	}

	if cCtx.Bool(flagDiscordEnabled) {
		providers["discord"] = oauth2.NewDiscordProvider(
			cCtx.String(flagDiscordClientID),
			cCtx.String(flagDiscordClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderDiscord, cCtx.StringSlice(flagDiscordScope)),
		)
	}

	if cCtx.Bool(flagSpotifyEnabled) {
		providers["spotify"] = oauth2.NewSpotifyProvider(
			cCtx.String(flagSpotifyClientID),
			cCtx.String(flagSpotifyClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderSpotify, cCtx.StringSlice(flagSpotifyScope)),
		)
	}

	if cCtx.Bool(flagTwitchEnabled) {
		providers["twitch"] = oauth2.NewTwitchProvider(
			cCtx.String(flagTwitchClientID),
			cCtx.String(flagTwitchClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderTwitch, cCtx.StringSlice(flagTwitchScope)),
		)
	}

	if cCtx.Bool(flagGitlabEnabled) {
		providers["gitlab"] = oauth2.NewGitlabProvider(
			cCtx.String(flagGitlabClientID),
			cCtx.String(flagGitlabClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderGitlab, cCtx.StringSlice(flagGitlabScope)),
		)
	}

	if cCtx.Bool(flagBitbucketEnabled) {
		providers["bitbucket"] = oauth2.NewBitbucketProvider(
			cCtx.String(flagBitbucketClientID),
			cCtx.String(flagBitbucketClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderBitbucket, cCtx.StringSlice(flagBitbucketScope)),
		)
	}

	if cCtx.Bool(flagWorkosEnabled) {
		providers["workos"] = oauth2.NewWorkosProvider(
			cCtx.String(flagWorkosClientID),
			cCtx.String(flagWorkosClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderWorkos, cCtx.StringSlice(flagWorkosScope)),
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
			getScopes(api.SigninProviderAzuread, cCtx.StringSlice(flagAzureadScope)),
		)
	}

	if cCtx.Bool(flagFacebookEnabled) {
		providers["facebook"] = oauth2.NewFacebookProvider(
			cCtx.String(flagFacebookClientID),
			cCtx.String(flagFacebookClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderFacebook, cCtx.StringSlice(flagFacebookScope)),
		)
	}

	if cCtx.Bool(flagWindowsliveEnabled) {
		providers["windowslive"] = oauth2.NewWindowsliveProvider(
			cCtx.String(flagWindowsliveClientID),
			cCtx.String(flagWindowsliveClientSecret),
			cCtx.String(flagServerURL),
			getScopes("windowslive", cCtx.StringSlice(flagWindowsliveScope)),
		)
	}
	if cCtx.Bool(flagStravaEnabled) {
		providers["strava"] = oauth2.NewStravaProvider(
			cCtx.String(flagStravaClientID),
			cCtx.String(flagStravaClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderStrava, cCtx.StringSlice(flagStravaScope)),
		)
	}

	return oauth2.NewProviders(providers), nil
}
