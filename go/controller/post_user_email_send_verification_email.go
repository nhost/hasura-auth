package controller

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
)

func (ctrl *Controller) PostUserEmailSendVerificationEmail( //nolint:ireturn
	ctx context.Context,
	request api.PostUserEmailSendVerificationEmailRequestObject,
) (api.PostUserEmailSendVerificationEmailResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	options, apiErr := ctrl.wf.ValidateOptionsRedirectTo(request.Body.Options, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}
	request.Body.Options = options

	user, apiErr := ctrl.wf.GetUserByEmail(ctx, string(request.Body.Email), logger)
	switch {
	case errors.Is(apiErr, ErrUnverifiedUser):
	case apiErr == nil && !user.EmailVerified:
	case apiErr != nil:
		return ctrl.respondWithError(apiErr), nil
	default:
		return ctrl.respondWithError(ErrEmailAlreadyVerified), nil
	}

	ticket := generateTicket(TicketTypeVerifyEmail)
	expireAt := time.Now().Add(In30Days)
	if apiErr = ctrl.wf.SetTicket(ctx, user.ID, ticket, expireAt, logger); apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	if err := ctrl.wf.SendEmail(
		ctx,
		user.Email.String,
		user.Locale,
		LinkTypeEmailVerify,
		ticket,
		deptr(options.RedirectTo),
		notifications.TemplateNameEmailVerify,
		user.DisplayName,
		user.Email.String,
		"",
		logger,
	); err != nil {
		return ctrl.sendError(err), nil
	}

	return api.PostUserEmailSendVerificationEmail200JSONResponse(api.OK), nil
}
