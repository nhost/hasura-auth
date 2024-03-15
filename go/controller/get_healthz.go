package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

func (ctrl *Controller) GetHealth( //nolint:ireturn
	_ context.Context, _ api.GetHealthRequestObject,
) (api.GetHealthResponseObject, error) {
	return api.GetHealth200JSONResponse(api.OK), nil
}
