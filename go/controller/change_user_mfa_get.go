package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) ChangeUserMfa( //nolint:ireturn
	ctx context.Context, _ api.ChangeUserMfaRequestObject,
) (api.ChangeUserMfaResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	if !ctrl.config.MfaEnabled {
		logger.Warn("mfa disabled")
		return ctrl.sendError(ErrDisabledEndpoint), nil
	}

	user, apiErr := ctrl.wf.GetUserFromJWTInContext(ctx, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	accountName := user.Email.String
	if accountName == "" {
		accountName = user.DisplayName
	}
	if accountName == "" {
		accountName = user.ID.String()
	}

	secret, imgBase64, err := ctrl.totp.Generate(accountName)
	if err != nil {
		logger.Error("failed to generate TOTP: %v", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	if err := ctrl.wf.db.UpdateUserTotpSecret(
		ctx, sql.UpdateUserTotpSecretParams{
			ID:         user.ID,
			TotpSecret: sql.Text(secret),
		},
	); err != nil {
		logger.Error("failed to update TOTP secret: %v", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	return api.ChangeUserMfa200JSONResponse{
		ImageUrl:   "data:image/png;base64," + imgBase64,
		TotpSecret: secret,
	}, nil
}
