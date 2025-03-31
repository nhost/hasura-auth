package controller

import (
	"context"
	"fmt"
	"net/url"

	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/middleware"
	"github.com/nhost/hasura-auth/go/oauth2"
)

func (ctrl *Controller) getSigninProviderProviderCallbackRedirectError(
	errorMsg string,
) (api.GetSigninProviderProviderCallbackResponseObject, error) {
	return api.GetSigninProviderProviderCallback302Response{
		Headers: api.GetSigninProviderProviderCallback302ResponseHeaders{
			Location: fmt.Sprintf("http://localhost:3000/auth/error?error=%s", url.QueryEscape(errorMsg)),
		},
	}, nil
}

func (ctrl *Controller) GetSigninProviderProviderCallback( //nolint:ireturn
	ctx context.Context,
	req api.GetSigninProviderProviderCallbackRequestObject,
) (api.GetSigninProviderProviderCallbackResponseObject, error) {
	logger := middleware.LoggerFromContext(ctx)

	provider := ctrl.Providers.Get(string(req.Provider))
	if provider == nil {
		return nil, fmt.Errorf("provider not found: %s", req.Provider)
	}

	if req.Params.Error != nil && *req.Params.Error != "" {
		errorMsg := *req.Params.Error
		if req.Params.ErrorDescription != nil {
			errorMsg += ": " + *req.Params.ErrorDescription
		}
		return ctrl.getSigninProviderProviderCallbackRedirectError(errorMsg)
	}

	signInData := &oauth2.ProviderSignInData{} //nolint:exhaustruct
	if err := signInData.Decode(req.Params.NhostAuthProviderSignInData); err != nil {
		logger.Error("failed to decode sign-in data", logError(err))
		return ctrl.getSigninProviderProviderCallbackRedirectError("invalid_state")
	}

	if req.Params.State != signInData.State {
		logger.Error("state mismatch", "expected", signInData.State, "actual", req.Params.State)
		return ctrl.getSigninProviderProviderCallbackRedirectError("invalid_state")
	}

	token, err := provider.Oauth2().Exchange(ctx, req.Params.Code)
	if err != nil {
		logger.Error("failed to exchange token", logError(err))
		return ctrl.getSigninProviderProviderCallbackRedirectError("token_exchange_failed")
	}

	profile, err := provider.GetProfile(ctx, token.AccessToken)
	if err != nil {
		logger.Error("failed to get user info", logError(err))
		return ctrl.getSigninProviderProviderCallbackRedirectError("failed_to_get_user_info")
	}

	fmt.Println("githubUser", profile)

	// // Clear the state cookie
	// clearCookie := &http.Cookie{
	// 	Name:     "nhostAuthSigninGithubData",
	// 	Value:    "",
	// 	MaxAge:   -1,
	// 	Secure:   false, // Change to true in production
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteLaxMode,
	// 	Path:     "/signin/provider/" + string(req.Provider),
	// }

	return api.GetSigninProviderProviderCallback302Response{
		Headers: api.GetSigninProviderProviderCallback302ResponseHeaders{
			Location: "http://localhost:3000",
		},
	}, nil
}
