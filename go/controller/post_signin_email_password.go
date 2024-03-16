package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
)

func (ctrl *Controller) postSigninEmailPasswordWithTOTP( //nolint:ireturn
	ctx context.Context,
	userID uuid.UUID,
	logger *slog.Logger,
) (api.PostSigninEmailPasswordResponseObject, error) {
	ticket := "mfaTotp:" + uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Minute) //nolint:gomnd

	if apiErr := ctrl.SetTicket(ctx, userID, ticket, expiresAt, logger); apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
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

	user, _, apiErr := ctrl.wf.GetUserByEmail(
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

	session, err := ctrl.wf.NewSession(ctx, user, logger)
	if err != nil {
		logger.Error("error getting new session", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	return api.PostSigninEmailPassword200JSONResponse{
		Session: session,
		Mfa:     nil,
	}, nil
}
