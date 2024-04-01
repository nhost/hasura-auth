package controller

import (
	"context"

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
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	session, err := ctrl.wf.UpdateSession(ctx, user, request.Body.RefreshToken, logger)
	if err != nil {
		logger.Error("error getting new session", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	return api.PostToken200JSONResponse{
		Session: session,
	}, nil
}