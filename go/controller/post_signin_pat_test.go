package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/controller/mock"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/nhost/hasura-auth/go/testhelpers"
	"go.uber.org/mock/gomock"
)

//nolint:dupl
func TestPostSigninPat(t *testing.T) { //nolint:maintidx,gocognit,cyclop
	t.Parallel()

	userID := uuid.MustParse("db477732-48fa-4289-b694-2886a646b6eb")
	pat := uuid.MustParse("1fb17604-86c7-444e-b337-09a644465f2d")
	hashedPat := `\x9698157153010b858587119503cbeef0cf288f11775e51cdb6bfd65e930d9310`

	cases := []struct {
		name             string
		config           func() *controller.Config
		db               func(ctrl *gomock.Controller) controller.DBClient
		emailer          func(ctrl *gomock.Controller) *mock.MockEmailer
		hibp             func(ctrl *gomock.Controller) *mock.MockHIBPClient
		customClaimer    func(ctrl *gomock.Controller) controller.CustomClaimer
		request          api.PostSigninPatRequestObject
		expectedResponse api.PostSigninPatResponseObject
		expectedJWT      *jwt.Token
	}{
		{
			name:   "simple",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(getSigninUser(userID), nil)

				mock.EXPECT().GetUserRoles(
					gomock.Any(), userID,
				).Return([]sql.AuthUserRole{
					{UserID: userID, Role: "user"}, //nolint:exhaustruct
					{UserID: userID, Role: "me"},   //nolint:exhaustruct
				}, nil)

				mock.EXPECT().InsertRefreshtoken(
					gomock.Any(),
					cmpDBParams(sql.InsertRefreshtokenParams{
						UserID:           userID,
						RefreshTokenHash: pgtype.Text{}, //nolint:exhaustruct
						ExpiresAt:        sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
						Type:             sql.RefreshTokenTypeRegular,
						Metadata:         nil,
					}),
				).Return(uuid.New(), nil)

				mock.EXPECT().UpdateUserLastSeen(
					gomock.Any(), userID,
				).Return(sql.TimestampTz(time.Now()), nil)

				return mock
			},
			customClaimer: nil,
			hibp:          mock.NewMockHIBPClient,
			emailer:       mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: api.PostSigninPat200JSONResponse{
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
						"x-hasura-allowed-roles":    []any{"user", "me"},
						"x-hasura-default-role":     "user",
						"x-hasura-user-id":          "db477732-48fa-4289-b694-2886a646b6eb",
						"x-hasura-user-isAnonymous": "false",
					},
					"iat": float64(time.Now().Unix()),
					"iss": "hasura-auth",
					"sub": "db477732-48fa-4289-b694-2886a646b6eb",
				},
				Signature: []byte{},
				Valid:     true,
			},
		},

		{
			name:   "with custom claims",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(getSigninUser(userID), nil)

				mock.EXPECT().GetUserRoles(
					gomock.Any(), userID,
				).Return([]sql.AuthUserRole{
					{UserID: userID, Role: "user"}, //nolint:exhaustruct
					{UserID: userID, Role: "me"},   //nolint:exhaustruct
				}, nil)

				mock.EXPECT().InsertRefreshtoken(
					gomock.Any(),
					cmpDBParams(sql.InsertRefreshtokenParams{
						UserID:           userID,
						RefreshTokenHash: pgtype.Text{}, //nolint:exhaustruct
						ExpiresAt:        sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
						Type:             sql.RefreshTokenTypeRegular,
						Metadata:         nil,
					}),
				).Return(uuid.New(), nil)

				mock.EXPECT().UpdateUserLastSeen(
					gomock.Any(), userID,
				).Return(sql.TimestampTz(time.Now()), nil)

				return mock
			},
			customClaimer: func(ctrl *gomock.Controller) controller.CustomClaimer {
				mock := mock.NewMockCustomClaimer(ctrl)
				mock.EXPECT().GetClaims(
					gomock.Any(),
					"db477732-48fa-4289-b694-2886a646b6eb",
				).Return(map[string]any{
					"claim1":      "value1",
					"claim2":      "value2",
					"claimArray":  []any{"value1", "value2"},
					"claimObject": map[string]any{"key1": "value1", "key2": "value2"},
					"claimNil":    nil,
				}, nil)
				return mock
			},
			hibp:    mock.NewMockHIBPClient,
			emailer: mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: api.PostSigninPat200JSONResponse{
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
						"x-hasura-allowed-roles":    []any{"user", "me"},
						"x-hasura-claim1":           "value1",
						"x-hasura-claim2":           "value2",
						"x-hasura-claimarray":       `{"value1","value2"}`,
						"x-hasura-claimnil":         "null",
						"x-hasura-claimobject":      `{"key1":"value1","key2":"value2"}`,
						"x-hasura-default-role":     "user",
						"x-hasura-user-id":          "db477732-48fa-4289-b694-2886a646b6eb",
						"x-hasura-user-isAnonymous": "false",
					},
					"iat": float64(time.Now().Unix()),
					"iss": "hasura-auth",
					"sub": "db477732-48fa-4289-b694-2886a646b6eb",
				},
				Signature: []byte{},
				Valid:     true,
			},
		},

		{
			name: "user not found",
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
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				return mock
			},
			customClaimer: nil,
			hibp:          mock.NewMockHIBPClient,
			emailer:       mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-pat",
				Message: "Invalid or expired personal access token",
				Status:  401,
			},
			expectedJWT: nil,
		},

		{
			name: "email not allowed",
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
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(getSigninUser(userID), nil)

				return mock
			},
			customClaimer: nil,
			hibp:          mock.NewMockHIBPClient,
			emailer:       mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-email-password",
				Message: "Incorrect email or password",
				Status:  401,
			},
			expectedJWT: nil,
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
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(user, nil)

				return mock
			},
			customClaimer: nil,
			hibp:          mock.NewMockHIBPClient,
			emailer:       mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-user",
				Message: "User is disabled",
				Status:  401,
			},
			expectedJWT: nil,
		},

		{
			name: "verification required",
			config: func() *controller.Config {
				cfg := getConfig()
				cfg.RequireEmailVerification = true
				return cfg
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				user := getSigninUser(userID)
				user.EmailVerified = false

				mock.EXPECT().GetUserByRefreshTokenHash(
					gomock.Any(),
					sql.GetUserByRefreshTokenHashParams{
						RefreshTokenHash: sql.Text(hashedPat),
						Type:             "pat",
					},
				).Return(user, nil)

				return mock
			},
			customClaimer: nil,
			hibp:          mock.NewMockHIBPClient,
			emailer:       mock.NewMockEmailer,
			request: api.PostSigninPatRequestObject{
				Body: &api.SignInPATRequest{
					PersonalAccessToken: pat.String(),
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "unverified-user",
				Message: "User is not verified.",
				Status:  401,
			},
			expectedJWT: nil,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc := tc

			ctrl := gomock.NewController(t)

			var cc controller.CustomClaimer
			if tc.customClaimer != nil {
				cc = tc.customClaimer(ctrl)
			}

			jwtGetter, err := controller.NewJWTGetter(
				jwtSecret,
				time.Second*time.Duration(tc.config().AccessTokenExpiresIn),
				cc,
			)
			if err != nil {
				t.Fatalf("failed to create jwt getter: %v", err)
			}

			c, err := controller.New(
				tc.db(ctrl),
				*tc.config(),
				jwtGetter,
				tc.emailer(ctrl),
				tc.hibp(ctrl),
				"dev",
			)
			if err != nil {
				t.Fatalf("failed to create controller: %v", err)
			}

			resp, err := c.PostSigninPat(context.Background(), tc.request)
			if err != nil {
				t.Fatalf("failed to post signin email password: %v", err)
			}

			if diff := cmp.Diff(
				resp, tc.expectedResponse,
				testhelpers.FilterPathLast(
					[]string{".CreatedAt"}, cmpopts.EquateApproxTime(time.Minute),
				),
				testhelpers.FilterPathLast(
					[]string{".Ticket"},
					cmp.Comparer(cmpTicket),
				),
				cmpopts.IgnoreFields(api.Session{}, "RefreshToken", "AccessToken"), //nolint:exhaustruct
			); diff != "" {
				t.Fatalf("unexpected response: %s", diff)
			}

			resp200, ok := resp.(api.PostSigninPat200JSONResponse)
			if ok { //nolint:nestif
				var token *jwt.Token
				if resp200.Session == nil {
					token = nil
				} else {
					token, err = jwtGetter.Validate(resp200.Session.AccessToken)
					if err != nil {
						t.Fatalf("failed to get claims: %v", err)
					}
				}
				if diff := cmp.Diff(
					token,
					tc.expectedJWT,
					cmpopts.IgnoreFields(jwt.Token{}, "Raw", "Signature"), //nolint:exhaustruct
					cmpopts.EquateApprox(0, 10),
				); diff != "" {
					t.Fatalf("unexpected jwt: %s", diff)
				}
			}
		})
	}
}
