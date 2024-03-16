// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version (devel) DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for ErrorResponseError.
const (
	DefaultRoleMustBeInAllowedRoles ErrorResponseError = "default-role-must-be-in-allowed-roles"
	DisabledEndpoint                ErrorResponseError = "disabled-endpoint"
	DisabledUser                    ErrorResponseError = "disabled-user"
	EmailAlreadyInUse               ErrorResponseError = "email-already-in-use"
	ForbiddenAnonymous              ErrorResponseError = "forbidden-anonymous"
	InternalServerError             ErrorResponseError = "internal-server-error"
	InvalidEmailPassword            ErrorResponseError = "invalid-email-password"
	InvalidPat                      ErrorResponseError = "invalid-pat"
	InvalidRequest                  ErrorResponseError = "invalid-request"
	LocaleNotAllowed                ErrorResponseError = "locale-not-allowed"
	PasswordInHibpDatabase          ErrorResponseError = "password-in-hibp-database"
	PasswordTooShort                ErrorResponseError = "password-too-short"
	RedirecToNotAllowed             ErrorResponseError = "redirecTo-not-allowed"
	RoleNotAllowed                  ErrorResponseError = "role-not-allowed"
	SignupDisabled                  ErrorResponseError = "signup-disabled"
	UnverifiedUser                  ErrorResponseError = "unverified-user"
)

// Defines values for OKResponse.
const (
	OK OKResponse = "OK"
)

// CreatePATRequest defines model for CreatePATRequest.
type CreatePATRequest struct {
	// ExpiresAt Expiration date of the PAT
	ExpiresAt time.Time               `json:"expiresAt"`
	Metadata  *map[string]interface{} `json:"metadata,omitempty"`
}

// CreatePATResponse defines model for CreatePATResponse.
type CreatePATResponse struct {
	// Id ID of the PAT
	Id string `json:"id"`

	// PersonalAccessToken PAT
	PersonalAccessToken string `json:"personalAccessToken"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error code that identifies the application error
	Error ErrorResponseError `json:"error"`

	// Message Human friendly error message
	Message string `json:"message"`

	// Status HTTP status error code
	Status int `json:"status"`
}

// ErrorResponseError Error code that identifies the application error
type ErrorResponseError string

// MFAChallengePayload defines model for MFAChallengePayload.
type MFAChallengePayload struct {
	Ticket string `json:"ticket"`
}

// OKResponse defines model for OKResponse.
type OKResponse string

// OptionsRedirectTo defines model for OptionsRedirectTo.
type OptionsRedirectTo struct {
	RedirectTo *string `json:"redirectTo,omitempty"`
}

// Session defines model for Session.
type Session struct {
	AccessToken          string `json:"accessToken"`
	AccessTokenExpiresIn int64  `json:"accessTokenExpiresIn"`

	// RefreshToken Refresh token during authentication or when refreshing the JWT
	RefreshToken string `json:"refreshToken"`
	User         *User  `json:"user,omitempty"`
}

// SessionPayload defines model for SessionPayload.
type SessionPayload struct {
	Session *Session `json:"session,omitempty"`
}

// SignInEmailPasswordRequest defines model for SignInEmailPasswordRequest.
type SignInEmailPasswordRequest struct {
	// Email A valid email
	Email openapi_types.Email `json:"email"`

	// Password A password of minimum 3 characters
	Password string `json:"password"`
}

// SignInEmailPasswordResponse defines model for SignInEmailPasswordResponse.
type SignInEmailPasswordResponse struct {
	Mfa     *MFAChallengePayload `json:"mfa,omitempty"`
	Session *Session             `json:"session,omitempty"`
}

// SignInPATRequest defines model for SignInPATRequest.
type SignInPATRequest struct {
	// PersonalAccessToken PAT
	PersonalAccessToken string `json:"personalAccessToken"`
}

// SignInPasswordlessEmailRequest defines model for SignInPasswordlessEmailRequest.
type SignInPasswordlessEmailRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *SignUpOptions      `json:"options,omitempty"`
}

// SignUpEmailPasswordRequest defines model for SignUpEmailPasswordRequest.
type SignUpEmailPasswordRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *SignUpOptions      `json:"options,omitempty"`

	// Password A password of minimum 3 characters
	Password string `json:"password"`
}

// SignUpOptions defines model for SignUpOptions.
type SignUpOptions struct {
	AllowedRoles *[]string `json:"allowedRoles,omitempty"`
	DefaultRole  *string   `json:"defaultRole,omitempty"`
	DisplayName  *string   `json:"displayName,omitempty"`

	// Locale A two-characters locale
	Locale     *string                 `json:"locale,omitempty"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	RedirectTo *string                 `json:"redirectTo,omitempty"`
}

// User defines model for User.
type User struct {
	AvatarUrl   string    `json:"avatarUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	DefaultRole string    `json:"defaultRole"`
	DisplayName string    `json:"displayName"`

	// Email A valid email
	Email         openapi_types.Email `json:"email"`
	EmailVerified bool                `json:"emailVerified"`

	// Id Id of the user
	Id          string `json:"id"`
	IsAnonymous bool   `json:"isAnonymous"`

	// Locale A two-characters locale
	Locale              string                 `json:"locale"`
	Metadata            map[string]interface{} `json:"metadata"`
	PhoneNumber         string                 `json:"phoneNumber"`
	PhoneNumberVerified bool                   `json:"phoneNumberVerified"`
	Roles               []string               `json:"roles"`
}

// UserEmailChangeRequest defines model for UserEmailChangeRequest.
type UserEmailChangeRequest struct {
	// NewEmail A valid email
	NewEmail openapi_types.Email `json:"newEmail"`
	Options  *OptionsRedirectTo  `json:"options,omitempty"`
}

// UserPasswordResetRequest defines model for UserPasswordResetRequest.
type UserPasswordResetRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *OptionsRedirectTo  `json:"options,omitempty"`
}

// PostPatJSONRequestBody defines body for PostPat for application/json ContentType.
type PostPatJSONRequestBody = CreatePATRequest

// PostSigninEmailPasswordJSONRequestBody defines body for PostSigninEmailPassword for application/json ContentType.
type PostSigninEmailPasswordJSONRequestBody = SignInEmailPasswordRequest

// PostSigninPasswordlessEmailJSONRequestBody defines body for PostSigninPasswordlessEmail for application/json ContentType.
type PostSigninPasswordlessEmailJSONRequestBody = SignInPasswordlessEmailRequest

// PostSigninPatJSONRequestBody defines body for PostSigninPat for application/json ContentType.
type PostSigninPatJSONRequestBody = SignInPATRequest

// PostSignupEmailPasswordJSONRequestBody defines body for PostSignupEmailPassword for application/json ContentType.
type PostSignupEmailPasswordJSONRequestBody = SignUpEmailPasswordRequest

// PostUserEmailChangeJSONRequestBody defines body for PostUserEmailChange for application/json ContentType.
type PostUserEmailChangeJSONRequestBody = UserEmailChangeRequest

// PostUserPasswordResetJSONRequestBody defines body for PostUserPasswordReset for application/json ContentType.
type PostUserPasswordResetJSONRequestBody = UserPasswordResetRequest
