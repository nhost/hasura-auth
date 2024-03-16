package controller

import (
	"context"
	"log/slog"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
)

func (ctrl *Controller) PostUserPasswordReset( //nolint:ireturn
	ctx context.Context,
	request api.PostUserPasswordResetRequestObject,
) (api.PostUserPasswordResetResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	options, err := ctrl.validate.OptionsRedirectTo(request.Body.Options, logger)
	if err != nil {
		return ctrl.respondWithError(err), nil
	}
	request.Body.Options = options

	if !ctrl.validate.ValidateEmail(string(request.Body.Email)) {
		logger.Warn("email didn't pass access control checks")
		return ctrl.sendError(api.InvalidEmailPassword), nil
	}

	user, _, apiErr := ctrl.validate.GetUserByEmail(ctx, string(request.Body.Email), logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	ticket, apiErr := ctrl.SetTicket(ctx, user.ID, TicketTypePasswordReset, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	if err := ctrl.SendEmail(
		string(request.Body.Email),
		user.Locale,
		LinkTypePasswordReset,
		ticket,
		deptr(request.Body.Options.RedirectTo),
		notifications.TemplateNamePasswordReset,
		user.DisplayName,
		string(request.Body.Email),
		"",
		logger,
	); err != nil {
		return nil, err
	}

	return api.PostUserPasswordReset200JSONResponse(api.OK), nil
}
