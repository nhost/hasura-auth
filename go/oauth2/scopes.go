package oauth2

//nolint:gochecknoglobals
var (
	// DefaultGoogleScopes defines the default scopes for Google OAuth2.
	DefaultGoogleScopes = []string{"openid", "email", "profile"}

	// DefaultGithubScopes defines the default scopes for GitHub OAuth2.
	DefaultGithubScopes = []string{"user:email"}

	// DefaultAppleScopes defines the default scopes for Apple OAuth2.
	DefaultAppleScopes = []string{"name", "email"}

	// DefaultLinkedInScopes defines the default scopes for LinkedIn OAuth2.
	DefaultLinkedInScopes = []string{"openid", "profile", "email"}

	// DefaultDiscordScopes defines the default scopes for Discord OAuth2.
	DefaultDiscordScopes = []string{"identify", "email"}

	// DefaultSpotifyScopes defines the default scopes for Spotify OAuth2.
	DefaultSpotifyScopes = []string{"user-read-email", "user-read-private"}

	// DefaultTwitchScopes defines the default scopes for Twitch OAuth2.
	DefaultTwitchScopes = []string{"user:read:email"}
)
