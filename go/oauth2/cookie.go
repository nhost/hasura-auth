package oauth2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/nhost/hasura-auth/go/api"
)

type ProviderSignInData struct {
	State   string            `json:"state"`
	Connect *string           `json:"connect"`
	Options api.SignUpOptions `json:"options"`
}

// Encode serializes the ProviderSignInData to a base64 URL-encoded string.
func (data *ProviderSignInData) Encode() (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal sign-in data: %w", err)
	}

	encoded := base64.RawURLEncoding.EncodeToString(b)
	return encoded, nil
}

// Decode deserializes a base64 URL-encoded string into ProviderSignInData.
func (data *ProviderSignInData) Decode(encoded string) error {
	b, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("failed to decode base64 string: %w", err)
	}

	if err := json.Unmarshal(b, data); err != nil {
		return fmt.Errorf("failed to unmarshal sign-in data: %w", err)
	}

	return nil
}
