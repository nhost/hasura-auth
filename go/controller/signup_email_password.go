package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

func (a *Auth) PostSignupEmailPassword(
	ctx context.Context,
	request api.PostSignupEmailPasswordRequestObject,
) (api.PostSignupEmailPasswordResponseObject, error) {
	return api.PostSignupEmailPassword200JSONResponse{}, nil
}
