package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

func (a *Auth) PostSigninEmailPassword(
	ctx context.Context,
	request api.PostSigninEmailPasswordRequestObject,
) (api.PostSigninEmailPasswordResponseObject, error) {
	return api.PostSigninEmailPassword200JSONResponse{}, nil
}
