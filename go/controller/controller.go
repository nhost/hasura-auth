//go:generate mockgen -package mock -destination mock/controller.go --source=controller.go
package controller

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	db          DBClient
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
