package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
)

func (ctrl *Controller) ElevateWebauthn( //nolint:ireturn
	ctx context.Context,
	_ api.ElevateWebauthnRequestObject,
) (api.ElevateWebauthnResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	if !ctrl.config.WebauthnEnabled {
		logger.Error("webauthn is disabled")
		return ctrl.sendError(ErrDisabledEndpoint), nil
	}

	user, apiErr := ctrl.wf.GetUserFromJWTInContext(ctx, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	keys, apiErr := ctrl.wf.GetUserSecurityKeys(ctx, user.ID, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	if len(keys) == 0 {
		logger.Error("user has no security keys")
		return ctrl.sendError(ErrInvalidRequest), nil
	}

	creds, apiErr := webauthnCredentials(keys, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	waUser := WebauthnUser{
		ID:           user.ID,
		Name:         user.DisplayName,
		Email:        user.Email.String,
		Credentials:  creds,
		Discoverable: false,
	}

	creation, apiErr := ctrl.Webauthn.BeginLogin(waUser, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	return api.ElevateWebauthn200JSONResponse(creation.Response), nil
}
