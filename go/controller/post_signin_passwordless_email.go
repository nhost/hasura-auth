package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) postSigninPasswordlessEmailCreateUser(
	ctx context.Context,
	email string,
	options *api.SignUpOptions,
	logger *slog.Logger,
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

	insertedUser, err := ctrl.db.InsertUser(
		ctx, sql.InsertUserParams{
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
		},
	)
	if err != nil {
		logger.Error("error inserting user", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
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

func (ctrl *Controller) postSigninPasswordlessEmailValidateRequest(
	request api.PostSigninPasswordlessEmailRequestObject,
	logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	if !ctrl.config.EmailPasswordlessEnabled {
		logger.Warn("email passwordless signin is disabled")
		return nil, &APIError{api.DisabledEndpoint}
	}

	if !ctrl.validate.ValidateEmail(string(request.Body.Email)) {
		logger.Warn("email didn't pass access control checks")
		return nil, &APIError{api.InvalidEmailPassword}
	}

	options, apiErr := ctrl.validate.SignUpOptions(
		request.Body.Options, string(request.Body.Email), logger,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	return options, nil
}

func (ctrl *Controller) PostSigninPasswordlessEmail( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninPasswordlessEmailRequestObject,
) (api.PostSigninPasswordlessEmailResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	options, apiErr := ctrl.postSigninPasswordlessEmailValidateRequest(request, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	user, err := ctrl.db.GetUserByEmail(ctx, sql.Text(request.Body.Email))
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Info("user does not exist, creating user")

		user, apiErr = ctrl.postSigninPasswordlessEmailCreateUser(
			ctx, string(request.Body.Email), options, logger,
		)
		if apiErr != nil {
			return ctrl.respondWithError(apiErr), nil
		}
	} else if err != nil {
		logger.Error("error getting user by email", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	if user.Disabled {
		logger.Warn("user is disabled")
		return ctrl.sendError(api.DisabledUser), nil
	}

	ticket, apiErr := ctrl.SetTicket(ctx, user.ID, TicketTypePasswordLessEmail, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	if err := ctrl.SendEmail(
		string(request.Body.Email),
		user.Locale,
		LinkTypePasswordlessEmail,
		ticket,
		deptr(options.RedirectTo),
		notifications.TemplateNameSigninPasswordless,
		user.DisplayName,
		string(request.Body.Email),
		"",
		logger,
	); err != nil {
		return nil, err
	}

	return api.PostSigninPasswordlessEmail200JSONResponse(api.OK), nil
}
