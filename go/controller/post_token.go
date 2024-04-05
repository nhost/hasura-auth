package controller

import (
	"context"
	"errors"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) PostToken( //nolint:ireturn
	ctx context.Context, request api.PostTokenRequestObject,
) (api.PostTokenResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	user, apiErr := ctrl.wf.GetUserByRefreshTokenHash(
		ctx,
		request.Body.RefreshToken,
		sql.RefreshTokenTypeRegular,
		logger,
	)

	switch {
	case errors.Is(apiErr, ErrForbiddenAnonymous):
	case errors.Is(apiErr, ErrInvalidEmailPassword):
	default:
		if apiErr != nil {
			return ctrl.respondWithError(apiErr), nil
		}
	}

	session, err := ctrl.wf.UpdateSession(ctx, user, request.Body.RefreshToken, logger)
	if err != nil {
		logger.Error("error updating session", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	return api.PostToken200JSONResponse(*session), nil
}
