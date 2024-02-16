package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nhost/hasura-auth/go/api"
)

func isSensitive(err api.ErrorResponseError) bool {
	switch err {
	case api.SignupDisabled, api.EmailAlreadyInUse:
		return true
	case api.InvalidRequest, api.InternalServerError:
		return false
	}
	return false
}

type ErrorResponse api.ErrorResponse

func (response ErrorResponse) visit(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	return json.NewEncoder(w).Encode(response) //nolint:wrapcheck
}

func (response ErrorResponse) VisitPostSignupEmailPasswordResponse(w http.ResponseWriter) error {
	return response.visit(w)
}

func (ctrl *Controller) sendError(errType api.ErrorResponseError) ErrorResponse {
	invalidRequest := ErrorResponse{
		Code:    http.StatusBadRequest,
		Error:   api.InvalidRequest,
		Message: "invalid-request",
	}

	if ctrl.config.ConcealErrors && isSensitive(errType) {
		return invalidRequest
	}

	switch errType {
	case api.SignupDisabled:
		return ErrorResponse{
			Code:    http.StatusForbidden,
			Error:   errType,
			Message: "Signup is disabled",
		}
	case api.EmailAlreadyInUse:
		return ErrorResponse{
			Code:    http.StatusConflict,
			Error:   errType,
			Message: "Email already in use",
		}
	case api.InvalidRequest, api.InternalServerError:
	}

	return invalidRequest
}
