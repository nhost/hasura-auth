package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

func (ctrl *Controller) GetHealthz( //nolint:ireturn
	_ context.Context, _ api.GetHealthzRequestObject,
) (api.GetHealthzResponseObject, error) {
	return api.GetHealthz200JSONResponse(api.OK), nil
}
