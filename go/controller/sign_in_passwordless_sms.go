package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) SignInPasswordlessSms( //nolint:ireturn
	ctx context.Context,
	request api.SignInPasswordlessSmsRequestObject,
) (api.SignInPasswordlessSmsResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("phoneNumber", request.Body.PhoneNumber))

	if !ctrl.config.SMSPasswordlessEnabled {
		logger.Warn("SMS passwordless signin is disabled")
		return ctrl.sendError(ErrDisabledEndpoint), nil
	}

	options, apiErr := ctrl.signinSmsValidateRequest(
		request.Body.PhoneNumber, request.Body.Options, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	user, apiErr := ctrl.wf.GetUserByPhoneNumber(ctx, request.Body.PhoneNumber, logger)
	switch {
	case errors.Is(apiErr, ErrUserPhoneNumberNotFound):
		logger.Info("user does not exist, creating user")
		if apiErr := ctrl.postSigninPasswordlessSmsSignup(
			ctx, request.Body.PhoneNumber, options, logger,
		); apiErr != nil {
			logger.Error("error signing up user", logError(apiErr))
			return ctrl.respondWithError(apiErr), nil
		}
		return api.SignInPasswordlessSms200JSONResponse(api.OK), nil
	case apiErr != nil:
		logger.Error("error getting user by phone number", logError(apiErr))
		return ctrl.respondWithError(apiErr), nil
	}

	if apiErr := ctrl.postSigninPasswordlessSmsSignin(ctx, user, logger); apiErr != nil {
		logger.Error("error signing in user", logError(apiErr))
		return ctrl.respondWithError(apiErr), nil
	}

	return api.SignInPasswordlessSms200JSONResponse(api.OK), nil
}

func (ctrl *Controller) signinSmsValidateRequest(
	phoneNumber string,
	options *api.SignUpOptions,
	logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	options, apiErr := ctrl.wf.ValidateSignUpOptions(options, phoneNumber, logger)
	if apiErr != nil {
		return nil, apiErr
	}

	return options, nil
}

func (ctrl *Controller) postSigninPasswordlessSmsSignin(
	ctx context.Context,
	user sql.AuthUser,
	logger *slog.Logger,
) *APIError {
	otpHash, expiresAt, err := ctrl.wf.sms.SendVerificationCode(
		user.PhoneNumber.String,
		user.Locale,
	)
	if err != nil {
		logger.Error("error sending SMS verification code", logError(err))
		return ErrCannotSendSMS
	}

	if _, err := ctrl.wf.db.UpdateUserOTPHash(ctx, sql.UpdateUserOTPHashParams{
		ID:                user.ID,
		OtpHash:           sql.Text(otpHash),
		OtpHashExpiresAt:  sql.TimestampTz(expiresAt),
		OtpMethodLastUsed: sql.Text("sms"),
	}); err != nil {
		logger.Error("error updating user OTP hash", logError(err))
		return ErrInternalServerError
	}

	return nil
}

func (ctrl *Controller) postSigninPasswordlessSmsSignup(
	ctx context.Context,
	phoneNumber string,
	options *api.SignUpOptions,
	logger *slog.Logger,
) *APIError {
	ohash, expiresAt, err := ctrl.wf.sms.SendVerificationCode(
		phoneNumber,
		deptr(options.Locale),
	)
	if err != nil {
		logger.Error("error sending SMS verification code", logError(err))
		return ErrCannotSendSMS
	}

	var otpHash pgtype.Text
	if ohash != "" {
		otpHash = sql.Text(ohash)
	}

	apiErr := ctrl.wf.SignupUserWithouthSession(
		ctx,
		"", // email is empty for SMS signup
		options,
		false,
		func(
			_ pgtype.Text,
			_ pgtype.Timestamptz,
			metadata []byte,
			gravatarURL string,
		) error {
			_, err := ctrl.wf.db.InsertUser(ctx, sql.InsertUserParams{
				ID:                uuid.New(),
				Disabled:          ctrl.config.DisableNewUsers,
				DisplayName:       deptr(options.DisplayName),
				AvatarUrl:         gravatarURL,
				PhoneNumber:       sql.Text(phoneNumber),
				OtpHash:           otpHash,
				OtpHashExpiresAt:  sql.TimestampTz(expiresAt),
				OtpMethodLastUsed: sql.Text("sms"),
				Email:             pgtype.Text{},        //nolint:exhaustruct
				PasswordHash:      pgtype.Text{},        //nolint:exhaustruct
				Ticket:            pgtype.Text{},        //nolint:exhaustruct
				TicketExpiresAt:   pgtype.Timestamptz{}, //nolint:exhaustruct
				EmailVerified:     false,
				Locale:            deptr(options.Locale),
				DefaultRole:       deptr(options.DefaultRole),
				Metadata:          metadata,
				Roles:             deptr(options.AllowedRoles),
			})
			if err != nil {
				return fmt.Errorf("error inserting user: %w", err)
			}

			return nil
		},
		logger,
	)

	return apiErr
}
