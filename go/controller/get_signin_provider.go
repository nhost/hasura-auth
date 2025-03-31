package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/oauth2"
)

const oauth2StateLength = 16

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", fmt.Errorf("error reading random bytes: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (ctrl *Controller) getSigninProviderValidateRequest(
	req api.GetSigninProviderProviderRequestObject,
	logger *slog.Logger,
) (*url.URL, *APIError) {
	options, apiErr := ctrl.wf.ValidateOptionsRedirectTo(
		&api.OptionsRedirectTo{
			RedirectTo: req.Params.RedirectTo,
		},
		logger,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	redirectTo, err := url.Parse(*options.RedirectTo)
	if err != nil {
		logger.Error("error parsing redirect URL",
			slog.String("redirectTo", *options.RedirectTo), logError(err))
		return nil, ErrInvalidRequest
	}

	return redirectTo, nil
}

func (ctrl *Controller) GetSigninProviderProvider( //nolint:ireturn
	ctx context.Context,
	req api.GetSigninProviderProviderRequestObject,
) (api.GetSigninProviderProviderResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx).
		With(slog.String("provider", string(req.Provider)))

	redirectTo, apiErr := ctrl.getSigninProviderValidateRequest(req, logger)
	if apiErr != nil {
		return ctrl.sendError(apiErr), nil
	}

	state, err := randString(oauth2StateLength)
	if err != nil {
		logger.Error("error generating random string", logError(err))
		return ctrl.sendRedirectError(redirectTo, ErrInternalServerError), nil
	}

	provider := ctrl.Providers.Get(string(req.Provider))
	if provider == nil {
		logger.Error("provider not enabled")
		return ctrl.sendRedirectError(redirectTo, ErrDisabledEndpoint), nil
	}

	url := provider.Oauth2().AuthCodeURL(
		state,
	)

	data := oauth2.ProviderSignInData{
		State: state,
		Options: api.SignUpOptions{
			AllowedRoles: req.Params.AllowedRoles,
			DefaultRole:  req.Params.DefaultRole,
			DisplayName:  req.Params.DisplayName,
			Locale:       req.Params.Locale,
			Metadata:     req.Params.Metadata,
			RedirectTo:   req.Params.RedirectTo,
		},
	}
	dataEncoded, err := data.Encode()
	if err != nil {
		logger.Error("error encoding provider sign in data", logError(err))
		return ctrl.sendRedirectError(redirectTo, ErrInternalServerError), nil
	}

	cookie := &http.Cookie{ //nolint:exhaustruct
		Name:     "nhostAuthProviderSignInData",
		Value:    dataEncoded,
		MaxAge:   int(time.Minute.Seconds()),
		Secure:   ctrl.config.UseSecureCookies(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/signin/provider/" + string(req.Provider),
	}

	return api.GetSigninProviderProvider302Response{
		Headers: api.GetSigninProviderProvider302ResponseHeaders{
			Location:  url,
			SetCookie: cookie.String(),
		},
	}, nil
}
