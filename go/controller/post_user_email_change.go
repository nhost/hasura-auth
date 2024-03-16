package controller

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) PostUserEmailChange( //nolint:ireturn
	ctx context.Context, request api.PostUserEmailChangeRequestObject,
) (api.PostUserEmailChangeResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	options, apiErr := ctrl.validate.OptionsRedirectTo(request.Body.Options, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}
	request.Body.Options = options

	user, apiErr := ctrl.GetUserFromJWTInContext(ctx, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	_, err := ctrl.db.GetUserByEmail(ctx, sql.Text(string(request.Body.NewEmail)))
	if err == nil {
		logger.Warn("email already exists")
		return ctrl.sendError(api.EmailAlreadyInUse), nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		logger.Error("error getting user by email", logError(err))
		return ctrl.respondWithError(err), nil
	}

	ticket := newTicket(TicketTypeEmailConfirmChange)
	ticketExpiresAt := time.Now().Add(time.Hour)

	updatedUser, err := ctrl.db.UpdateUserChangeEmail(
		ctx,
		sql.UpdateUserChangeEmailParams{
			ID:              user.ID,
			Ticket:          sql.Text(ticket),
			TicketExpiresAt: sql.TimestampTz(ticketExpiresAt),
			NewEmail:        sql.Text(string(request.Body.NewEmail)),
		},
	)
	if err != nil {
		logger.Error("error updating user ticket", logError(err))
		return ctrl.respondWithError(err), nil
	}

	if err := ctrl.SendEmail(
		string(request.Body.NewEmail),
		updatedUser.Locale,
		LinkTypeEmailConfirmChange,
		ticket,
		deptr(request.Body.Options.RedirectTo),
		notifications.TemplateNameEmailConfirmChange,
		updatedUser.DisplayName,
		updatedUser.Email.String,
		string(request.Body.NewEmail),
		logger,
	); err != nil {
		return nil, err
	}

	return api.PostUserEmailChange200JSONResponse(api.OK), nil
}
