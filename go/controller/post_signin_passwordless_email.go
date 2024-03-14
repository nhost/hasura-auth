package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
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
) (uuid.UUID, bool, error) {
	if ctrl.config.DisableSignup {
		logger.Warn("signup disabled")
		return uuid.UUID{}, false, &APIError{api.SignupDisabled}
	}

	metadata, err := json.Marshal(options.Metadata)
	if err != nil {
		logger.Error("error marshaling metadata", logError(err))
		return uuid.UUID{}, false, &APIError{api.InternalServerError}
	}

	gravatarURL := ctrl.gravatarURL(email)

	fmt.Println(6666, ctrl.config.DisableNewUsers)
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
		return uuid.UUID{}, false, &APIError{api.InternalServerError}
	}

	return insertedUser.UserID, ctrl.config.DisableNewUsers, nil
}

func (ctrl *Controller) PostSigninPasswordlessEmail( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninPasswordlessEmailRequestObject,
) (api.PostSigninPasswordlessEmailResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("email", string(request.Body.Email)))

	if !ctrl.config.EmailPasswordlessEnabled {
		logger.Warn("email passwordless signin is disabled")
		return ctrl.sendError(api.DisabledEndpoint), nil
	}

	if !ctrl.validator.emailValidator(string(request.Body.Email)) {
		logger.Warn("email didn't pass access control checks")
		return ctrl.sendError(api.InvalidEmailPassword), nil
	}

	options, err := ctrl.validator.postSignUpOptions(
		request.Body.Options, string(request.Body.Email), logger,
	)
	if err != nil {
		return ctrl.respondWithError(err), nil
	}

	user, err := ctrl.db.GetUserByEmail(ctx, sql.Text(request.Body.Email))
	userID := user.ID
	disabled := user.Disabled
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Info("user does not exist, creating user")

		userID, disabled, err = ctrl.postSigninPasswordlessEmailCreateUser(
			ctx, string(request.Body.Email), options, logger,
		)
		if err != nil {
			logger.Error("error validating signup request", logError(err))
			return ctrl.respondWithError(err), nil
		}
	}
	if err != nil {
		logger.Error("error getting user by email", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	if disabled {
		logger.Warn("user is disabled")
		return ctrl.sendError(api.DisabledUser), nil
	}

	user, err = ctrl.setTicket(ctx, userID, TicketTypePasswordLessEmail, logger)
	if err != nil {
		return ctrl.respondWithError(err), nil
	}

	if err := ctrl.sendEmail(
		string(request.Body.Email),
		user.Locale,
		LinkTypeEmailVerify,
		user.Ticket.String,
		deptr(options.RedirectTo),
		notifications.TemplateNameSigninPasswordless,
		user.DisplayName,
		string(request.Body.Email),
		"",
		logger,
	); err != nil {
		return nil, err
	}

	return nil, nil
}
