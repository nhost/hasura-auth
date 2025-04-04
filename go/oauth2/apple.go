package oauth2

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nhost/hasura-auth/go/oidc"
	"golang.org/x/oauth2"
)

type Apple struct {
	*oauth2.Config
}

// NewAppleProvider creates a new Apple OAuth2 provider with JWT token generation support
func NewAppleProvider(
	clientID, clientSecret, authServerURL string,
	scopes []string,
) *Apple {
	return &Apple{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  authServerURL + "/signin/provider/apple/callback",
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://appleid.apple.com/auth/authorize",
				TokenURL: "https://appleid.apple.com/auth/token",
			},
		},
	}
}

// GenerateClientSecret creates a JWT token for Apple authentication
// This should be called instead of using a static client secret
func GenerateClientSecret(teamID, keyID, clientID, privateKeyBase64 string) (string, error) {
	// Replace escaped newlines with actual newlines
	privateKeyPEM := strings.ReplaceAll(privateKeyBase64, "\\n", "\n")

	// Check if the private key is base64 encoded
	if !strings.HasPrefix(privateKeyPEM, "-----BEGIN") {
		// Try to decode from base64
		decodedBytes, err := base64.StdEncoding.DecodeString(privateKeyPEM)
		if err != nil {
			return "", fmt.Errorf("failed to decode private key from base64: %w", err)
		}
		privateKeyPEM = string(decodedBytes)
	}

	// Parse the private key
	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// Create the token
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": teamID,                               // Team ID
		"iat": now.Unix(),                           // Issued at
		"exp": now.Add(time.Hour * 24 * 180).Unix(), // 180 days validity
		"aud": "https://appleid.apple.com",          // Apple's authorization server
		"sub": clientID,                             // Client ID
	})

	// Set the key ID in the header
	token.Header["kid"] = keyID

	// Sign the token
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return tokenString, nil
}

type appleUser struct {
	ID            string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func (a *Apple) GetProfile(ctx context.Context, accessToken string) (oidc.Profile, error) {
	var user appleUser
	if err := fetchOAuthProfile(
		ctx,
		"https://appleid.apple.com/auth/userinfo",
		accessToken,
		&user,
		fetchProfileTimeout,
	); err != nil {
		return oidc.Profile{}, fmt.Errorf("Apple API error: %w", err)
	}

	return oidc.Profile{
		ProviderUserID: user.ID,
		Email:          user.Email,
		EmailVerified:  user.EmailVerified,
		Name:           user.Name,
		Picture:        user.Picture,
	}, nil
}
