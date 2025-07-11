package controller

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
)

func (ctrl *Controller) postSignupWebauthnValidateRequest(
	request api.SignUpWebauthnRequestObject,
	logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	if !ctrl.config.WebauthnEnabled {
		logger.Error("webauthn is disabled")
		return nil, ErrDisabledEndpoint
	}

	if ctrl.config.DisableSignup {
		logger.Error("signup is disabled")
		return nil, ErrSignupDisabled
	}

	options, apiErr := ctrl.wf.ValidateSignUpOptions(
		request.Body.Options, string(request.Body.Email), logger,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	if apiErr := ctrl.wf.ValidateSignupEmail(request.Body.Email, logger); apiErr != nil {
		return nil, apiErr
	}

	return options, nil
}

func (ctrl *Controller) SignUpWebauthn( //nolint:ireturn
	ctx context.Context,
	request api.SignUpWebauthnRequestObject,
) (api.SignUpWebauthnResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	options, apiErr := ctrl.postSignupWebauthnValidateRequest(request, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	user := WebauthnUser{
		ID:           uuid.New(),
		Name:         deptr(options.DisplayName),
		Email:        string(request.Body.Email),
		Credentials:  nil,
		Discoverable: false,
	}

	creation, apiErr := ctrl.Webauthn.BeginRegistration(user, options, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	return api.SignUpWebauthn200JSONResponse(creation.Response), nil
}
