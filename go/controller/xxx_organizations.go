package controller

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/sql"
)

func (ctrl *Controller) GetOrganizations( //nolint:ireturn
	ctx context.Context,
	_ api.GetOrganizationsRequestObject,
) (api.GetOrganizationsResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	jwtToken, ok := ctrl.wf.jwtGetter.FromContext(ctx)
	if !ok {
		logger.Error("error getting jwt token from context")
		return ctrl.sendError(ErrInternalServerError), nil
	}

	userID, err := ctrl.wf.jwtGetter.GetUserID(jwtToken)
	if err != nil {
		logger.Error("error getting user id from jwt token", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	orgs, err := ctrl.wf.db.GetOrganizations(ctx, userID)
	if err != nil {
		logger.Error("error getting organizations", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	response := make(api.GetOrganizations200JSONResponse, len(orgs))
	for i, org := range orgs {
		response[i] = api.Organization{
			Id:   org.ID.String(),
			Name: org.Name,
		}
	}

	return response, nil
}

func (ctrl *Controller) PostOrganizations( //nolint:ireturn
	ctx context.Context,
	request api.PostOrganizationsRequestObject,
) (api.PostOrganizationsResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	jwtToken, ok := ctrl.wf.jwtGetter.FromContext(ctx)
	if !ok {
		logger.Error("error getting jwt token from context")
		return ctrl.sendError(ErrInternalServerError), nil
	}

	userID, err := ctrl.wf.jwtGetter.GetUserID(jwtToken)
	if err != nil {
		logger.Error("error getting user id from jwt token", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	org, err := ctrl.wf.db.InsertOrganization(ctx, sql.InsertOrganizationParams{
		Name:   request.Body.Name,
		UserID: userID,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return ctrl.sendError(ErrOnlyOneOrganizationPerUserAllowed), nil
	}

	if err != nil {
		logger.Error("error inserting organization", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	return api.PostOrganizations201JSONResponse{
		Id:   org.OrganizationID.String(),
		Name: org.Name,
	}, nil
}

func (ctrl *Controller) DeleteOrganizationsId( //nolint:ireturn,revive,stylecheck
	ctx context.Context,
	request api.DeleteOrganizationsIdRequestObject,
) (api.DeleteOrganizationsIdResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	jwtToken, ok := ctrl.wf.jwtGetter.FromContext(ctx)
	if !ok {
		logger.Error("error getting jwt token from context")
		return ctrl.sendError(ErrInternalServerError), nil
	}

	userID, err := ctrl.wf.jwtGetter.GetUserID(jwtToken)
	if err != nil {
		logger.Error("error getting user id from jwt token", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	_, err = ctrl.wf.db.DeleteOrganization(ctx, sql.DeleteOrganizationParams{
		OrganizationID: request.Id,
		UserID:         userID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return ctrl.sendError(ErrOrganizationNotFoundOrUserNotAdmin), nil
	}
	if err != nil {
		logger.Error("error deleting organization", logError(err))
		return ctrl.sendError(ErrInternalServerError), nil
	}

	return api.DeleteOrganizationsId200JSONResponse(api.OK), nil
}
