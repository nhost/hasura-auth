package controller_test

import (
	"crypto"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/controller"
)

//nolint:lll,gochecknoglobals
var (
	jwtSecret                    = []byte(`{"type":"HS256", "key":"5152fa850c02dc222631cca898ed1485821a70912a6e3649c49076912daa3b62182ba013315915d64f40cddfbb8b58eb5bd11ba225336a6af45bbae07ca873f3"}`)
	jwtSecretWithIssuer          = []byte(`{"type":"HS256", "key":"5152fa850c02dc222631cca898ed1485821a70912a6e3649c49076912daa3b62182ba013315915d64f40cddfbb8b58eb5bd11ba225336a6af45bbae07ca873f3","issuer":"some-issuer"}`)
	jwtSecretWithClaimsNamespace = []byte(`{"type":"HS256", "key":"5152fa850c02dc222631cca898ed1485821a70912a6e3649c49076912daa3b62182ba013315915d64f40cddfbb8b58eb5bd11ba225336a6af45bbae07ca873f3","claims_namespace":"some/namespace"}`)
)

func TestGetJWTFunc(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("585e21fc-3664-4d03-8539-69945342a4f4")

	cases := []struct {
		name          string
		key           []byte
		userID        uuid.UUID
		expiresIn     time.Duration
		expectedToken *jwt.Token
	}{
		{
			name:      "with valid key",
			key:       jwtSecret,
			userID:    userID,
			expiresIn: time.Hour,
			expectedToken: &jwt.Token{
				Raw:    "ignored",
				Method: &jwt.SigningMethodHMAC{Name: "HS256", Hash: crypto.SHA256},
				Header: map[string]interface{}{"alg": string("HS256"), "typ": string("JWT")},
				Claims: jwt.MapClaims{
					"exp": float64(1.708103735e+09),
					"https://hasura.io/jwt/claims": map[string]interface{}{
						"x-hasura-allowed-roles": []interface{}{
							string("admin"),
							string("user"),
							string("project_manager"),
							string("anonymous"),
						},
						"x-hasura-default-role":     string("user"),
						"x-hasura-user-id":          string("585e21fc-3664-4d03-8539-69945342a4f4"),
						"x-hasura-user-isAnonymous": string("false"),
					},
					"iat": float64(1.708100135e+09),
					"iss": string("hasura-auth"),
					"sub": string("585e21fc-3664-4d03-8539-69945342a4f4"),
				},
				Signature: []uint8{},
				Valid:     true,
			},
		},

		{
			name:      "with valid key with issuer",
			key:       jwtSecretWithIssuer,
			userID:    userID,
			expiresIn: time.Hour,
			expectedToken: &jwt.Token{
				Raw:    "ignored",
				Method: &jwt.SigningMethodHMAC{Name: "HS256", Hash: crypto.SHA256},
				Header: map[string]interface{}{"alg": string("HS256"), "typ": string("JWT")},
				Claims: jwt.MapClaims{
					"exp": float64(1.708103735e+09),
					"https://hasura.io/jwt/claims": map[string]interface{}{
						"x-hasura-allowed-roles": []interface{}{
							string("admin"),
							string("user"),
							string("project_manager"),
							string("anonymous"),
						},
						"x-hasura-default-role":     string("user"),
						"x-hasura-user-id":          string("585e21fc-3664-4d03-8539-69945342a4f4"),
						"x-hasura-user-isAnonymous": string("false"),
					},
					"iat": float64(1.708100135e+09),
					"iss": string("some-issuer"),
					"sub": string("585e21fc-3664-4d03-8539-69945342a4f4"),
				},
				Signature: []uint8{},
				Valid:     true,
			},
		},

		{
			name:      "with valid key with claims namespace",
			key:       jwtSecretWithClaimsNamespace,
			userID:    userID,
			expiresIn: time.Hour,
			expectedToken: &jwt.Token{
				Raw:    "ignored",
				Method: &jwt.SigningMethodHMAC{Name: "HS256", Hash: crypto.SHA256},
				Header: map[string]interface{}{"alg": string("HS256"), "typ": string("JWT")},
				Claims: jwt.MapClaims{
					"exp": float64(1.708103735e+09),
					"some/namespace": map[string]interface{}{
						"x-hasura-allowed-roles": []interface{}{
							string("admin"),
							string("user"),
							string("project_manager"),
							string("anonymous"),
						},
						"x-hasura-default-role":     string("user"),
						"x-hasura-user-id":          string("585e21fc-3664-4d03-8539-69945342a4f4"),
						"x-hasura-user-isAnonymous": string("false"),
					},
					"iat": float64(1.708100135e+09),
					"iss": string("hasura-auth"),
					"sub": string("585e21fc-3664-4d03-8539-69945342a4f4"),
				},
				Signature: []uint8{},
				Valid:     true,
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc := tc

			getterFn, err := controller.NewJWTGetter(tc.key, tc.expiresIn)
			if err != nil {
				t.Fatalf("GetJWTFunc() err = %v; want nil", err)
			}

			accessToken, _, err := getterFn(tc.userID)
			if err != nil {
				t.Fatalf("fn() err = %v; want nil", err)
			}
			t.Logf("token = %v", accessToken)

			validatorFn, err := controller.JWTValidateFn(tc.key)
			if err != nil {
				t.Fatalf("JWTValidatorFn() err = %v; want nil", err)
			}

			decodedToken, err := validatorFn(accessToken)
			if err != nil {
				t.Fatalf("fn() err = %v; want nil", err)
			}

			cmpopts := []cmp.Option{
				cmpopts.IgnoreFields(jwt.Token{}, "Raw", "Signature"), //nolint:exhaustruct
				cmpopts.IgnoreMapEntries(func(key string, value interface{}) bool {
					return key == "iat" || key == "exp"
				}),
			}
			if diff := cmp.Diff(decodedToken, tc.expectedToken, cmpopts...); diff != "" {
				t.Errorf("decodedToken mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
