// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version (devel) DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for ErrorResponseError.
const (
	DefaultRoleMustBeInAllowedRoles ErrorResponseError = "default-role-must-be-in-allowed-roles"
	EmailAlreadyInUse               ErrorResponseError = "email-already-in-use"
	InternalServerError             ErrorResponseError = "internal-server-error"
	InvalidRequest                  ErrorResponseError = "invalid-request"
	LocaleNotAllowed                ErrorResponseError = "locale-not-allowed"
	PasswordInHibpDatabase          ErrorResponseError = "password-in-hibp-database"
	PasswordTooShort                ErrorResponseError = "password-too-short"
	RedirecToNotAllowed             ErrorResponseError = "redirecTo-not-allowed"
	RoleNotAllowed                  ErrorResponseError = "role-not-allowed"
	SignupDisabled                  ErrorResponseError = "signup-disabled"
)

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

// SignUpEmailPasswordSchema defines model for SignUpEmailPasswordSchema.
type SignUpEmailPasswordSchema struct {
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
	Id          *string `json:"id,omitempty"`
	IsAnonymous bool    `json:"isAnonymous"`

	// Locale A two-characters locale
	Locale              string                 `json:"locale"`
	Metadata            map[string]interface{} `json:"metadata"`
	PhoneNumber         string                 `json:"phoneNumber"`
	PhoneNumberVerified bool                   `json:"phoneNumberVerified"`
	Roles               []string               `json:"roles"`
}

// PostSignupEmailPasswordJSONRequestBody defines body for PostSignupEmailPassword for application/json ContentType.
type PostSignupEmailPasswordJSONRequestBody = SignUpEmailPasswordSchema
