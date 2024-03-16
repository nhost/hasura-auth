package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/sql"
)

type TicketType string

const (
	TicketTypeEmailConfirmChange TicketType = "emailConfirmChange"
	TicketTypePasswordLessEmail  TicketType = "passwordlessEmail"
	TicketTypeVerifyEmail        TicketType = "verifyEmail"
	TicketTypePasswordReset      TicketType = "passwordReset"
)

func newTicket(ticketType TicketType) string {
	return fmt.Sprintf("%s:%s", ticketType, uuid.NewString())
}

func (ctrl *Controller) SetTicket(
	ctx context.Context,
	userID uuid.UUID,
	ticket string,
	expiresAt time.Time,
	logger *slog.Logger,
) *APIError {
	_, err := ctrl.validate.db.UpdateUserTicket(
		ctx,
		sql.UpdateUserTicketParams{
			ID:              userID,
			Ticket:          sql.Text(ticket),
			TicketExpiresAt: sql.TimestampTz(expiresAt),
		},
	)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Error("user not found")
		return &APIError{api.InvalidRequest}
	}
	if err != nil {
		logger.Error("error updating user ticket", logError(err))
		return &APIError{api.InternalServerError}
	}

	return nil
}
