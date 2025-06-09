package cmd

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/provider"
	"github.com/urfave/cli/v2"
)

//nolint:cyclop
func getDefaultScopes(providerName api.SigninProvider) []string {
	switch providerName {
	case api.SigninProviderGoogle:
		return provider.DefaultGoogleScopes
	case api.SigninProviderDiscord:
		return provider.DefaultDiscordScopes
	case api.SigninProviderGithub:
		return provider.DefaultGithubScopes
	case api.SigninProviderApple:
		return provider.DefaultAppleScopes
	case api.SigninProviderLinkedin:
		return provider.DefaultLinkedInScopes
	case api.SigninProviderSpotify:
		return provider.DefaultSpotifyScopes
	case api.SigninProviderTwitch:
		return provider.DefaultTwitchScopes
	case api.SigninProviderGitlab:
		return provider.DefaultGitlabScopes
	case api.SigninProviderBitbucket:
		return provider.DefaultBitbucketScopes
	case api.SigninProviderWorkos:
		return provider.DefaultWorkOSScopes
	case api.SigninProviderAzuread:
		return provider.DefaultAzureadScopes
	case api.SigninProviderFacebook:
		return provider.DefaultFacebookScopes
	case api.SigninProviderWindowslive:
		return provider.DefaultWindowsliveScopes
	case api.SigninProviderStrava:
		return provider.DefaultStravaScopes
	default:
		panic("Unknown OAuth2 provider: " + providerName)
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
) (provider.Map, error) {
	providers := make(provider.Map)

	if cCtx.Bool(flagGoogleEnabled) {
		providers["google"] = provider.NewGoogleProvider(
			cCtx.String(flagGoogleClientID),
			cCtx.String(flagGoogleClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderGoogle, cCtx.StringSlice(flagGoogleScope)),
		)
	}

	if cCtx.Bool(flagGithubEnabled) {
		providers["github"] = provider.NewGithubProvider(
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
		clientSecret, err := provider.GenerateClientSecret(
			cCtx.String(flagAppleTeamID),
			cCtx.String(flagAppleKeyID),
			cCtx.String(flagAppleClientID),
			cCtx.String(flagApplePrivateKey),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to generate Apple client secret: %w", err)
		}

		providers["apple"], err = provider.NewAppleProvider(
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
		providers["linkedin"] = provider.NewLinkedInProvider(
			cCtx.String(flagLinkedInClientID),
			cCtx.String(flagLinkedInClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderLinkedin, cCtx.StringSlice(flagLinkedInScope)),
		)
	}

	if cCtx.Bool(flagDiscordEnabled) {
		providers["discord"] = provider.NewDiscordProvider(
			cCtx.String(flagDiscordClientID),
			cCtx.String(flagDiscordClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderDiscord, cCtx.StringSlice(flagDiscordScope)),
		)
	}

	if cCtx.Bool(flagSpotifyEnabled) {
		providers["spotify"] = provider.NewSpotifyProvider(
			cCtx.String(flagSpotifyClientID),
			cCtx.String(flagSpotifyClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderSpotify, cCtx.StringSlice(flagSpotifyScope)),
		)
	}

	if cCtx.Bool(flagTwitchEnabled) {
		providers["twitch"] = provider.NewTwitchProvider(
			cCtx.String(flagTwitchClientID),
			cCtx.String(flagTwitchClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderTwitch, cCtx.StringSlice(flagTwitchScope)),
		)
	}

	if cCtx.Bool(flagGitlabEnabled) {
		providers["gitlab"] = provider.NewGitlabProvider(
			cCtx.String(flagGitlabClientID),
			cCtx.String(flagGitlabClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderGitlab, cCtx.StringSlice(flagGitlabScope)),
		)
	}

	if cCtx.Bool(flagBitbucketEnabled) {
		providers["bitbucket"] = provider.NewBitbucketProvider(
			cCtx.String(flagBitbucketClientID),
			cCtx.String(flagBitbucketClientSecret),
			cCtx.String(flagServerURL),
			getScopes(api.SigninProviderBitbucket, cCtx.StringSlice(flagBitbucketScope)),
		)
	}

	if cCtx.Bool(flagWorkosEnabled) {
		providers["workos"] = provider.NewWorkosProvider(
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
		providers["azuread"] = provider.NewAzureadProvider(
			cCtx.String(flagAzureadClientID),
			cCtx.String(flagAzureadClientSecret),
			cCtx.String(flagServerURL),
			cCtx.String(flagAzureadTenant),
			getScopes(api.SigninProviderAzuread, cCtx.StringSlice(flagAzureadScope)),
		)
	}

	if cCtx.Bool(flagFacebookEnabled) {
		providers["facebook"] = provider.NewFacebookProvider(
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
			getScopes(api.SigninProviderWindowslive, cCtx.StringSlice(flagWindowsliveScope)),
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

	if cCtx.Bool(flagTwitterEnabled) {
		providers["twitter"] = provider.NewTwitterProvider(
			cCtx.String(flagTwitterConsumerKey),
			cCtx.String(flagTwitterConsumerSecret),
			cCtx.String(flagServerURL),
		)
	}

	return providers, nil
}
