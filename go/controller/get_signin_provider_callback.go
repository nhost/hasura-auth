package controller

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/oauth2"
	"github.com/nhost/hasura-auth/go/oidc"
)

func (ctrl *Controller) getSigninProviderProviderCallbackValidate( //nolint:ireturn
	req api.GetSigninProviderProviderCallbackRequestObject,
	clearCookie *http.Cookie,
	logger *slog.Logger,
) (api.GetSigninProviderProviderCallbackResponseObject, *api.SignUpOptions, *string, *url.URL, *APIError) {
	redirectTo := ctrl.config.ClientURL

	cookieData := &oauth2.ProviderSignInData{} //nolint:exhaustruct
	if err := cookieData.Decode(req.Params.NhostAuthProviderSignInData); err != nil {
		logger.Error("failed to decode cookie", logError(err))
		return nil, nil, nil, redirectTo, ErrInvalidState
	}

	// we just care about the redirect URL for now, the rest is handled by the signin flow
	options, apiErr := ctrl.wf.ValidateOptionsRedirectTo(
		&api.OptionsRedirectTo{
			RedirectTo: cookieData.Options.RedirectTo,
		},
		logger,
	)
	if apiErr != nil {
		return nil, nil, nil, redirectTo, apiErr
	}

	if req.Params.Error != nil && *req.Params.Error != "" {
		values := redirectTo.Query()
		values.Add("error", deptr(req.Params.Error))
		values.Add("error_description", deptr(req.Params.ErrorDescription))
		values.Add("error_url", deptr(req.Params.ErrorUri))
		redirectTo.RawQuery = values.Encode()

		return api.GetSigninProviderProviderCallback302Response{
			Headers: api.GetSigninProviderProviderCallback302ResponseHeaders{
				Location:  redirectTo.String(),
				SetCookie: clearCookie.String(),
			},
		}, nil, nil, nil, nil
	}

	if req.Params.State != cookieData.State {
		logger.Error("state mismatch", "expected", cookieData.State, "actual", req.Params.State)
		return nil, nil, nil, redirectTo, ErrInvalidState
	}

	optionsRedirectTo, err := url.Parse(*options.RedirectTo)
	if err != nil {
		logger.Error("error parsing redirect URL", logError(err))
		return nil, nil, nil, redirectTo, ErrInvalidRequest
	}

	return nil, &cookieData.Options, cookieData.Connect, optionsRedirectTo, nil
}

func (ctrl *Controller) getSigninProviderProviderCallbackOauthFlow(
	ctx context.Context,
	req api.GetSigninProviderProviderCallbackRequestObject,
	logger *slog.Logger,
) (oidc.Profile, *APIError) {
	provider := ctrl.Providers.Get(string(req.Provider))
	if provider == nil {
		logger.Error("provider not enabled")
		return oidc.Profile{}, ErrDisabledEndpoint
	}

	token, err := provider.Exchange(ctx, req.Params.Code)
	if err != nil {
		logger.Error("failed to exchange token", logError(err))
		return oidc.Profile{}, ErrOauthTokenExchangeFailed
	}

	profile, err := provider.GetProfile(ctx, token.AccessToken)
	if err != nil {
		logger.Error("failed to get user info", logError(err))
		return oidc.Profile{}, ErrOauthProfileFetchFailed
	}

	return profile, nil
}

func (ctrl *Controller) GetSigninProviderProviderCallback( //nolint:ireturn
	ctx context.Context,
	req api.GetSigninProviderProviderCallbackRequestObject,
) (api.GetSigninProviderProviderCallbackResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	// Clear the state cookie
	clearCookie := &http.Cookie{ //nolint:exhaustruct
		Name:     signinProviderCookieName,
		Value:    "",
		MaxAge:   0,
		Secure:   ctrl.config.UseSecureCookies(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/signin/provider/" + string(req.Provider),
	}

	quickReturn, options, connnect, redirectTo, apiErr := ctrl.getSigninProviderProviderCallbackValidate(
		req, clearCookie, logger,
	)
	if apiErr != nil {
		return ctrl.sendRedirectError(redirectTo, apiErr), nil
	}
	if quickReturn != nil {
		// we use quick return to forward the error message from the oauth provider
		return quickReturn, nil
	}

	profile, apiErr := ctrl.getSigninProviderProviderCallbackOauthFlow(ctx, req, logger)
	if apiErr != nil {
		return ctrl.sendRedirectError(redirectTo, apiErr), nil
	}

	if connnect != nil {
		// TODO
	}

	session, apiErr := ctrl.providerSignInFlow(
		ctx, profile, string(req.Provider), options, logger,
	)
	if apiErr != nil {
		return ctrl.sendRedirectError(redirectTo, apiErr), nil
	}

	if session != nil {
		values := redirectTo.Query()
		values.Add("refreshToken", session.RefreshToken)
		redirectTo.RawQuery = values.Encode()
	}

	return api.GetSigninProviderProviderCallback302Response{
		Headers: api.GetSigninProviderProviderCallback302ResponseHeaders{
			Location:  redirectTo.String(),
			SetCookie: clearCookie.String(),
		},
	}, nil
}
