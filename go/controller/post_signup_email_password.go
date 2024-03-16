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
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/oapi-codegen/runtime/types"
)

func (ctrl *Controller) postSignupEmailPasswordValidateRequest(
	ctx context.Context, req api.PostSignupEmailPasswordRequestObject, logger *slog.Logger,
) (api.PostSignupEmailPasswordRequestObject, *APIError) {
	if ctrl.config.DisableSignup {
		logger.Warn("signup disabled")
		return api.PostSignupEmailPasswordRequestObject{}, //nolint:exhaustruct
			&APIError{api.SignupDisabled}
	}

	if err := ctrl.validate.SignupEmail(ctx, req.Body.Email, logger); err != nil {
		return api.PostSignupEmailPasswordRequestObject{}, err //nolint:exhaustruct
	}

	if err := ctrl.validate.Password(ctx, req.Body.Password, logger); err != nil {
		return api.PostSignupEmailPasswordRequestObject{}, err //nolint:exhaustruct
	}

	options, err := ctrl.validate.SignUpOptions(
		req.Body.Options, string(req.Body.Email), logger,
	)
	if err != nil {
		return api.PostSignupEmailPasswordRequestObject{}, err //nolint:exhaustruct
	}

	req.Body.Options = options

	return req, nil
}

func (ctrl *Controller) PostSignupEmailPassword( //nolint:ireturn
	ctx context.Context,
	req api.PostSignupEmailPasswordRequestObject,
) (api.PostSignupEmailPasswordResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("email", string(req.Body.Email)))

	req, apiError := ctrl.postSignupEmailPasswordValidateRequest(ctx, req, logger)
	if apiError != nil {
		return ctrl.respondWithError(apiError), nil
	}

	hashedPassword, err := hashPassword(req.Body.Password)
	if err != nil {
		logger.Error("error hashing password", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	metadata, err := json.Marshal(req.Body.Options.Metadata)
	if err != nil {
		logger.Error("error marshaling metadata", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	gravatarURL := ctrl.gravatarURL(string(req.Body.Email))

	if ctrl.config.RequireEmailVerification || ctrl.config.DisableNewUsers {
		return ctrl.postSignupEmailPasswordWithEmailVerificationOrUserDisabled(
			ctx,
			string(req.Body.Email),
			req.Body.Password,
			req.Body.Options,
			logger,
		)
	}

	return ctrl.postSignupEmailPasswordWithoutEmailVerification(
		ctx,
		sql.Text(req.Body.Email),
		hashedPassword,
		gravatarURL,
		req.Body.Options,
		metadata,
		logger,
	)
}

func (ctrl *Controller) postSignupEmailPasswordWithEmailVerificationOrUserDisabled( //nolint:ireturn
	ctx context.Context,
	email string,
	password string,
	options *api.SignUpOptions,
	logger *slog.Logger,
) (api.PostSignupEmailPasswordResponseObject, error) {
	ticket := newTicket(TicketTypeVerifyEmail)

	if _, err := ctrl.SignUpUser(
		ctx,
		email,
		options,
		logger,
		SignupUserWithTicket(ticket, time.Now().Add(30*24*time.Hour)),
		SignupUserWithPassword(password),
	); err != nil {
		return ctrl.respondWithError(err), nil
	}

	if ctrl.config.DisableNewUsers {
		return api.PostSignupEmailPassword200JSONResponse{Session: nil}, nil
	}

	if err := ctrl.SendEmail(
		email,
		deptr(options.Locale),
		LinkTypeEmailVerify,
		ticket,
		deptr(options.RedirectTo),
		notifications.TemplateNameEmailVerify,
		deptr(options.DisplayName),
		email,
		"",
		logger,
	); err != nil {
		return nil, err
	}

	return api.PostSignupEmailPassword200JSONResponse{Session: nil}, nil
}

func (ctrl *Controller) postSignupEmailPasswordWithoutEmailVerification( //nolint:ireturn,funlen
	ctx context.Context,
	email pgtype.Text,
	hashedPassword string,
	gravatarURL string,
	options *api.SignUpOptions,
	metadata []byte,
	logger *slog.Logger,
) (api.PostSignupEmailPasswordResponseObject, error) {
	refreshToken := uuid.New()
	expiresAt := time.Now().Add(time.Duration(ctrl.config.RefreshTokenExpiresIn) * time.Second)
	user, err := ctrl.db.InsertUserWithRefreshToken(
		ctx, sql.InsertUserWithRefreshTokenParams{
			Disabled:              ctrl.config.DisableNewUsers,
			DisplayName:           deptr(options.DisplayName),
			AvatarUrl:             gravatarURL,
			Email:                 email,
			PasswordHash:          sql.Text(hashedPassword),
			Ticket:                pgtype.Text{}, //nolint:exhaustruct
			TicketExpiresAt:       sql.TimestampTz(time.Now()),
			EmailVerified:         false,
			Locale:                deptr(options.Locale),
			DefaultRole:           deptr(options.DefaultRole),
			Metadata:              metadata,
			Roles:                 deptr(options.AllowedRoles),
			RefreshTokenHash:      sql.Text(hashRefreshToken([]byte(refreshToken.String()))),
			RefreshTokenExpiresAt: sql.TimestampTz(expiresAt),
		},
	)
	if err != nil {
		logger.Error("error inserting user", logError(err))
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	accessToken, expiresIn, err := ctrl.jwtGetter.GetToken(
		ctx, user.UserID, deptr(options.AllowedRoles), *options.DefaultRole, logger,
	)
	if err != nil {
		logger.Error("error getting jwt", logError(err))
		return nil, fmt.Errorf("error getting jwt: %w", err)
	}

	return api.PostSignupEmailPassword200JSONResponse{
		Session: &api.Session{
			AccessToken:          accessToken,
			AccessTokenExpiresIn: expiresIn,
			RefreshToken:         refreshToken.String(),
			User: &api.User{
				AvatarUrl:           gravatarURL,
				CreatedAt:           user.CreatedAt.Time,
				DefaultRole:         *options.DefaultRole,
				DisplayName:         deptr(options.DisplayName),
				Email:               types.Email(email.String),
				EmailVerified:       false,
				Id:                  user.UserID.String(),
				IsAnonymous:         false,
				Locale:              deptr(options.Locale),
				Metadata:            deptr(options.Metadata),
				PhoneNumber:         "",
				PhoneNumberVerified: false,
				Roles:               deptr(options.AllowedRoles),
			},
		},
	}, nil
}
