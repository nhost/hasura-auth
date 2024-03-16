package controller

import (
	"context"
	"log/slog"
	"time"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
)

func (ctrl *Controller) postSigninPasswordlessEmailValidateRequest(
	request api.PostSigninPasswordlessEmailRequestObject,
	logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	if !ctrl.config.EmailPasswordlessEnabled {
		logger.Warn("email passwordless signin is disabled")
		return nil, &APIError{api.DisabledEndpoint}
	}

	if !ctrl.validate.ValidateEmail(string(request.Body.Email)) {
		logger.Warn("email didn't pass access control checks")
		return nil, &APIError{api.InvalidEmailPassword}
	}

	options, apiErr := ctrl.validate.SignUpOptions(
		request.Body.Options, string(request.Body.Email), logger,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	return options, nil
}

func (ctrl *Controller) PostSigninPasswordlessEmail( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninPasswordlessEmailRequestObject,
) (api.PostSigninPasswordlessEmailResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("email", string(request.Body.Email)))

	options, apiErr := ctrl.postSigninPasswordlessEmailValidateRequest(request, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	user, isMissing, apiErr := ctrl.validate.GetUserByEmail(ctx, string(request.Body.Email), logger)
	var ticket string
	switch {
	case isMissing:
		logger.Info("user does not exist, creating user")

		ticket = newTicket(TicketTypePasswordLessEmail)
		user, apiErr = ctrl.SignUpUser(
			ctx,
			string(request.Body.Email),
			options,
			logger,
			SignupUserWithTicket(ticket, time.Now().Add(time.Hour)),
		)
		if apiErr != nil {
			return ctrl.respondWithError(apiErr), nil
		}
	case apiErr != nil:
		logger.Error("error getting user by email", logError(apiErr))
		return ctrl.respondWithError(apiErr), nil
	default:
		ticket, apiErr = ctrl.SetTicket(ctx, user.ID, TicketTypePasswordLessEmail, logger)
		if apiErr != nil {
			return ctrl.respondWithError(apiErr), nil
		}
	}

	if err := ctrl.SendEmail(
		string(request.Body.Email),
		user.Locale,
		LinkTypePasswordlessEmail,
		ticket,
		deptr(options.RedirectTo),
		notifications.TemplateNameSigninPasswordless,
		user.DisplayName,
		string(request.Body.Email),
		"",
		logger,
	); err != nil {
		return nil, err
	}

	return api.PostSigninPasswordlessEmail200JSONResponse(api.OK), nil
}
