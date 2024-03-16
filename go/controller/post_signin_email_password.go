package controller

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) postSigninEmailPasswordWithTOTP( //nolint:ireturn
	ctx context.Context,
	userID uuid.UUID,
	logger *slog.Logger,
) (api.PostSigninEmailPasswordResponseObject, error) {
	ticket := "mfaTotp:" + uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Minute) //nolint:gomnd

	if _, err := ctrl.db.UpdateUserTicket(ctx, sql.UpdateUserTicketParams{
		ID:              userID,
		Ticket:          sql.Text(ticket),
		TicketExpiresAt: sql.TimestampTz(expiresAt),
	}); err != nil {
		logger.Error("error updating user ticket", logError(err))
		return nil, fmt.Errorf("error updating user ticket: %w", err)
	}

	return api.PostSigninEmailPassword200JSONResponse{
		Mfa: &api.MFAChallengePayload{
			Ticket: ticket,
		},
		Session: nil,
	}, nil
}

func (ctrl *Controller) PostSigninEmailPassword( //nolint:ireturn
	ctx context.Context, request api.PostSigninEmailPasswordRequestObject,
) (api.PostSigninEmailPasswordResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	user, apiErr := ctrl.validate.GetUserByEmail(
		ctx, string(request.Body.Email), logger.WithGroup("validator"),
	)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	if !verifyHashPassword(request.Body.Password, user.PasswordHash.String) {
		logger.Warn("password doesn't match")
		return ctrl.sendError(api.InvalidEmailPassword), nil
	}

	if user.ActiveMfaType.String == "totp" {
		return ctrl.postSigninEmailPasswordWithTOTP(ctx, user.ID, logger)
	}

	session, err := ctrl.GetNewSession(ctx, user, logger)
	if err != nil {
		logger.Error("error getting new session", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	return api.PostSigninEmailPassword200JSONResponse{
		Session: session,
		Mfa:     nil,
	}, nil
}
