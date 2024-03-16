package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) PostPat( //nolint:ireturn
	ctx context.Context, request api.PostPatRequestObject,
) (api.PostPatResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)
	user, apiErr := ctrl.GetUserFromJWTInContext(ctx, logger)
	if apiErr != nil {
		return ctrl.respondWithError(apiErr), nil
	}

	var metadata []byte
	var err error
	if request.Body.Metadata != nil {
		metadata, err = json.Marshal(*request.Body.Metadata)
		if err != nil {
			logger.Error("error marshalling metadata", logError(err))
			return ctrl.sendError(api.InternalServerError), nil
		}
	}

	pat := uuid.New()
	refreshToken, err := ctrl.db.InsertRefreshtoken(ctx, sql.InsertRefreshtokenParams{
		UserID:           user.ID,
		RefreshTokenHash: sql.Text(hashRefreshToken([]byte(pat.String()))),
		ExpiresAt:        sql.TimestampTz(request.Body.ExpiresAt),
		Type:             sql.RefreshTokenTypePAT,
		Metadata:         metadata,
	})
	if err != nil {
		return nil, fmt.Errorf("error inserting refresh token: %w", err)
	}

	return api.PostPat200JSONResponse{
		Id:                  refreshToken.String(),
		PersonalAccessToken: pat.String(),
	}, nil
}
