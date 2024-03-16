package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/oapi-codegen/runtime/types"
)

func (ctrl *Controller) GetNewSession(
	ctx context.Context,
	user sql.AuthUser,
	logger *slog.Logger,
) (*api.Session, error) {
	userRoles, err := ctrl.db.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting roles by user id: %w", err)
	}
	allowedRoles := make([]string, len(userRoles))
	for i, role := range userRoles {
		allowedRoles[i] = role.Role
	}

	refreshToken := uuid.New()
	expiresAt := time.Now().Add(time.Duration(ctrl.config.RefreshTokenExpiresIn) * time.Second)
	if _, err := ctrl.db.InsertRefreshtoken(ctx, sql.InsertRefreshtokenParams{
		UserID:           user.ID,
		RefreshTokenHash: sql.Text(hashRefreshToken([]byte(refreshToken.String()))),
		ExpiresAt:        sql.TimestampTz(expiresAt),
		Type:             sql.RefreshTokenTypeRegular,
		Metadata:         nil,
	}); err != nil {
		return nil, fmt.Errorf("error inserting refresh token: %w", err)
	}

	if _, err := ctrl.db.UpdateUserLastSeen(ctx, user.ID); err != nil {
		return nil, fmt.Errorf("error updating last seen: %w", err)
	}

	accessToken, expiresIn, err := ctrl.jwtGetter.GetToken(
		ctx, user.ID, allowedRoles, user.DefaultRole, logger,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting jwt: %w", err)
	}

	var metadata map[string]any
	if err := json.Unmarshal(user.Metadata, &metadata); err != nil {
		return nil, fmt.Errorf("error unmarshalling user metadata: %w", err)
	}
	return &api.Session{
		AccessToken:          accessToken,
		AccessTokenExpiresIn: expiresIn,
		RefreshToken:         refreshToken.String(),
		User: &api.User{
			AvatarUrl:           user.AvatarUrl,
			CreatedAt:           user.CreatedAt.Time,
			DefaultRole:         user.DefaultRole,
			DisplayName:         user.DisplayName,
			Email:               types.Email(user.Email.String),
			EmailVerified:       user.EmailVerified,
			Id:                  user.ID.String(),
			IsAnonymous:         false,
			Locale:              user.Locale,
			Metadata:            metadata,
			PhoneNumber:         user.PhoneNumber.String,
			PhoneNumberVerified: user.PhoneNumberVerified,
			Roles:               allowedRoles,
		},
	}, nil
}

func (ctrl *Controller) GetUserFromJWTInContext(
	ctx context.Context,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	jwtToken, ok := ctrl.jwtGetter.FromContext(ctx)
	if !ok {
		logger.Error(
			"jwt token not found in context, this should not be possilble due to middleware",
		)
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	sub, err := jwtToken.Claims.GetSubject()
	if err != nil {
		logger.Error("error getting user id from jwt token", logError(err))
		return sql.AuthUser{}, &APIError{api.InvalidRequest} //nolint:exhaustruct
	}
	logger = logger.With(slog.String("user_id", sub))

	userID, err := uuid.Parse(sub)
	if err != nil {
		logger.Error("error parsing user id from jwt token's subject", logError(err))
		return sql.AuthUser{}, &APIError{api.InvalidRequest} //nolint:exhaustruct
	}

	user, apiErr := ctrl.validate.GetUser(ctx, userID, logger)
	if apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	if apiErr := ctrl.validate.User(user, logger); apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	return user, nil
}
