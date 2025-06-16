package controller

import (
	"context"

	"github.com/nhost/hasura-auth/go/api"
)

// func (ctrl *Controller) postSigninPasswordlessSmsOtpWithTOTP( //nolint:ireturn
// 	ctx context.Context,
// 	userID uuid.UUID,
// 	logger *slog.Logger,
// ) (api.PostSigninPasswordlessSmsOtpResponseObject, error) {
// 	ticket := "mfaTotp:" + uuid.NewString()
// 	expiresAt := time.Now().Add(In5Minutes)

// 	if apiErr := ctrl.wf.SetTicket(ctx, userID, ticket, expiresAt, logger); apiErr != nil {
// 		return ctrl.respondWithError(apiErr), nil
// 	}

// 	return api.PostSigninPasswordlessSmsOtp200JSONResponse{
// 		Mfa: &api.MFAChallengePayload{
// 			Ticket: ticket,
// 		},
// 		Session: nil,
// 	}, nil
// }

func (ctrl *Controller) PostSigninPasswordlessSmsOtp( //nolint:ireturn
	ctx context.Context,
	request api.PostSigninPasswordlessSmsOtpRequestObject,
) (api.PostSigninPasswordlessSmsOtpResponseObject, error) {
	return nil, nil
	// logger := middleware.LoggerFromContext(ctx).
	// 	With(slog.String("phoneNumber", string(request.Body.PhoneNumber)))

	// if !ctrl.config.SMSPasswordlessEnabled {
	// 	logger.Warn("SMS passwordless signin is disabled")
	// 	return ctrl.sendError(ErrDisabledEndpoint), nil
	// }

	// user, apiErr := ctrl.wf.GetUserByPhoneNumberAndOTP(
	// 	ctx, string(request.Body.PhoneNumber), request.Body.Otp, logger)

	// if apiErr != nil {
	// 	return ctrl.respondWithError(apiErr), nil
	// }

	// if user.ActiveMfaType.String == string(api.Totp) {
	// 	return ctrl.postSigninPasswordlessSmsOtpWithTOTP(ctx, user.ID, logger)
	// }

	// session, err := ctrl.wf.NewSession(ctx, user, nil, logger)
	// if err != nil {
	// 	logger.Error("error getting new session", logError(err))
	// 	return ctrl.sendError(ErrInternalServerError), nil
	// }

	// return api.PostSigninPasswordlessSmsOtp200JSONResponse{
	// 	Session: session,
	// 	Mfa:     nil,
	// }, nil
}
