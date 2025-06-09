package oauth1

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Config struct {
	ConsumerKey     string   `json:"consumer_key"`
	ConsumerSecret  string   `json:"consumer_secret"`
	CallbackURL     string   `json:"callback_url"`
	AccessURL       string   `json:"access_url"`
	RequestTokenURL string   `json:"request_token_url"`
	AuthorizeURL    *url.URL `json:"authorize_url"`
}

// SignedRequestOptions contains options for creating a signed OAuth 1.0a request.
type SignedRequestOptions struct {
	Method      string
	URL         string
	Body        io.Reader
	Headers     map[string]string
	ExtraParams map[string]string
	OAuthToken  string
	TokenSecret string
	ContentType string
	Timeout     time.Duration
}

// SignedRequest creates and executes a signed OAuth 1.0a HTTP request.
func (c *Config) SignedRequest(
	ctx context.Context, opts SignedRequestOptions,
) (*http.Response, error) {
	// Set default values
	if opts.Method == "" {
		opts.Method = http.MethodGet
	}
	if opts.ContentType == "" {
		opts.ContentType = "application/x-www-form-urlencoded"
	}
	if opts.Timeout == 0 {
		opts.Timeout = 10 * time.Second //nolint:mnd
	}

	// Build OAuth 1.0a parameters
	params := map[string]string{
		"oauth_consumer_key":     c.ConsumerKey,
		"oauth_nonce":            Nonce(),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_version":          "1.0",
	}

	// Add OAuth token if provided
	if opts.OAuthToken != "" {
		params["oauth_token"] = opts.OAuthToken
	}

	// Add extra parameters
	maps.Copy(params, opts.ExtraParams)

	// Create signature
	params["oauth_signature"] = CreateSignature(
		opts.Method, opts.URL, c.ConsumerSecret, params, opts.TokenSecret,
	)

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, opts.Method, opts.URL, opts.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set Authorization header
	req.Header.Set("Authorization", AuthHeader(params))
	req.Header.Set("Content-Type", opts.ContentType)

	// Add custom headers
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Execute request
	client := &http.Client{Timeout: opts.Timeout} //nolint:exhaustruct
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return resp, nil
}

func (c *Config) RequestToken(ctx context.Context, state string) (string, error) {
	// Use the generalized SignedRequest function
	resp, err := c.SignedRequest(ctx, SignedRequestOptions{ //nolint:exhaustruct
		Method: http.MethodPost,
		URL:    c.RequestTokenURL,
		ExtraParams: map[string]string{
			"oauth_callback": c.CallbackURL + "?state=" + url.QueryEscape(state),
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf(
			"%w: request token request failed (status %d): %s",
			ErrHTTPRequestFailed, resp.StatusCode, string(body),
		)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	oauthToken := values.Get("oauth_token")
	callbackConfirmed := values.Get("oauth_callback_confirmed")

	if oauthToken == "" {
		return "", fmt.Errorf("%w: oauth_token not found in response", ErrInvalidResponse)
	}
	if callbackConfirmed != "true" {
		return "", fmt.Errorf("%w: oauth_callback_confirmed is not true", ErrInvalidResponse)
	}

	// Return authorization URL
	query := c.AuthorizeURL.Query()
	query.Set("oauth_token", oauthToken)
	c.AuthorizeURL.RawQuery = query.Encode()
	return c.AuthorizeURL.String(), nil
}

func (c *Config) AuthCodeURL(ctx context.Context, state string) (string, error) {
	authURL, err := c.RequestToken(ctx, state)
	if err != nil {
		return "", fmt.Errorf("failed to get auth code URL: %w", err)
	}
	return authURL, nil
}

// AccessToken exchanges an OAuth verifier for an access token.
func (c *Config) AccessToken(
	ctx context.Context, requestToken, verifier string,
) (string, string, error) {
	resp, err := c.SignedRequest(ctx, SignedRequestOptions{ //nolint:exhaustruct
		Method:     http.MethodPost,
		URL:        c.AccessURL,
		OAuthToken: requestToken,
		ExtraParams: map[string]string{
			"oauth_verifier": verifier,
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to exchange token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf(
			"%w: access token request failed (status %d): %s",
			ErrHTTPRequestFailed, resp.StatusCode, string(body),
		)
	}

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

	if accessToken == "" {
		return "", "", fmt.Errorf("%w: oauth_token not found in response", ErrInvalidResponse)
	}
	if accessTokenSecret == "" {
		return "", "", fmt.Errorf("%w: oauth_token_secret not found in response", ErrInvalidResponse)
	}

	return accessToken, accessTokenSecret, nil
}
