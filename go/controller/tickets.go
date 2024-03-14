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
)

func newTicket(ticketType TicketType) string {
	return fmt.Sprintf("%s:%s", ticketType, uuid.NewString())
}

func (ctrl *Controller) setTicket(
	ctx context.Context,
	userID uuid.UUID,
	ticketType TicketType,
	logger *slog.Logger,
) (sql.AuthUser, error) {
	ticket := newTicket(ticketType)
	ticketExpiresAt := time.Now().Add(time.Hour)

	user, err := ctrl.db.UpdateUserTicket(
		ctx,
		sql.UpdateUserTicketParams{
			ID:              userID,
			Ticket:          sql.Text(ticket),
			TicketExpiresAt: sql.TimestampTz(ticketExpiresAt),
		},
	)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Error("user not found")
		return sql.AuthUser{}, //nolint:exhaustruct
			&APIError{api.InvalidRequest}
	}
	if err != nil {
		logger.Error("error updating user ticket", logError(err))
		return sql.AuthUser{}, //nolint:exhaustruct
			&APIError{api.InternalServerError}
	}

	return user, nil
}
