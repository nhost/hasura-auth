package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nhost/hasura-auth/go/oidc"
	"github.com/nhost/hasura-auth/go/provider/oauth1"
)

type Twitter struct {
	*oauth1.Config
}

func NewTwitterProvider(
	consumerKey, consumerSecret, authServerURL string,
) *Provider {
	authorizeURL, err := url.Parse("https://api.twitter.com/oauth/authorize")
	if err != nil {
		panic("invalid Twitter authorize URL: " + err.Error())
	}

	twitter := &Twitter{
		Config: &oauth1.Config{
			ConsumerKey:     consumerKey,
			ConsumerSecret:  consumerSecret,
			CallbackURL:     authServerURL + "/signin/provider/twitter/callback",
			AccessURL:       "https://api.twitter.com/oauth/authenticate",
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    authorizeURL,
		},
	}

	return NewOauth1Provider(twitter)
}

type twitterUser struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	ProfileImageURL string `json:"profile_image_url_https"`
}

func (t *Twitter) GetProfile(
	ctx context.Context,
	_ string,
	_ *string,
	extras map[string]any,
) (oidc.Profile, error) {
	oauthToken, ok := extras["oauth_token"].(string)
	if !ok || oauthToken == "" {
		return oidc.Profile{}, fmt.Errorf("missing or invalid oauth_token in extras")
	}

	oauthVerifier, ok := extras["oauth_verifier"].(string)
	if !ok || oauthVerifier == "" {
		return oidc.Profile{}, fmt.Errorf("missing or invalid oauth_verifier in extras")
	}

	// Exchange oauth_token and oauth_verifier for access tokens
	accessTokenValue, accessTokenSecret, err := t.exchangeForAccessToken(oauthToken, oauthVerifier)
	if err != nil {
		return oidc.Profile{}, fmt.Errorf("failed to exchange for access token: %w", err)
	}

	// Fetch user profile using OAuth 1.0a
	var user twitterUser
	if err := t.fetchTwitterProfile(ctx, accessTokenValue, accessTokenSecret, &user); err != nil {
		return oidc.Profile{}, fmt.Errorf("Twitter API error: %w", err)
	}

	return oidc.Profile{
		ProviderUserID: user.ID,
		Email:          user.Email,
		EmailVerified:  user.Email != "",
		Name:           user.Name,
		Picture:        user.ProfileImageURL,
	}, nil
}

// fetchTwitterProfile makes an OAuth 1.0a authenticated request to Twitter API
func (t *Twitter) fetchTwitterProfile(
	ctx context.Context,
	oauthToken, oauthTokenSecret string,
	result any,
) error {
	// Twitter API endpoint for user profile (base URL without query params for signature)
	baseURL := "https://api.twitter.com/1.1/account/verify_credentials.json"

	// OAuth 1.0a parameters (including the query parameter for signature)
	params := map[string]string{
		"oauth_consumer_key":     t.ConsumerKey,
		"oauth_nonce":            oauth1.Nonce(),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_token":            oauthToken,
		"oauth_version":          "1.0",
		"include_email":          "true", // This must be included in signature
	}

	// Create signature using base URL without query params
	params["oauth_signature"] = oauth1.CreateSignature(
		http.MethodGet, baseURL, t.ConsumerSecret, params, oauthTokenSecret,
	)

	// Create authorization header (only include oauth_ parameters)
	var oauthParams []string
	for k, v := range params {
		if strings.HasPrefix(k, "oauth_") {
			oauthParams = append(oauthParams, url.QueryEscape(k)+"=\""+url.QueryEscape(v)+"\"")
		}
	}
	sort.Strings(oauthParams)
	authHeader := "OAuth " + strings.Join(oauthParams, ", ")

	// Use the full URL with query parameter
	reqURL := baseURL + "?include_email=true"

	// Make request
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	fmt.Printf("Response body: %s\n", string(body)) // Debugging line

	if err := json.NewDecoder(strings.NewReader(string(body))).Decode(result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

// exchangeForAccessToken exchanges the oauth_token and oauth_verifier for access tokens
func (t *Twitter) exchangeForAccessToken(oauthToken, oauthVerifier string) (string, string, error) {
	// OAuth 1.0a access token parameters
	params := map[string]string{
		"oauth_consumer_key":     t.ConsumerKey,
		"oauth_nonce":            oauth1.Nonce(),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_token":            oauthToken,
		"oauth_verifier":         oauthVerifier,
		"oauth_version":          "1.0",
	}

	// For access token exchange, we need the request token secret
	// This should be stored from the initial request token call
	// For now, we'll use empty string as it's not available in this context
	params["oauth_signature"] = oauth1.CreateSignature(
		"POST",
		"https://api.twitter.com/oauth/access_token",
		t.ConsumerSecret,
		params,
		"",
	)

	// Create authorization header
	authHeader := oauth1.AuthHeader(params)

	// Create form data for POST body
	formData := url.Values{}
	formData.Set("oauth_verifier", oauthVerifier)

	// Make request
	req, err := http.NewRequest(
		"POST",
		"https://api.twitter.com/oauth/access_token",
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf(
			"access token request failed (status %d): %s",
			resp.StatusCode,
			string(body),
		)
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %w", err)
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", "", fmt.Errorf("failed to parse response: %w", err)
	}

	accessToken := values.Get("oauth_token")
	accessTokenSecret := values.Get("oauth_token_secret")

	if accessToken == "" || accessTokenSecret == "" {
		return "", "", fmt.Errorf("access token or secret not found in response")
	}

	return accessToken, accessTokenSecret, nil
}
