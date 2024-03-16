package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) PostSigninPat( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninPatRequestObject,
) (api.PostSigninPatResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	user, apiErr := ctrl.validate.GetUserByRefreshTokenHash(
		ctx,
		request.Body.PersonalAccessToken,
		sql.RefreshTokenTypePAT,
		logger,
	)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	session, err := ctrl.GetNewSession(ctx, user, logger)
	if err != nil {
		logger.Error("error getting new session", logError(err))
		return ctrl.sendError(api.InternalServerError), nil
	}

	return api.PostSigninPat200JSONResponse{
		Session: session,
	}, nil
}
