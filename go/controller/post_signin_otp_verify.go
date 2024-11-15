package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

func (ctrl *Controller) PostSigninOtpVerify( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninOtpVerifyRequestObject,
) (api.PostSigninOtpVerifyResponseObject, error) {
	return nil, nil
}
