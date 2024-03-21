package controller_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/controller/mock"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/nhost/hasura-auth/go/testhelpers"
	"go.uber.org/mock/gomock"
)

func webAuthnTouchID(
	t *testing.T,
) (*protocol.CredentialCreationResponse, controller.WebauthnChallenge) {
	t.Helper()

	//nolint:lll
	rawCredResp := []byte(`{
        "id": "LychOomEPgZu4XNwiDvzlP5hd1U",
        "rawId": "LychOomEPgZu4XNwiDvzlP5hd1U",
        "response": {
            "attestationObject": "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViY0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U",
            "clientDataJSON": "eyJjaGFsbGVuZ2UiOiJ6em56dGp2RlZVTTBFMnA4WlY2c2hYRWN3MmY0dGJ6NVJyZlpXazRWUFhJIiwib3JpZ2luIjoiaHR0cHM6Ly9yZWFjdC1hcG9sbG8uZXhhbXBsZS5uaG9zdC5pbyIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ",
            "transports": [
                "internal"
            ],
            "publicKeyAlgorithm": -7,
            "publicKey": "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV-G1gqCVxBrzZZ3dwmjPZjUlMqWGIvv3xsYIbanJZH-jUJQRuCdSrkbsVjo7Om1xJBBmrrJXddW7mIzQxZEfZQ",
            "authenticatorData": "0RE6Bmg2J-FxNrC8136ZQSeTWKWtdni_Lpfv5XR4bDtdAAAAAPv8MAcVTk7MjAtuAgVX170AFC8nITqJhD4GbuFzcIg785T-YXdVpQECAyYgASFYIFfhtYKglcQa82Wd3cJoz2Y1JTKlhiL798bGCG2pyWR_Ilggo1CUEbgnUq5G7FY6OzptcSQQZq6yV3XVu5iM0MWRH2U"
        },
        "type": "public-key",
        "clientExtensionResults": {},
        "authenticatorAttachment": "platform"
    }`)

	var resp *protocol.CredentialCreationResponse
	if err := json.Unmarshal(rawCredResp, &resp); err != nil {
		t.Fatal(err)
	}

	rawChallenge := []byte(`{
          "Session": {
            "challenge": "zznztjvFVUM0E2p8ZV6shXEcw2f4tbz5RrfZWk4VPXI",
            "user_id": "Y2Y5MWQxYmMtODc1ZS00OWJjLTg5N2YtZmJjY2YzMmVkZTEx",
            "expires": "0001-01-01T00:00:00Z",
            "userVerification": "preferred"
          },
          "User": {
            "ID": "cf91d1bc-875e-49bc-897f-fbccf32ede11",
            "Name": "Jane Doe",
            "Email": "jane@acme.com"
          },
          "Options": {
            "allowedRoles": [
              "user",
              "me"
            ],
            "defaultRole": "user",
            "displayName": "Jane Doe",
            "locale": "en",
            "redirectTo": "http://localhost:3000"
          }
        }`)

	var challenge controller.WebauthnChallenge
	if err := json.Unmarshal(rawChallenge, &challenge); err != nil {
		t.Fatal(err)
	}

	return resp, challenge
}

//nolint:dupl
func TestPostSignupWebauthnVerify(t *testing.T) { //nolint:maintidx
	t.Parallel()

	userID := uuid.MustParse("cf91d1bc-875e-49bc-897f-fbccf32ede11")

	touchIDRequest, touchIDWebauthnChallenge := webAuthnTouchID(t)

	cases := []testRequest[api.PostSignupWebauthnVerifyRequestObject, api.PostSignupWebauthnVerifyResponseObject]{
		{
			name:   "touchID - no email verify",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().InsertUserWithSecurityKeyAndRefreshToken(
					gomock.Any(),
					cmpDBParams(sql.InsertUserWithSecurityKeyAndRefreshTokenParams{
						ID:                    userID,
						Disabled:              false,
						DisplayName:           "Jane Doe",
						AvatarUrl:             "",
						Email:                 sql.Text("jane@acme.com"),
						Ticket:                pgtype.Text{}, //nolint:exhaustruct
						TicketExpiresAt:       sql.TimestampTz(time.Now()),
						EmailVerified:         false,
						Locale:                "en",
						DefaultRole:           "user",
						Metadata:              []byte("null"),
						Roles:                 []string{"user", "me"},
						RefreshTokenHash:      pgtype.Text{}, //nolint:exhaustruct
						RefreshTokenExpiresAt: sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
						CredentialID:          "LychOomEPgZu4XNwiDvzlP5hd1U",
						CredentialPublicKey: []uint8{
							0xa5, 0x01, 0x02, 0x03, 0x26, 0x20, 0x01, 0x21, 0x58, 0x20, 0x57, 0xe1, 0xb5, 0x82, 0xa0, 0x95, 0xc4, 0x1a, 0xf3, 0x65, 0x9d, 0xdd, 0xc2, 0x68, 0xcf, 0x66, 0x35, 0x25, 0x32, 0xa5, 0x86, 0x22, 0xfb, 0xf7, 0xc6, 0xc6, 0x08, 0x6d, 0xa9, 0xc9, 0x64, 0x7f, 0x22, 0x58, 0x20, 0xa3, 0x50, 0x94, 0x11, 0xb8, 0x27, 0x52, 0xae, 0x46, 0xec, 0x56, 0x3a, 0x3b, 0x3a, 0x6d, 0x71, 0x24, 0x10, 0x66, 0xae, 0xb2, 0x57, 0x75, 0xd5, 0xbb, 0x98, 0x8c, 0xd0, 0xc5, 0x91, 0x1f, 0x65, //nolint:lll
						},
						Nickname: pgtype.Text{}, //nolint:exhaustruct
					}),
				).Return(userID, nil)

				return mock
			},
			emailer: func(ctrl *gomock.Controller) *mock.MockEmailer {
				mock := mock.NewMockEmailer(ctrl)
				return mock
			},
			hibp: func(ctrl *gomock.Controller) *mock.MockHIBPClient {
				mock := mock.NewMockHIBPClient(ctrl)
				return mock
			},
			customClaimer: nil,
			request: api.PostSignupWebauthnVerifyRequestObject{
				Body: &api.SignUpWebauthnVerifyRequest{
					Credential:           touchIDRequest,
					Options:              nil,
					AdditionalProperties: nil,
				},
			},
			expectedResponse: api.PostSignupWebauthnVerify200JSONResponse{
				Session: &api.Session{
					AccessToken:          "xxxxx",
					AccessTokenExpiresIn: time.Now().Add(15 * time.Minute).Unix(),
					RefreshToken:         "ff0499a1-7935-4052-baea-6c3a573b1b6a",
					User: &api.User{
						AvatarUrl:           "",
						CreatedAt:           time.Now(),
						DefaultRole:         "user",
						DisplayName:         "Jane Doe",
						Email:               "jane@acme.com",
						EmailVerified:       false,
						Id:                  "cf91d1bc-875e-49bc-897f-fbccf32ede11",
						IsAnonymous:         false,
						Locale:              "en",
						Metadata:            nil,
						PhoneNumber:         "",
						PhoneNumberVerified: false,
						Roles:               []string{"user", "me"},
					},
				},
			},
			expectedJWT: &jwt.Token{
				Raw:    "",
				Method: jwt.SigningMethodHS256,
				Header: map[string]any{"alg": string("HS256"), "typ": string("JWT")},
				Claims: jwt.MapClaims{
					"exp": float64(time.Now().Add(900 * time.Second).Unix()),
					"https://hasura.io/jwt/claims": map[string]any{
						"x-hasura-allowed-roles":    []any{"user", "me"},
						"x-hasura-default-role":     string("user"),
						"x-hasura-user-id":          string("cf91d1bc-875e-49bc-897f-fbccf32ede11"),
						"x-hasura-user-isAnonymous": string("false"),
					},
					"iat": float64(time.Now().Unix()),
					"iss": string("hasura-auth"),
					"sub": string("cf91d1bc-875e-49bc-897f-fbccf32ede11"),
				},
				Signature: []uint8{},
				Valid:     true,
			},
			jwtTokenFn: nil,
		},

		{
			name: "touchID - email verify",
			config: func() *controller.Config {
				c := getConfig()
				c.RequireEmailVerification = true
				return c
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().InsertUserWithSecurityKey(
					gomock.Any(),
					cmpDBParams(sql.InsertUserWithSecurityKeyParams{
						ID:              userID,
						Disabled:        false,
						DisplayName:     "Jane Doe",
						AvatarUrl:       "",
						Email:           sql.Text("jane@acme.com"),
						Ticket:          sql.Text("verifyEmail:xxxx"),
						TicketExpiresAt: sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
						EmailVerified:   false,
						Locale:          "en",
						DefaultRole:     "user",
						Metadata:        []byte("null"),
						Roles:           []string{"user", "me"},
						CredentialID:    "LychOomEPgZu4XNwiDvzlP5hd1U",
						CredentialPublicKey: []uint8{
							0xa5, 0x01, 0x02, 0x03, 0x26, 0x20, 0x01, 0x21, 0x58, 0x20, 0x57, 0xe1, 0xb5, 0x82, 0xa0, 0x95, 0xc4, 0x1a, 0xf3, 0x65, 0x9d, 0xdd, 0xc2, 0x68, 0xcf, 0x66, 0x35, 0x25, 0x32, 0xa5, 0x86, 0x22, 0xfb, 0xf7, 0xc6, 0xc6, 0x08, 0x6d, 0xa9, 0xc9, 0x64, 0x7f, 0x22, 0x58, 0x20, 0xa3, 0x50, 0x94, 0x11, 0xb8, 0x27, 0x52, 0xae, 0x46, 0xec, 0x56, 0x3a, 0x3b, 0x3a, 0x6d, 0x71, 0x24, 0x10, 0x66, 0xae, 0xb2, 0x57, 0x75, 0xd5, 0xbb, 0x98, 0x8c, 0xd0, 0xc5, 0x91, 0x1f, 0x65, //nolint:lll
						},
						Nickname: pgtype.Text{}, //nolint:exhaustruct
					}),
				).Return(userID, nil)

				return mock
			},
			emailer: func(ctrl *gomock.Controller) *mock.MockEmailer {
				mock := mock.NewMockEmailer(ctrl)

				mock.EXPECT().SendEmail(
					"jane@acme.com",
					"en",
					notifications.TemplateNameEmailVerify,
					testhelpers.GomockCmpOpts(
						notifications.TemplateData{
							Link:        "https://local.auth.nhost.run/verify?redirectTo=http%3A%2F%2Flocalhost%3A3000&ticket=verifyEmail%3Ac2ee89db-095c-4904-b796-f6a507ee1260&type=emailVerify", //nolint:lll
							DisplayName: "Jane Doe",
							Email:       "jane@acme.com",
							NewEmail:    "",
							Ticket:      "verifyEmail:c2ee89db-095c-4904-b796-f6a507ee1260",
							RedirectTo:  "http://localhost:3000",
							Locale:      "en",
							ServerURL:   "https://local.auth.nhost.run",
							ClientURL:   "http://localhost:3000",
						},
						testhelpers.FilterPathLast(
							[]string{".Ticket"}, cmp.Comparer(cmpTicket)),

						testhelpers.FilterPathLast(
							[]string{".Link"}, cmp.Comparer(cmpLink)),
					)).Return(nil)

				return mock
			},
			hibp: func(ctrl *gomock.Controller) *mock.MockHIBPClient {
				mock := mock.NewMockHIBPClient(ctrl)
				return mock
			},
			customClaimer: nil,
			request: api.PostSignupWebauthnVerifyRequestObject{
				Body: &api.SignUpWebauthnVerifyRequest{
					Credential:           touchIDRequest,
					Options:              nil,
					AdditionalProperties: nil,
				},
			},
			expectedResponse: api.PostSignupWebauthnVerify200JSONResponse{
				Session: nil,
			},
			expectedJWT: nil,
			jwtTokenFn:  nil,
		},

		{
			name: "webauthn disabled",
			config: func() *controller.Config {
				c := getConfig()
				c.WebauthnEnabled = false
				return c
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				return mock
			},
			emailer: func(ctrl *gomock.Controller) *mock.MockEmailer {
				mock := mock.NewMockEmailer(ctrl)

				return mock
			},
			hibp: func(ctrl *gomock.Controller) *mock.MockHIBPClient {
				mock := mock.NewMockHIBPClient(ctrl)
				return mock
			},
			customClaimer: nil,
			request: api.PostSignupWebauthnVerifyRequestObject{
				Body: &api.SignUpWebauthnVerifyRequest{
					Credential:           touchIDRequest,
					Options:              nil,
					AdditionalProperties: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-endpoint",
				Message: "This endpoint is disabled",
				Status:  409,
			},
			expectedJWT: nil,
			jwtTokenFn:  nil,
		},

		{
			name: "signed up users disabled",
			config: func() *controller.Config {
				c := getConfig()
				c.DisableNewUsers = true
				return c
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().InsertUserWithSecurityKey(
					gomock.Any(),
					cmpDBParams(sql.InsertUserWithSecurityKeyParams{
						ID:              userID,
						Disabled:        true,
						DisplayName:     "Jane Doe",
						AvatarUrl:       "",
						Email:           sql.Text("jane@acme.com"),
						Ticket:          sql.Text("verifyEmail:xxxx"),
						TicketExpiresAt: sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
						EmailVerified:   false,
						Locale:          "en",
						DefaultRole:     "user",
						Metadata:        []byte("null"),
						Roles:           []string{"user", "me"},
						CredentialID:    "LychOomEPgZu4XNwiDvzlP5hd1U",
						CredentialPublicKey: []uint8{
							0xa5, 0x01, 0x02, 0x03, 0x26, 0x20, 0x01, 0x21, 0x58, 0x20, 0x57, 0xe1, 0xb5, 0x82, 0xa0, 0x95, 0xc4, 0x1a, 0xf3, 0x65, 0x9d, 0xdd, 0xc2, 0x68, 0xcf, 0x66, 0x35, 0x25, 0x32, 0xa5, 0x86, 0x22, 0xfb, 0xf7, 0xc6, 0xc6, 0x08, 0x6d, 0xa9, 0xc9, 0x64, 0x7f, 0x22, 0x58, 0x20, 0xa3, 0x50, 0x94, 0x11, 0xb8, 0x27, 0x52, 0xae, 0x46, 0xec, 0x56, 0x3a, 0x3b, 0x3a, 0x6d, 0x71, 0x24, 0x10, 0x66, 0xae, 0xb2, 0x57, 0x75, 0xd5, 0xbb, 0x98, 0x8c, 0xd0, 0xc5, 0x91, 0x1f, 0x65, //nolint:lll
						},
						Nickname: pgtype.Text{}, //nolint:exhaustruct
					}),
				).Return(userID, nil)

				return mock
			},
			emailer: func(ctrl *gomock.Controller) *mock.MockEmailer {
				mock := mock.NewMockEmailer(ctrl)

				return mock
			},
			hibp: func(ctrl *gomock.Controller) *mock.MockHIBPClient {
				mock := mock.NewMockHIBPClient(ctrl)
				return mock
			},
			customClaimer: nil,
			request: api.PostSignupWebauthnVerifyRequestObject{
				Body: &api.SignUpWebauthnVerifyRequest{
					Credential:           touchIDRequest,
					Options:              nil,
					AdditionalProperties: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-user",
				Message: "User is disabled",
				Status:  401,
			},
			expectedJWT: nil,
			jwtTokenFn:  nil,
		},

		{
			name: "disable sign ups",
			config: func() *controller.Config {
				c := getConfig()
				c.DisableSignup = true
				return c
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				return mock
			},
			emailer: func(ctrl *gomock.Controller) *mock.MockEmailer {
				mock := mock.NewMockEmailer(ctrl)

				return mock
			},
			hibp: func(ctrl *gomock.Controller) *mock.MockHIBPClient {
				mock := mock.NewMockHIBPClient(ctrl)
				return mock
			},
			customClaimer: nil,
			request: api.PostSignupWebauthnVerifyRequestObject{
				Body: &api.SignUpWebauthnVerifyRequest{
					Credential:           touchIDRequest,
					Options:              nil,
					AdditionalProperties: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "signup-disabled",
				Message: "Sign up is disabled.",
				Status:  403,
			},
			expectedJWT: nil,
			jwtTokenFn:  nil,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc := tc

			ctrl := gomock.NewController(t)

			c, jwtGetter := getController(t, ctrl, tc.config, tc.db, getControllerOpts{
				customClaimer: tc.customClaimer,
				emailer:       tc.emailer,
				hibp:          tc.hibp,
			})

			c.Webauthn.Storage["zznztjvFVUM0E2p8ZV6shXEcw2f4tbz5RrfZWk4VPXI"] = touchIDWebauthnChallenge

			resp := assertRequest(
				context.Background(),
				t,
				c.PostSignupWebauthnVerify,
				tc.request,
				tc.expectedResponse,
			)

			resp200, ok := resp.(api.PostSignupWebauthnVerify200JSONResponse)
			if ok {
				assertSession(t, jwtGetter, resp200.Session, tc.expectedJWT)

				if _, ok := c.Webauthn.Storage["zznztjvFVUM0E2p8ZV6shXEcw2f4tbz5RrfZWk4VPXI"]; ok {
					t.Errorf("challenge should've been removed")
				}
			}
		})
	}
}
