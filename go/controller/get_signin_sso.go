package controller

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/zitadel/oidc/v3/pkg/client/rp"
)

func (ctrl *Controller) getSigninSsoProviderValidateRequest(
	request api.GetSigninSsoProviderRequestObject,
	logger *slog.Logger,
) (*url.URL, *api.SignUpOptions, *APIError) {
	redirectToURL := ctrl.config.ClientURL
	redirectTo := ptr(ctrl.config.ClientURL.String())
	if request.Params.RedirectTo != nil {
		if !ctrl.wf.redirectURLValidator(deptr(request.Params.RedirectTo)) {
			logger.Warn("redirect URL not allowed",
				slog.String("redirectTo", deptr(request.Params.RedirectTo)))
			return nil, nil, ErrRedirecToNotAllowed
		}
	}

	return redirectToURL, &api.SignUpOptions{
		AllowedRoles: nil,
		DefaultRole:  nil,
		DisplayName:  nil,
		Locale:       nil,
		Metadata:     nil,
		RedirectTo:   redirectTo,
	}, nil
}

func getSigninProviderRedirectError(
	redirectUrl *url.URL, error api.ErrorResponseError,
) api.GetSigninSsoProvider302Response {
	q := redirectUrl.Query()
	q.Set("error", string(error))
	redirectUrl.RawQuery = q.Encode()

	return api.GetSigninSsoProvider302Response{
		Headers: api.GetSigninSsoProvider302ResponseHeaders{
			Location:  redirectUrl.String(),
			SetCookie: "",
		},
	}
}

type OIDCFlowData struct {
	State   string             `json:"state"`
	Options *api.SignUpOptions `json:"options"`
}

func oidcFlowDataToCookie(state string, options *api.SignUpOptions) (string, error) {
	s := OIDCFlowData{
		State:   state,
		Options: options,
	}

	b, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to marshal OIDC flow data: %w", err)
	}

	// encode in base64
	encoded := base64.StdEncoding.EncodeToString(b)

	return fmt.Sprintf("nhost_auth_oidc_flow_data=%s; Path=/; HttpOnly; SameSite=Lax;", encoded), nil
}

func oidcFlowDataFromCookie(cookie string) (OIDCFlowData, error) {
	decoded, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return OIDCFlowData{}, fmt.Errorf("failed to decode OIDC flow data: %w", err)
	}

	var s OIDCFlowData
	if err := json.Unmarshal(decoded, &s); err != nil {
		return s, fmt.Errorf("failed to unmarshal OIDC flow data: %w", err)
	}

	return s, nil
}

func (c *Controller) GetSigninSsoProvider(
	ctx context.Context,
	request api.GetSigninSsoProviderRequestObject,
) (api.GetSigninSsoProviderResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).With("provider", request.Provider)

	clientURL, options, apiErr := c.getSigninSsoProviderValidateRequest(request, logger)
	if apiErr != nil {
		return getSigninProviderRedirectError(clientURL, apiErr.t), nil
	}

	state := "some-state"

	provider, ok := c.oidcProviders[request.Provider]
	if !ok {
		logger.Warn("provider not found")
		return getSigninProviderRedirectError(clientURL, api.DisabledEndpoint), nil
	}

	oidcFlowDataCookie, err := oidcFlowDataToCookie(state, options)
	if err != nil {
		logger.Error("failed to create OIDC flow data cookie", logError(err))
		return getSigninProviderRedirectError(clientURL, api.InternalServerError), nil
	}

	return api.GetSigninSsoProvider302Response{
		Headers: api.GetSigninSsoProvider302ResponseHeaders{
			Location:  rp.AuthURL(state, provider),
			SetCookie: oidcFlowDataCookie,
		},
	}, nil
}
