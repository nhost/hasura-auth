//go:generate mockgen -package mock -destination mock/controller.go --source=controller.go
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
)

func deptr[T any](x *T) T { //nolint:ireturn
	if x == nil {
		return *new(T)
	}
	return *x
}

func ptr[T any](x T) *T {
	return &x
}

type Emailer interface {
	SendEmail(
		to string,
		locale string,
		templateName notifications.TemplateName,
		data notifications.TemplateData,
	) error
}

type DBClient interface {
	GetUser(ctx context.Context, id uuid.UUID) (sql.AuthUser, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (sql.AuthUser, error)
	GetUserByRefreshTokenHash(
		ctx context.Context, arg sql.GetUserByRefreshTokenHashParams,
	) (sql.AuthUser, error)
	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]sql.AuthUserRole, error)
	InsertUser(ctx context.Context, arg sql.InsertUserParams) (sql.InsertUserRow, error)
	InsertUserWithRefreshToken(
		ctx context.Context, arg sql.InsertUserWithRefreshTokenParams,
	) (sql.InsertUserWithRefreshTokenRow, error)
	InsertRefreshtoken(ctx context.Context, arg sql.InsertRefreshtokenParams) (uuid.UUID, error)
	UpdateUserChangeEmail(
		ctx context.Context,
		arg sql.UpdateUserChangeEmailParams,
	) (sql.UpdateUserChangeEmailRow, error)
	UpdateUserLastSeen(ctx context.Context, id uuid.UUID) (pgtype.Timestamptz, error)
	UpdateUserTicket(ctx context.Context, arg sql.UpdateUserTicketParams) (uuid.UUID, error)
}

type Controller struct {
	db DBClient
	// validate object helps to validate the input. It provides functionality to validate
	// input objects and, as well, to fetch users and other data from the database
	// validating the user in the process
	validate    *Validator
	config      Config
	gravatarURL func(string) string
	jwtGetter   *JWTGetter
	email       Emailer
	version     string
}

func New(
	db DBClient,
	config Config,
	jwtGetter *JWTGetter,
	emailer Emailer,
	hibp HIBPClient,
	version string,
) (*Controller, error) {
	validator, err := NewValidator(&config, db, hibp)
	if err != nil {
		return nil, fmt.Errorf("error creating validator: %w", err)
	}

	return &Controller{
		db:       db,
		config:   config,
		validate: validator,
		gravatarURL: GravatarURLFunc(
			config.GravatarEnabled, config.GravatarDefault, config.GravatarRating,
		),
		jwtGetter: jwtGetter,
		email:     emailer,
		version:   version,
	}, nil
}

type SignUpFn func(input *sql.InsertUserParams) error

func SignupUserWithTicket(ticket string, expiresAt time.Time) SignUpFn {
	return func(input *sql.InsertUserParams) error {
		input.Ticket = sql.Text(ticket)
		input.TicketExpiresAt = sql.TimestampTz(expiresAt)
		return nil
	}
}

func SignupUserWithPassword(password string) SignUpFn {
	return func(input *sql.InsertUserParams) error {
		hashedPassword, err := hashPassword(password)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}

		input.PasswordHash = sql.Text(hashedPassword)

		return nil
	}
}

func (ctrl *Controller) SignUpUser(
	ctx context.Context,
	email string,
	options *api.SignUpOptions,
	logger *slog.Logger,
	withInputFn ...SignUpFn,
) (sql.AuthUser, *APIError) {
	if ctrl.config.DisableSignup {
		logger.Warn("signup disabled")
		return sql.AuthUser{}, &APIError{api.SignupDisabled} //nolint:exhaustruct
	}

	metadata, err := json.Marshal(options.Metadata)
	if err != nil {
		logger.Error("error marshaling metadata", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	gravatarURL := ctrl.gravatarURL(email)

	input := sql.InsertUserParams{
		Disabled:        ctrl.config.DisableNewUsers,
		DisplayName:     deptr(options.DisplayName),
		AvatarUrl:       gravatarURL,
		Email:           sql.Text(email),
		PasswordHash:    pgtype.Text{}, //nolint:exhaustruct
		Ticket:          pgtype.Text{}, //nolint:exhaustruct
		TicketExpiresAt: sql.TimestampTz(time.Now()),
		EmailVerified:   false,
		Locale:          deptr(options.Locale),
		DefaultRole:     deptr(options.DefaultRole),
		Metadata:        metadata,
		Roles:           deptr(options.AllowedRoles),
	}

	for _, fn := range withInputFn {
		if err := fn(&input); err != nil {
			logger.Error("error applying input function to user signup", logError(err))
			return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
		}
	}

	insertedUser, err := ctrl.db.InsertUser(ctx, input)
	if err != nil {
		logger.Error("error inserting user", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if ctrl.config.DisableNewUsers {
		logger.Warn("new user disabled")
		return sql.AuthUser{}, &APIError{api.DisabledUser} //nolint:exhaustruct
	}

	return sql.AuthUser{ //nolint:exhaustruct
		ID:                  insertedUser.UserID,
		Disabled:            ctrl.config.DisableNewUsers,
		DisplayName:         deptr(options.DisplayName),
		AvatarUrl:           gravatarURL,
		Locale:              deptr(options.Locale),
		Email:               sql.Text(email),
		EmailVerified:       false,
		PhoneNumberVerified: false,
		DefaultRole:         deptr(options.DefaultRole),
		IsAnonymous:         false,
		Metadata:            metadata,
	}, nil
}

func (ctrl *Controller) SendEmail(
	to string,
	locale string,
	linkType LinkType,
	ticket string,
	redirectTo string,
	templateName notifications.TemplateName,
	displayName string,
	email string,
	newEmail string,
	logger *slog.Logger,
) error {
	link, err := GenLink(
		*ctrl.config.ServerURL,
		linkType,
		ticket,
		redirectTo,
	)
	if err != nil {
		logger.Error("problem generating email verification link", logError(err))
		return fmt.Errorf("problem generating email verification link: %w", err)
	}

	if err := ctrl.email.SendEmail(
		to,
		locale,
		templateName,
		notifications.TemplateData{
			Link:        link,
			DisplayName: displayName,
			Email:       email,
			NewEmail:    newEmail,
			Ticket:      ticket,
			RedirectTo:  redirectTo,
			Locale:      locale,
			ServerURL:   ctrl.config.ServerURL.String(),
			ClientURL:   ctrl.config.ClientURL.String(),
		},
	); err != nil {
		logger.Error("problem sending email", logError(err))
		return fmt.Errorf("problem sending email: %w", err)
	}

	return nil
}
