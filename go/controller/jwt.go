package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTSecret struct {
	Key             string `json:"key"`
	Type            string `json:"type"`
	Issuer          string `json:"issuer"`
	ClaimsNamespace string `json:"claims_namespace"`
}

func decodeJWTSecret(jwtSecretb []byte) (JWTSecret, error) {
	var jwtSecret JWTSecret
	if err := json.Unmarshal(jwtSecretb, &jwtSecret); err != nil {
		return JWTSecret{}, fmt.Errorf("error unmarshalling jwt secret: %w", err)
	}

	if jwtSecret.Issuer == "" {
		jwtSecret.Issuer = "hasura-auth"
	}

	if jwtSecret.ClaimsNamespace == "" {
		jwtSecret.ClaimsNamespace = "https://hasura.io/jwt/claims"
	}

	return jwtSecret, nil
}

func NewJWTGetter(
	jwtSecretb []byte,
	accessTokenExpiresIn time.Duration,
) (
	func(userID uuid.UUID) (string, int64, error),
	error,
) {
	jwtSecret, err := decodeJWTSecret(jwtSecretb)
	if err != nil {
		return nil, err
	}

	mySigningKey := []byte(jwtSecret.Key)
	method := jwt.GetSigningMethod(jwtSecret.Type)

	return func(userID uuid.UUID) (string, int64, error) {
		now := time.Now()
		iat := now.Unix()
		exp := now.Add(accessTokenExpiresIn).Unix()

		// Create the Claims
		claims := &jwt.MapClaims{
			"sub": userID.String(),
			"iss": jwtSecret.Issuer,
			"iat": iat,
			"exp": exp,
			jwtSecret.ClaimsNamespace: map[string]any{
				"x-hasura-allowed-roles": []string{
					"admin",
					"user",
					"project_manager",
					"anonymous",
				},
				"x-hasura-default-role":     "user",
				"x-hasura-user-id":          userID.String(),
				"x-hasura-user-isAnonymous": "false",
			},
		}
		token := jwt.NewWithClaims(method, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			return "", 0, fmt.Errorf("error signing token: %w", err)
		}

		return ss, exp, nil
	}, nil
}

func JWTValidateFn(
	jwtSecretb []byte,
) (
	func(jwt string) (*jwt.Token, error),
	error,
) {
	jwtSecret, err := decodeJWTSecret(jwtSecretb)
	if err != nil {
		return nil, err
	}

	mySigningKey := []byte(jwtSecret.Key)

	return func(accessToken string) (*jwt.Token, error) {
		jwtToken, err := jwt.Parse(
			accessToken,
			func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			},
			jwt.WithValidMethods([]string{jwtSecret.Type}),
			jwt.WithIssuer(jwtSecret.Issuer),
			jwt.WithIssuedAt(),
			jwt.WithExpirationRequired(),
		)
		if err != nil {
			return nil, fmt.Errorf("error parsing token: %w", err)
		}
		return jwtToken, nil
	}, nil
}
