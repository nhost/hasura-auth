package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/zitadel/oidc/v3/pkg/client"
	"github.com/zitadel/oidc/v3/pkg/client/rp"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func (ctrl *Controller) getSigninSsoProviderCallbackValidateRequest(
	request api.GetSigninSsoProviderCallbackRequestObject,
	logger *slog.Logger,
) (*url.URL, *api.SignUpOptions, *APIError) {
	clientURL := ctrl.config.ClientURL

	SsoFlowData, err := oidcFlowDataFromCookie(request.Params.NhostAuthSsoFlowData)
	if err != nil {
		logger.Warn("failed to parse oidc flow data", logError(err))
		return nil, nil, ErrInvalidRequest
	}

	if SsoFlowData.State != request.Params.State {
		logger.Warn("state mismatch", slog.String("state", request.Params.State))
		return nil, nil, ErrInvalidRequest
	}

	options, apiErr := ctrl.wf.ValidateSignUpOptions(
		SsoFlowData.Options, "", logger,
	)
	if apiErr != nil {
		return clientURL, nil, apiErr
	}

	return clientURL, options, nil
}

func getSigninProviderCallbackRedirectError(
	redirectUrl *url.URL, error api.ErrorResponseError,
) api.GetSigninSsoProviderCallback302Response {
	q := redirectUrl.Query()
	q.Set("error", string(error))
	redirectUrl.RawQuery = q.Encode()

	return api.GetSigninSsoProviderCallback302Response{
		Headers: api.GetSigninSsoProviderCallback302ResponseHeaders{
			Location:  redirectUrl.String(),
			SetCookie: "",
		},
	}
}

func (c *Controller) GetSigninSsoProviderCallback(
	ctx context.Context,
	request api.GetSigninSsoProviderCallbackRequestObject,
) (api.GetSigninSsoProviderCallbackResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).With(slog.String("provider", request.Provider))

	clientURL, _, apiErr := c.getSigninSsoProviderCallbackValidateRequest(request, logger)
	if apiErr != nil {
		return getSigninProviderCallbackRedirectError(clientURL, apiErr.t), nil
	}

	provider, ok := c.oidcProviders[request.Provider]
	if !ok {
		panic("provider not found: " + request.Provider)
	}

	codeOpts := make([]rp.CodeExchangeOpt, 0)
	if provider.Signer() != nil {
		assertion, err := client.SignedJWTProfileAssertion(
			provider.OAuthConfig().ClientID,
			[]string{provider.Issuer()},
			time.Hour,
			provider.Signer(),
		)
		if err != nil {
			panic("failed to build assertion: " + err.Error())
		}
		codeOpts = append(codeOpts, rp.WithClientAssertionJWT(assertion))
	}

	tokens, err := rp.CodeExchange[*oidc.IDTokenClaims](ctx, request.Params.Code, provider, codeOpts...)
	if err != nil {
		panic("failed to exchange token: " + err.Error())
	}

	info, err := rp.Userinfo[*oidc.UserInfo](
		ctx,
		tokens.AccessToken,
		tokens.TokenType,
		tokens.IDTokenClaims.GetSubject(),
		provider,
	)
	if err != nil {
		panic("userinfo failed: " + err.Error())
	}

	b, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(b))

	return api.GetSigninSsoProviderCallback302Response{
		Headers: api.GetSigninSsoProviderCallback302ResponseHeaders{
			Location:  "http://localhost:4000/version",
			SetCookie: "sso=; Max-Age=0",
		},
	}, nil
}
