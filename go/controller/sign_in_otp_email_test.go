package controller_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/controller"
	"github.com/nhost/hasura-auth/go/controller/mock"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/nhost/hasura-auth/go/testhelpers"
	"go.uber.org/mock/gomock"
)

func TestSignInOTPEmail(t *testing.T) { //nolint:maintidx
	t.Parallel()

	userID := uuid.MustParse("DB477732-48FA-4289-B694-2886A646B6EB")

	cases := []testRequest[api.SignInOTPEmailRequestObject, api.SignInOTPEmailResponseObject]{ //nolint:dupl
		{
			name:   "signup required",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient { //nolint:dupl
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				mock.EXPECT().InsertUser(
					gomock.Any(),
					cmpDBParams(sql.InsertUserParams{
						ID:                uuid.UUID{},
						Disabled:          false,
						DisplayName:       "jane@acme.com",
						AvatarUrl:         "",
						Email:             sql.Text("jane@acme.com"),
						PasswordHash:      pgtype.Text{}, //nolint:exhaustruct
						Ticket:            sql.Text("xxx"),
						TicketExpiresAt:   sql.TimestampTz(time.Now().Add(time.Hour)),
						EmailVerified:     false,
						Locale:            "en",
						DefaultRole:       "user",
						Metadata:          []byte("null"),
						Roles:             []string{"user", "me"},
						PhoneNumber:       pgtype.Text{},        //nolint:exhaustruct
						OtpHash:           pgtype.Text{},        //nolint:exhaustruct
						OtpHashExpiresAt:  pgtype.Timestamptz{}, //nolint:exhaustruct
						OtpMethodLastUsed: pgtype.Text{},        //nolint:exhaustruct
					},
						cmpopts.IgnoreFields(sql.InsertUserParams{}, "ID"), //nolint:exhaustruct
					),
				).Return(sql.InsertUserRow{
					UserID:    userID,
					CreatedAt: sql.TimestampTz(time.Now()),
				}, nil)

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: api.SignInOTPEmail200JSONResponse(api.OK),
			jwtTokenFn:       nil,
			expectedJWT:      nil,
			getControllerOpts: []getControllerOptsFunc{
				withEmailer(func(ctrl *gomock.Controller) *mock.MockEmailer {
					mock := mock.NewMockEmailer(ctrl)

					mock.EXPECT().SendEmail(
						gomock.Any(),
						"jane@acme.com",
						"en",
						notifications.TemplateNameSigninOTP,
						testhelpers.GomockCmpOpts(
							notifications.TemplateData{
								Link:        "",
								DisplayName: "jane@acme.com",
								Email:       "jane@acme.com",
								NewEmail:    "",
								Ticket:      "xxx",
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
				}),
			},
		},

		{
			name: "signup required - otp disabled",
			config: func() *controller.Config {
				config := getConfig()
				config.OTPEmailEnabled = false
				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)
				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "disabled-endpoint",
				Message: "This endpoint is disabled",
				Status:  409,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name: "signup required - email not allowed",
			config: func() *controller.Config {
				config := getConfig()
				config.AllowedEmails = []string{"sad@acme.com"}
				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)
				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "invalid-email-password",
				Message: "Incorrect email or password",
				Status:  401,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name: "signup required - role not allowed",
			config: func() *controller.Config {
				config := getConfig()
				config.DisableSignup = true
				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)
				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email: "jane@acme.com",
					Options: &api.SignUpOptions{
						AllowedRoles: &[]string{"admin"},
						DefaultRole:  nil,
						DisplayName:  nil,
						Locale:       nil,
						Metadata:     nil,
						RedirectTo:   nil,
					},
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "role-not-allowed",
				Message: "Role not allowed",
				Status:  400,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "signup required - locale not allowed",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient { //nolint:dupl
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				mock.EXPECT().InsertUser(
					gomock.Any(),
					cmpDBParams(sql.InsertUserParams{
						ID:                uuid.UUID{},
						Disabled:          false,
						DisplayName:       "jane@acme.com",
						AvatarUrl:         "",
						Email:             sql.Text("jane@acme.com"),
						PasswordHash:      pgtype.Text{}, //nolint:exhaustruct
						Ticket:            sql.Text("xxx"),
						TicketExpiresAt:   sql.TimestampTz(time.Now().Add(time.Hour)),
						EmailVerified:     false,
						Locale:            "en",
						DefaultRole:       "user",
						Metadata:          []byte("null"),
						Roles:             []string{"user", "me"},
						PhoneNumber:       pgtype.Text{},        //nolint:exhaustruct
						OtpHash:           pgtype.Text{},        //nolint:exhaustruct
						OtpHashExpiresAt:  pgtype.Timestamptz{}, //nolint:exhaustruct
						OtpMethodLastUsed: pgtype.Text{},        //nolint:exhaustruct
					},
						cmpopts.IgnoreFields(sql.InsertUserParams{}, "ID"), //nolint:exhaustruct
					),
				).Return(sql.InsertUserRow{
					UserID:    userID,
					CreatedAt: sql.TimestampTz(time.Now()),
				}, nil)

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email: "jane@acme.com",
					Options: &api.SignUpOptions{
						AllowedRoles: nil,
						DefaultRole:  nil,
						DisplayName:  nil,
						Locale:       ptr("xx"),
						Metadata:     nil,
						RedirectTo:   nil,
					},
				},
			},
			expectedResponse: api.SignInOTPEmail200JSONResponse(api.OK),
			jwtTokenFn:       nil,
			expectedJWT:      nil,
			getControllerOpts: []getControllerOptsFunc{
				withEmailer(func(ctrl *gomock.Controller) *mock.MockEmailer {
					mock := mock.NewMockEmailer(ctrl)

					mock.EXPECT().SendEmail(
						gomock.Any(),
						"jane@acme.com",
						"en",
						notifications.TemplateNameSigninOTP,
						testhelpers.GomockCmpOpts(
							notifications.TemplateData{
								Link:        "",
								DisplayName: "jane@acme.com",
								Email:       "jane@acme.com",
								NewEmail:    "",
								Ticket:      "xxx",
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
				}),
			},
		},

		{
			name:   "signup required - redirect not allowed",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)
				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email: "jane@acme.com",
					Options: &api.SignUpOptions{
						AllowedRoles: nil,
						DefaultRole:  nil,
						DisplayName:  nil,
						Locale:       nil,
						Metadata:     nil,
						RedirectTo:   ptr("https://evil.com"),
					},
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "redirectTo-not-allowed",
				Message: `The value of "options.redirectTo" is not allowed.`,
				Status:  400,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name: "signup required - options",
			config: func() *controller.Config {
				config := getConfig()
				config.AllowedLocales = []string{"en", "fr"}
				config.AllowedRedirectURLs = []string{"http://myapp"}
				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				mock.EXPECT().InsertUser(
					gomock.Any(),
					cmpDBParams(sql.InsertUserParams{
						ID:                uuid.UUID{},
						Disabled:          false,
						DisplayName:       "Jane Doe",
						AvatarUrl:         "",
						Email:             sql.Text("jane@acme.com"),
						PasswordHash:      pgtype.Text{}, //nolint:exhaustruct
						Ticket:            sql.Text("xxx"),
						TicketExpiresAt:   sql.TimestampTz(time.Now().Add(time.Hour)),
						EmailVerified:     false,
						Locale:            "fr",
						DefaultRole:       "user",
						Metadata:          []byte(`{"asd":"asd"}`),
						Roles:             []string{"user"},
						PhoneNumber:       pgtype.Text{},        //nolint:exhaustruct
						OtpHash:           pgtype.Text{},        //nolint:exhaustruct
						OtpHashExpiresAt:  pgtype.Timestamptz{}, //nolint:exhaustruct
						OtpMethodLastUsed: pgtype.Text{},        //nolint:exhaustruct
					},
						cmpopts.IgnoreFields(sql.InsertUserParams{}, "ID"), //nolint:exhaustruct
					),
				).Return(sql.InsertUserRow{
					UserID:    userID,
					CreatedAt: sql.TimestampTz(time.Now()),
				}, nil)

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email: "jane@acme.com",
					Options: &api.SignUpOptions{
						AllowedRoles: &[]string{"user"},
						DefaultRole:  ptr("user"),
						DisplayName:  ptr("Jane Doe"),
						Locale:       ptr("fr"),
						Metadata:     &map[string]any{"asd": "asd"},
						RedirectTo:   ptr("http://myapp"),
					},
				},
			},
			expectedResponse: api.SignInOTPEmail200JSONResponse(api.OK),
			jwtTokenFn:       nil,
			expectedJWT:      nil,
			getControllerOpts: []getControllerOptsFunc{
				withEmailer(func(ctrl *gomock.Controller) *mock.MockEmailer {
					mock := mock.NewMockEmailer(ctrl)

					mock.EXPECT().SendEmail(
						gomock.Any(),
						"jane@acme.com",
						"fr",
						notifications.TemplateNameSigninOTP,
						testhelpers.GomockCmpOpts(
							notifications.TemplateData{
								Link:        "",
								DisplayName: "Jane Doe",
								Email:       "jane@acme.com",
								NewEmail:    "",
								Ticket:      "xxx",
								RedirectTo:  "http://myapp",
								Locale:      "fr",
								ServerURL:   "https://local.auth.nhost.run",
								ClientURL:   "http://localhost:3000",
							},
							testhelpers.FilterPathLast(
								[]string{".Ticket"}, cmp.Comparer(cmpTicket)),

							testhelpers.FilterPathLast(
								[]string{".Link"}, cmp.Comparer(cmpLink)),
						)).Return(nil)

					return mock
				}),
			},
		},

		{
			name: "signup required - signup disabled",
			config: func() *controller.Config {
				config := getConfig()
				config.DisableSignup = true
				return config
			},
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{}, pgx.ErrNoRows) //nolint:exhaustruct

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error:   "signup-disabled",
				Message: "Sign up is disabled.",
				Status:  403,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},

		{
			name:   "signup not required",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{
					ID:                       userID,
					CreatedAt:                pgtype.Timestamptz{}, //nolint:exhaustruct
					UpdatedAt:                pgtype.Timestamptz{}, //nolint:exhaustruct
					LastSeen:                 pgtype.Timestamptz{}, //nolint:exhaustruct
					Disabled:                 false,
					DisplayName:              "jane@acme.com",
					AvatarUrl:                "",
					Locale:                   "en",
					Email:                    sql.Text("jane@acme.com"),
					PhoneNumber:              pgtype.Text{}, //nolint:exhaustruct
					PasswordHash:             pgtype.Text{}, //nolint:exhaustruct
					EmailVerified:            false,
					PhoneNumberVerified:      false,
					NewEmail:                 pgtype.Text{},        //nolint:exhaustruct
					OtpMethodLastUsed:        pgtype.Text{},        //nolint:exhaustruct
					OtpHash:                  pgtype.Text{},        //nolint:exhaustruct
					OtpHashExpiresAt:         pgtype.Timestamptz{}, //nolint:exhaustruct
					DefaultRole:              "",
					IsAnonymous:              false,
					TotpSecret:               pgtype.Text{},        //nolint:exhaustruct
					ActiveMfaType:            pgtype.Text{},        //nolint:exhaustruct
					Ticket:                   pgtype.Text{},        //nolint:exhaustruct
					TicketExpiresAt:          pgtype.Timestamptz{}, //nolint:exhaustruct
					Metadata:                 []byte{},
					WebauthnCurrentChallenge: pgtype.Text{}, //nolint:exhaustruct
				}, nil)

				mock.EXPECT().UpdateUserTicket(
					gomock.Any(),
					cmpDBParams(sql.UpdateUserTicketParams{
						ID:              userID,
						Ticket:          sql.Text("xxx"),
						TicketExpiresAt: sql.TimestampTz(time.Now().Add(time.Hour)),
					}),
				).Return(userID, nil)

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: api.SignInOTPEmail200JSONResponse(api.OK),
			jwtTokenFn:       nil,
			expectedJWT:      nil,
			getControllerOpts: []getControllerOptsFunc{
				withEmailer(func(ctrl *gomock.Controller) *mock.MockEmailer {
					mock := mock.NewMockEmailer(ctrl)

					mock.EXPECT().SendEmail(
						gomock.Any(),
						"jane@acme.com",
						"en",
						notifications.TemplateNameSigninOTP,
						testhelpers.GomockCmpOpts(
							notifications.TemplateData{
								Link:        "",
								DisplayName: "jane@acme.com",
								Email:       "jane@acme.com",
								NewEmail:    "",
								Ticket:      "xxx",
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
				}),
			},
		},

		{
			name:   "signup not required - user disabled",
			config: getConfig,
			db: func(ctrl *gomock.Controller) controller.DBClient {
				mock := mock.NewMockDBClient(ctrl)

				mock.EXPECT().GetUserByEmail(
					gomock.Any(),
					sql.Text("jane@acme.com"),
				).Return(sql.AuthUser{
					ID:                       userID,
					CreatedAt:                pgtype.Timestamptz{}, //nolint:exhaustruct
					UpdatedAt:                pgtype.Timestamptz{}, //nolint:exhaustruct
					LastSeen:                 pgtype.Timestamptz{}, //nolint:exhaustruct
					Disabled:                 true,
					DisplayName:              "jane@acme.com",
					AvatarUrl:                "",
					Locale:                   "en",
					Email:                    sql.Text("jane@acme.com"),
					PhoneNumber:              pgtype.Text{}, //nolint:exhaustruct
					PasswordHash:             pgtype.Text{}, //nolint:exhaustruct
					EmailVerified:            false,
					PhoneNumberVerified:      false,
					NewEmail:                 pgtype.Text{},        //nolint:exhaustruct
					OtpMethodLastUsed:        pgtype.Text{},        //nolint:exhaustruct
					OtpHash:                  pgtype.Text{},        //nolint:exhaustruct
					OtpHashExpiresAt:         pgtype.Timestamptz{}, //nolint:exhaustruct
					DefaultRole:              "",
					IsAnonymous:              false,
					TotpSecret:               pgtype.Text{},        //nolint:exhaustruct
					ActiveMfaType:            pgtype.Text{},        //nolint:exhaustruct
					Ticket:                   pgtype.Text{},        //nolint:exhaustruct
					TicketExpiresAt:          pgtype.Timestamptz{}, //nolint:exhaustruct
					Metadata:                 []byte{},
					WebauthnCurrentChallenge: pgtype.Text{}, //nolint:exhaustruct
				}, nil)

				return mock
			},
			request: api.SignInOTPEmailRequestObject{
				Body: &api.SignInOTPEmailJSONRequestBody{
					Email:   "jane@acme.com",
					Options: nil,
				},
			},
			expectedResponse: controller.ErrorResponse{
				Error: "disabled-user", Message: "User is disabled", Status: 401,
			},
			jwtTokenFn:        nil,
			expectedJWT:       nil,
			getControllerOpts: []getControllerOptsFunc{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			c, _ := getController(t, ctrl, tc.config, tc.db, tc.getControllerOpts...)

			assertRequest(
				t.Context(),
				t,
				c.SignInOTPEmail,
				tc.request,
				tc.expectedResponse,
			)
		})
	}
}
