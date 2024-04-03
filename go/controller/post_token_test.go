package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/controller/mock"
	"github.com/nhost/hasura-auth/go/sql"
	"go.uber.org/mock/gomock"
)

func TestPostToken(t *testing.T) {
	t.Parallel()

	userID := uuid.MustParse("db477732-48fa-4289-b694-2886a646b6eb")
	token := uuid.MustParse("1fb17604-86c7-444e-b337-09a644465f2d")
	hashedToken := `\x9698157153010b858587119503cbeef0cf288f11775e51cdb6bfd65e930d9310`

	cases := []testRequest[api.PostTokenRequestObject, api.PostTokenResponseObject]{
		{
			name:   "simple",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedToken),
						Type:             "regular",
					},
				).Return(getSigninUser(userID), nil)

				mock.EXPECT().RefreshTokenAndGetUserRoles(
					gomock.Any(),
					cmpDBParams(sql.RefreshTokenAndGetUserRolesParams{
						RefreshTokenHash: sql.Text(hashedToken),
						ExpiresAt: sql.TimestampTz(
							time.Now().Add(time.Duration(2592000) * time.Second),
						),
					}),
				).Return([]string{"user", "me"}, nil)

				return mock
			},
			request: api.PostTokenRequestObject{
				Body: &api.RefreshTokenRequest{
					RefreshToken: token.String(),
				},
			},
			expectedResponse: api.PostToken200JSONResponse{
				Session: &api.Session{
					AccessToken:          "",
					AccessTokenExpiresIn: time.Now().Add(900 * time.Second).Unix(),
					RefreshToken:         "1fb17604-86c7-444e-b337-09a644465f2d",
					User: &api.User{
						AvatarUrl:           "",
						CreatedAt:           time.Now(),
						DefaultRole:         "user",
						DisplayName:         "Jane Doe",
						Email:               "jane@acme.com",
						EmailVerified:       true,
						Id:                  "db477732-48fa-4289-b694-2886a646b6eb",
						IsAnonymous:         false,
						Locale:              "en",
						Metadata:            map[string]any{},
						PhoneNumber:         "",
						PhoneNumberVerified: false,
						Roles:               []string{"user", "me"},
					},
				},
			},
			expectedJWT: &jwt.Token{
				Raw:    "",
				Method: jwt.SigningMethodHS256,
				Header: map[string]any{
					"alg": "HS256",
					"typ": "JWT",
				},
				Claims: jwt.MapClaims{
					"exp": float64(time.Now().Add(900 * time.Second).Unix()),
					"https://hasura.io/jwt/claims": map[string]any{
						"x-hasura-allowed-roles":     []any{"user", "me"},
						"x-hasura-default-role":      "user",
						"x-hasura-user-id":           "db477732-48fa-4289-b694-2886a646b6eb",
						"x-hasura-user-is-anonymous": "false",
					},
					"iat": float64(time.Now().Unix()),
					"iss": "hasura-auth",
					"sub": "db477732-48fa-4289-b694-2886a646b6eb",
				},
				Signature: []byte{},
				Valid:     true,
			},
			customClaimer: nil,
			emailer:       nil,
			hibp:          nil,
			jwtTokenFn:    nil,
		},
		{
			name: "invalid refresh token",
			config: func() *controller.Config {
				cfg := getConfig()
				cfg.AllowedEmails = []string{"asd@asd.com"}
				return cfg
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedToken),
						Type:             "regular",
					},
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				return mock
			},
			customClaimer: nil,
			request: api.PostTokenRequestObject{
				Body: &api.RefreshTokenRequest{
					RefreshToken: token.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-refresh-token",
				Message: "Invalid or expired refresh token",
				Status:  401,
			},
			expectedJWT: nil,
			emailer:     nil,
			hibp:        nil,
			jwtTokenFn:  nil,
		},
		{
			name:   "user disabled",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				user := getSigninUser(userID)
				user.Disabled = true

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedToken),
						Type:             "regular",
					},
				).Return(user, nil)

				return mock
			},
			customClaimer: nil,
			request: api.PostTokenRequestObject{
				Body: &api.RefreshTokenRequest{
					RefreshToken: token.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-user",
				Message: "User is disabled",
				Status:  401,
			},
			expectedJWT: nil,
			emailer:     nil,
			hibp:        nil,
			jwtTokenFn:  nil,
		},
		{
			name:   "user anonymous",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				user := getSigninUser(userID)
				user.IsAnonymous = true

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedToken),
						Type:             "regular",
					},
				).Return(user, nil)

				return mock
			},
			customClaimer: nil,
			request: api.PostTokenRequestObject{
				Body: &api.RefreshTokenRequest{
					RefreshToken: token.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "forbidden-anonymous",
				Message: "Forbidden, user is anonymous.",
				Status:  403,
			},
			expectedJWT: nil,
			emailer:     nil,
			hibp:        nil,
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
				emailer:       nil,
				hibp:          nil,
			})

			resp := assertRequest(
				context.Background(), t, c.PostToken, tc.request, tc.expectedResponse,
			)

			resp200, ok := resp.(api.PostToken200JSONResponse)
			if ok {
				assertSession(t, jwtGetter, resp200.Session, tc.expectedJWT)
			}
		})
	}
}
