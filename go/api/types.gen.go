// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version 2.4.1 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
	"time"

	protocol "github.com/go-webauthn/webauthn/protocol"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes         = "BearerAuth.Scopes"
	BearerAuthElevatedScopes = "BearerAuthElevated.Scopes"
)

// Defines values for ErrorResponseError.
const (
	CannotSendSms                   ErrorResponseError = "cannot-send-sms"
	DefaultRoleMustBeInAllowedRoles ErrorResponseError = "default-role-must-be-in-allowed-roles"
	DisabledEndpoint                ErrorResponseError = "disabled-endpoint"
	DisabledMfaTotp                 ErrorResponseError = "disabled-mfa-totp"
	DisabledUser                    ErrorResponseError = "disabled-user"
	EmailAlreadyInUse               ErrorResponseError = "email-already-in-use"
	EmailAlreadyVerified            ErrorResponseError = "email-already-verified"
	ForbiddenAnonymous              ErrorResponseError = "forbidden-anonymous"
	InternalServerError             ErrorResponseError = "internal-server-error"
	InvalidEmailPassword            ErrorResponseError = "invalid-email-password"
	InvalidOtp                      ErrorResponseError = "invalid-otp"
	InvalidPat                      ErrorResponseError = "invalid-pat"
	InvalidRefreshToken             ErrorResponseError = "invalid-refresh-token"
	InvalidRequest                  ErrorResponseError = "invalid-request"
	InvalidState                    ErrorResponseError = "invalid-state"
	InvalidTicket                   ErrorResponseError = "invalid-ticket"
	InvalidTotp                     ErrorResponseError = "invalid-totp"
	LocaleNotAllowed                ErrorResponseError = "locale-not-allowed"
	MfaTypeNotFound                 ErrorResponseError = "mfa-type-not-found"
	NoTotpSecret                    ErrorResponseError = "no-totp-secret"
	OauthProfileFetchFailed         ErrorResponseError = "oauth-profile-fetch-failed"
	OauthProviderError              ErrorResponseError = "oauth-provider-error"
	OauthTokenEchangeFailed         ErrorResponseError = "oauth-token-echange-failed"
	PasswordInHibpDatabase          ErrorResponseError = "password-in-hibp-database"
	PasswordTooShort                ErrorResponseError = "password-too-short"
	RedirectToNotAllowed            ErrorResponseError = "redirectTo-not-allowed"
	RoleNotAllowed                  ErrorResponseError = "role-not-allowed"
	SignupDisabled                  ErrorResponseError = "signup-disabled"
	TotpAlreadyActive               ErrorResponseError = "totp-already-active"
	UnverifiedUser                  ErrorResponseError = "unverified-user"
	UserNotAnonymous                ErrorResponseError = "user-not-anonymous"
)

// Defines values for IdTokenProvider.
const (
	IdTokenProviderApple  IdTokenProvider = "apple"
	IdTokenProviderGoogle IdTokenProvider = "google"
)

// Defines values for OKResponse.
const (
	OK OKResponse = "OK"
)

// Defines values for UserDeanonymizeRequestSignInMethod.
const (
	EmailPassword UserDeanonymizeRequestSignInMethod = "email-password"
	Passwordless  UserDeanonymizeRequestSignInMethod = "passwordless"
)

// Defines values for UserMfaRequestActiveMfaType.
const (
	Empty UserMfaRequestActiveMfaType = ""
	Totp  UserMfaRequestActiveMfaType = "totp"
)

// Defines values for SigninProvider.
const (
	SigninProviderApple       SigninProvider = "apple"
	SigninProviderAzuread     SigninProvider = "azuread"
	SigninProviderBitbucket   SigninProvider = "bitbucket"
	SigninProviderDiscord     SigninProvider = "discord"
	SigninProviderFacebook    SigninProvider = "facebook"
	SigninProviderGithub      SigninProvider = "github"
	SigninProviderGitlab      SigninProvider = "gitlab"
	SigninProviderGoogle      SigninProvider = "google"
	SigninProviderLinkedin    SigninProvider = "linkedin"
	SigninProviderSpotify     SigninProvider = "spotify"
	SigninProviderStrava      SigninProvider = "strava"
	SigninProviderTwitch      SigninProvider = "twitch"
	SigninProviderTwitter     SigninProvider = "twitter"
	SigninProviderWindowslive SigninProvider = "windowslive"
	SigninProviderWorkos      SigninProvider = "workos"
)

// Defines values for TicketTypeQuery.
const (
	TicketTypeQueryEmailConfirmChange TicketTypeQuery = "emailConfirmChange"
	TicketTypeQueryEmailVerify        TicketTypeQuery = "emailVerify"
	TicketTypeQueryPasswordReset      TicketTypeQuery = "passwordReset"
	TicketTypeQuerySigninPasswordless TicketTypeQuery = "signinPasswordless"
)

// Defines values for GetSigninProviderProviderParamsProvider.
const (
	GetSigninProviderProviderParamsProviderApple       GetSigninProviderProviderParamsProvider = "apple"
	GetSigninProviderProviderParamsProviderAzuread     GetSigninProviderProviderParamsProvider = "azuread"
	GetSigninProviderProviderParamsProviderBitbucket   GetSigninProviderProviderParamsProvider = "bitbucket"
	GetSigninProviderProviderParamsProviderDiscord     GetSigninProviderProviderParamsProvider = "discord"
	GetSigninProviderProviderParamsProviderFacebook    GetSigninProviderProviderParamsProvider = "facebook"
	GetSigninProviderProviderParamsProviderGithub      GetSigninProviderProviderParamsProvider = "github"
	GetSigninProviderProviderParamsProviderGitlab      GetSigninProviderProviderParamsProvider = "gitlab"
	GetSigninProviderProviderParamsProviderGoogle      GetSigninProviderProviderParamsProvider = "google"
	GetSigninProviderProviderParamsProviderLinkedin    GetSigninProviderProviderParamsProvider = "linkedin"
	GetSigninProviderProviderParamsProviderSpotify     GetSigninProviderProviderParamsProvider = "spotify"
	GetSigninProviderProviderParamsProviderStrava      GetSigninProviderProviderParamsProvider = "strava"
	GetSigninProviderProviderParamsProviderTwitch      GetSigninProviderProviderParamsProvider = "twitch"
	GetSigninProviderProviderParamsProviderTwitter     GetSigninProviderProviderParamsProvider = "twitter"
	GetSigninProviderProviderParamsProviderWindowslive GetSigninProviderProviderParamsProvider = "windowslive"
	GetSigninProviderProviderParamsProviderWorkos      GetSigninProviderProviderParamsProvider = "workos"
)

// Defines values for GetSigninProviderProviderCallbackParamsProvider.
const (
	GetSigninProviderProviderCallbackParamsProviderApple       GetSigninProviderProviderCallbackParamsProvider = "apple"
	GetSigninProviderProviderCallbackParamsProviderAzuread     GetSigninProviderProviderCallbackParamsProvider = "azuread"
	GetSigninProviderProviderCallbackParamsProviderBitbucket   GetSigninProviderProviderCallbackParamsProvider = "bitbucket"
	GetSigninProviderProviderCallbackParamsProviderDiscord     GetSigninProviderProviderCallbackParamsProvider = "discord"
	GetSigninProviderProviderCallbackParamsProviderFacebook    GetSigninProviderProviderCallbackParamsProvider = "facebook"
	GetSigninProviderProviderCallbackParamsProviderGithub      GetSigninProviderProviderCallbackParamsProvider = "github"
	GetSigninProviderProviderCallbackParamsProviderGitlab      GetSigninProviderProviderCallbackParamsProvider = "gitlab"
	GetSigninProviderProviderCallbackParamsProviderGoogle      GetSigninProviderProviderCallbackParamsProvider = "google"
	GetSigninProviderProviderCallbackParamsProviderLinkedin    GetSigninProviderProviderCallbackParamsProvider = "linkedin"
	GetSigninProviderProviderCallbackParamsProviderSpotify     GetSigninProviderProviderCallbackParamsProvider = "spotify"
	GetSigninProviderProviderCallbackParamsProviderStrava      GetSigninProviderProviderCallbackParamsProvider = "strava"
	GetSigninProviderProviderCallbackParamsProviderTwitch      GetSigninProviderProviderCallbackParamsProvider = "twitch"
	GetSigninProviderProviderCallbackParamsProviderTwitter     GetSigninProviderProviderCallbackParamsProvider = "twitter"
	GetSigninProviderProviderCallbackParamsProviderWindowslive GetSigninProviderProviderCallbackParamsProvider = "windowslive"
	GetSigninProviderProviderCallbackParamsProviderWorkos      GetSigninProviderProviderCallbackParamsProvider = "workos"
)

// Defines values for PostSigninProviderProviderCallbackParamsProvider.
const (
	Apple       PostSigninProviderProviderCallbackParamsProvider = "apple"
	Azuread     PostSigninProviderProviderCallbackParamsProvider = "azuread"
	Bitbucket   PostSigninProviderProviderCallbackParamsProvider = "bitbucket"
	Discord     PostSigninProviderProviderCallbackParamsProvider = "discord"
	Facebook    PostSigninProviderProviderCallbackParamsProvider = "facebook"
	Github      PostSigninProviderProviderCallbackParamsProvider = "github"
	Gitlab      PostSigninProviderProviderCallbackParamsProvider = "gitlab"
	Google      PostSigninProviderProviderCallbackParamsProvider = "google"
	Linkedin    PostSigninProviderProviderCallbackParamsProvider = "linkedin"
	Spotify     PostSigninProviderProviderCallbackParamsProvider = "spotify"
	Strava      PostSigninProviderProviderCallbackParamsProvider = "strava"
	Twitch      PostSigninProviderProviderCallbackParamsProvider = "twitch"
	Twitter     PostSigninProviderProviderCallbackParamsProvider = "twitter"
	Windowslive PostSigninProviderProviderCallbackParamsProvider = "windowslive"
	Workos      PostSigninProviderProviderCallbackParamsProvider = "workos"
)

// Defines values for GetVerifyParamsType.
const (
	GetVerifyParamsTypeEmailConfirmChange GetVerifyParamsType = "emailConfirmChange"
	GetVerifyParamsTypeEmailVerify        GetVerifyParamsType = "emailVerify"
	GetVerifyParamsTypePasswordReset      GetVerifyParamsType = "passwordReset"
	GetVerifyParamsTypeSigninPasswordless GetVerifyParamsType = "signinPasswordless"
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

// IdTokenProvider defines model for IdTokenProvider.
type IdTokenProvider string

// JWK defines model for JWK.
type JWK struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

// JWKSet defines model for JWKSet.
type JWKSet struct {
	Keys []JWK `json:"keys"`
}

// LinkIdTokenRequest defines model for LinkIdTokenRequest.
type LinkIdTokenRequest struct {
	// IdToken Apple ID token
	IdToken string `json:"idToken"`

	// Nonce Nonce used during sign in process
	Nonce    *string         `json:"nonce,omitempty"`
	Provider IdTokenProvider `json:"provider"`
}

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

// RefreshTokenRequest defines model for RefreshTokenRequest.
type RefreshTokenRequest struct {
	// RefreshToken Refresh Token
	RefreshToken string `json:"refreshToken"`
}

// Session defines model for Session.
type Session struct {
	AccessToken          string `json:"accessToken"`
	AccessTokenExpiresIn int64  `json:"accessTokenExpiresIn"`

	// RefreshToken Refresh token during authentication or when refreshing the JWT
	RefreshToken string `json:"refreshToken"`

	// RefreshTokenId Refresh token id
	RefreshTokenId string `json:"refreshTokenId"`
	User           *User  `json:"user,omitempty"`
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

// SignInIdTokenRequest defines model for SignInIdTokenRequest.
type SignInIdTokenRequest struct {
	// IdToken Apple ID token
	IdToken string `json:"idToken"`

	// Nonce Nonce used during sign in process
	Nonce    *string         `json:"nonce,omitempty"`
	Options  *SignUpOptions  `json:"options,omitempty"`
	Provider IdTokenProvider `json:"provider"`
}

// SignInMfaTotpRequest defines model for SignInMfaTotpRequest.
type SignInMfaTotpRequest struct {
	// Otp One time password
	Otp string `json:"otp"`

	// Ticket Ticket
	Ticket string `json:"ticket"`
}

// SignInOTPEmailRequest defines model for SignInOTPEmailRequest.
type SignInOTPEmailRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *SignUpOptions      `json:"options,omitempty"`
}

// SignInOTPEmailVerifyRequest defines model for SignInOTPEmailVerifyRequest.
type SignInOTPEmailVerifyRequest struct {
	// Email A valid email
	Email openapi_types.Email `json:"email"`

	// Otp One time password
	Otp string `json:"otp"`
}

// SignInOTPEmailVerifyResponse defines model for SignInOTPEmailVerifyResponse.
type SignInOTPEmailVerifyResponse struct {
	Session *Session `json:"session,omitempty"`
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

// SignInPasswordlessSmsOtpRequest defines model for SignInPasswordlessSmsOtpRequest.
type SignInPasswordlessSmsOtpRequest struct {
	// Otp One-time password received by SMS
	Otp string `json:"otp"`

	// PhoneNumber Phone number of the user
	PhoneNumber string `json:"phoneNumber"`
}

// SignInPasswordlessSmsOtpResponse defines model for SignInPasswordlessSmsOtpResponse.
type SignInPasswordlessSmsOtpResponse struct {
	Mfa     *MFAChallengePayload `json:"mfa,omitempty"`
	Session *Session             `json:"session,omitempty"`
}

// SignInPasswordlessSmsRequest defines model for SignInPasswordlessSmsRequest.
type SignInPasswordlessSmsRequest struct {
	Options *SignUpOptions `json:"options,omitempty"`

	// PhoneNumber Phone number of the user
	PhoneNumber string `json:"phoneNumber"`
}

// SignInWebauthnRequest defines model for SignInWebauthnRequest.
type SignInWebauthnRequest struct {
	// Email A valid email
	Email *openapi_types.Email `json:"email,omitempty"`
}

// SignInWebauthnResponse defines model for SignInWebauthnResponse.
type SignInWebauthnResponse = protocol.PublicKeyCredentialRequestOptions

// SignInWebauthnVerifyRequest defines model for SignInWebauthnVerifyRequest.
type SignInWebauthnVerifyRequest struct {
	Credential protocol.CredentialAssertionResponse `json:"credential"`

	// Email A valid email. Deprecated, no longer used
	// Deprecated:
	Email *openapi_types.Email `json:"email,omitempty"`
}

// SignOutRequest defines model for SignOutRequest.
type SignOutRequest struct {
	// All Sign out from all connected devices
	All *bool `json:"all,omitempty"`

	// RefreshToken Refresh Token
	RefreshToken *string `json:"refreshToken,omitempty"`
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

// SignUpWebauthnRequest defines model for SignUpWebauthnRequest.
type SignUpWebauthnRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *SignUpOptions      `json:"options,omitempty"`
}

// SignUpWebauthnResponse defines model for SignUpWebauthnResponse.
type SignUpWebauthnResponse = protocol.PublicKeyCredentialCreationOptions

// SignUpWebauthnVerifyRequest defines model for SignUpWebauthnVerifyRequest.
type SignUpWebauthnVerifyRequest struct {
	Credential protocol.CredentialCreationResponse `json:"credential"`

	// Nickname Nickname for the security key
	Nickname *string        `json:"nickname,omitempty"`
	Options  *SignUpOptions `json:"options,omitempty"`
}

// SigninAnonymousRequest defines model for SigninAnonymousRequest.
type SigninAnonymousRequest struct {
	DisplayName *string `json:"displayName,omitempty"`

	// Locale A two-characters locale
	Locale   *string                 `json:"locale,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
}

// TotpGenerateResponse defines model for TotpGenerateResponse.
type TotpGenerateResponse struct {
	// ImageUrl URL to QR code image for TOTP setup
	ImageUrl string `json:"imageUrl"`

	// TotpSecret TOTP secret for manual setup
	TotpSecret string `json:"totpSecret"`
}

// User defines model for User.
type User struct {
	// ActiveMfaType Active MFA type for the user
	ActiveMfaType *string   `json:"activeMfaType"`
	AvatarUrl     string    `json:"avatarUrl"`
	CreatedAt     time.Time `json:"createdAt"`
	DefaultRole   string    `json:"defaultRole"`
	DisplayName   string    `json:"displayName"`

	// Email A valid email
	Email         *openapi_types.Email `json:"email,omitempty"`
	EmailVerified bool                 `json:"emailVerified"`

	// Id Id of the user
	Id          string `json:"id"`
	IsAnonymous bool   `json:"isAnonymous"`

	// Locale A two-characters locale
	Locale              string                 `json:"locale"`
	Metadata            map[string]interface{} `json:"metadata"`
	PhoneNumber         *string                `json:"phoneNumber,omitempty"`
	PhoneNumberVerified bool                   `json:"phoneNumberVerified"`
	Roles               []string               `json:"roles"`
}

// UserAddSecurityKeyVerifyRequest defines model for UserAddSecurityKeyVerifyRequest.
type UserAddSecurityKeyVerifyRequest struct {
	Credential protocol.CredentialCreationResponse `json:"credential"`

	// Nickname Optional nickname for the security key
	Nickname *string `json:"nickname,omitempty"`
}

// UserAddSecurityKeyVerifyResponse defines model for UserAddSecurityKeyVerifyResponse.
type UserAddSecurityKeyVerifyResponse struct {
	// Id ID of the newly added security key
	Id string `json:"id"`

	// Nickname Nickname of the security key
	Nickname *string `json:"nickname,omitempty"`
}

// UserDeanonymizeRequest defines model for UserDeanonymizeRequest.
type UserDeanonymizeRequest struct {
	// Connection Deprecated, will be ignored
	// Deprecated:
	Connection *string `json:"connection,omitempty"`

	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *SignUpOptions      `json:"options,omitempty"`

	// Password A password of minimum 3 characters
	Password *string `json:"password,omitempty"`

	// SignInMethod Which sign-in method to use
	SignInMethod UserDeanonymizeRequestSignInMethod `json:"signInMethod"`
}

// UserDeanonymizeRequestSignInMethod Which sign-in method to use
type UserDeanonymizeRequestSignInMethod string

// UserEmailChangeRequest defines model for UserEmailChangeRequest.
type UserEmailChangeRequest struct {
	// NewEmail A valid email
	NewEmail openapi_types.Email `json:"newEmail"`
	Options  *OptionsRedirectTo  `json:"options,omitempty"`
}

// UserEmailSendVerificationEmailRequest defines model for UserEmailSendVerificationEmailRequest.
type UserEmailSendVerificationEmailRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *OptionsRedirectTo  `json:"options,omitempty"`
}

// UserMfaRequest defines model for UserMfaRequest.
type UserMfaRequest struct {
	// ActiveMfaType Type of MFA to activate. If not set or set to empty string, disable MFA
	ActiveMfaType *UserMfaRequestActiveMfaType `json:"activeMfaType,omitempty"`

	// Code MFA activation code
	Code string `json:"code"`
}

// UserMfaRequestActiveMfaType Type of MFA to activate. If not set or set to empty string, disable MFA
type UserMfaRequestActiveMfaType string

// UserPasswordRequest defines model for UserPasswordRequest.
type UserPasswordRequest struct {
	// NewPassword A password of minimum 3 characters
	NewPassword string `json:"newPassword"`

	// Ticket Ticket to reset the password, required if the user is not authenticated
	Ticket *string `json:"ticket,omitempty"`
}

// UserPasswordResetRequest defines model for UserPasswordResetRequest.
type UserPasswordResetRequest struct {
	// Email A valid email
	Email   openapi_types.Email `json:"email"`
	Options *OptionsRedirectTo  `json:"options,omitempty"`
}

// VerifyTokenRequest defines model for VerifyTokenRequest.
type VerifyTokenRequest struct {
	// Token JWT token to verify
	Token *string `json:"token,omitempty"`
}

// RedirectToQuery Target URL for the redirect
type RedirectToQuery = string

// SigninProvider defines model for SigninProvider.
type SigninProvider string

// TicketQuery Ticket
type TicketQuery = string

// TicketTypeQuery Type of the ticket
type TicketTypeQuery string

// GetSigninProviderProviderParams defines parameters for GetSigninProviderProvider.
type GetSigninProviderProviderParams struct {
	// AllowedRoles Array of allowed roles for the user
	AllowedRoles *[]string `form:"allowedRoles,omitempty" json:"allowedRoles,omitempty"`

	// DefaultRole Default role for the user
	DefaultRole *string `form:"defaultRole,omitempty" json:"defaultRole,omitempty"`

	// DisplayName Display name for the user
	DisplayName *string `form:"displayName,omitempty" json:"displayName,omitempty"`

	// Locale A two-characters locale
	Locale *string `form:"locale,omitempty" json:"locale,omitempty"`

	// Metadata Additional metadata for the user
	Metadata *map[string]interface{} `form:"metadata,omitempty" json:"metadata,omitempty"`

	// RedirectTo URI to redirect to
	RedirectTo *string `form:"redirectTo,omitempty" json:"redirectTo,omitempty"`

	// Connect If set, this means that the user is already authenticated and wants to link their account. This needs to be a valid JWT access token.
	Connect *string `form:"connect,omitempty" json:"connect,omitempty"`
}

// GetSigninProviderProviderParamsProvider defines parameters for GetSigninProviderProvider.
type GetSigninProviderProviderParamsProvider string

// GetSigninProviderProviderCallbackParams defines parameters for GetSigninProviderProviderCallback.
type GetSigninProviderProviderCallbackParams struct {
	// Code Authorization code provided by the authentication provider
	Code *string `form:"code,omitempty" json:"code,omitempty"`

	// IdToken ID token provided by the authentication provider
	IdToken *string `form:"id_token,omitempty" json:"id_token,omitempty"`

	// State State parameter to avoid CSRF attacks
	State string `form:"state" json:"state"`

	// OauthToken OAuth token for the provider (e.g., X)
	OauthToken *string `form:"oauth_token,omitempty" json:"oauth_token,omitempty"`

	// OauthVerifier OAuth verifier for the provider (e.g., X)
	OauthVerifier *string `form:"oauth_verifier,omitempty" json:"oauth_verifier,omitempty"`

	// Error Error message if authentication failed
	Error *string `form:"error,omitempty" json:"error,omitempty"`

	// ErrorDescription Detailed error description if authentication failed
	ErrorDescription *string `form:"error_description,omitempty" json:"error_description,omitempty"`

	// ErrorUri URI with more information about the error
	ErrorUri *string `form:"error_uri,omitempty" json:"error_uri,omitempty"`
}

// GetSigninProviderProviderCallbackParamsProvider defines parameters for GetSigninProviderProviderCallback.
type GetSigninProviderProviderCallbackParamsProvider string

// PostSigninProviderProviderCallbackFormdataBody defines parameters for PostSigninProviderProviderCallback.
type PostSigninProviderProviderCallbackFormdataBody struct {
	// Code Authorization code provided by the authentication provider
	Code *string `form:"code" json:"code"`

	// Error Error message if authentication failed
	Error *string `form:"error" json:"error"`

	// ErrorDescription Detailed error description if authentication failed
	ErrorDescription *string `form:"error_description" json:"error_description"`

	// ErrorUri URI with more information about the error
	ErrorUri *string `form:"error_uri" json:"error_uri"`

	// IdToken ID token provided by the authentication provider
	IdToken *string `form:"id_token" json:"id_token"`

	// State State parameter to avoid CSRF attacks
	State string `form:"state" json:"state"`

	// User JSON string containing user information (only provided on first authentication with Apple)
	User                 *string                `form:"user" json:"user"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// PostSigninProviderProviderCallbackParamsProvider defines parameters for PostSigninProviderProviderCallback.
type PostSigninProviderProviderCallbackParamsProvider string

// GetVerifyParams defines parameters for GetVerify.
type GetVerifyParams struct {
	// Ticket Ticket
	Ticket TicketQuery `form:"ticket" json:"ticket"`

	// Type Type of the ticket. Deprecated, no longer used
	Type *GetVerifyParamsType `form:"type,omitempty" json:"type,omitempty"`

	// RedirectTo Target URL for the redirect
	RedirectTo RedirectToQuery `form:"redirectTo" json:"redirectTo"`
}

// GetVerifyParamsType defines parameters for GetVerify.
type GetVerifyParamsType string

// PostElevateWebauthnVerifyJSONRequestBody defines body for PostElevateWebauthnVerify for application/json ContentType.
type PostElevateWebauthnVerifyJSONRequestBody = SignInWebauthnVerifyRequest

// PostLinkIdtokenJSONRequestBody defines body for PostLinkIdtoken for application/json ContentType.
type PostLinkIdtokenJSONRequestBody = LinkIdTokenRequest

// PostPatJSONRequestBody defines body for PostPat for application/json ContentType.
type PostPatJSONRequestBody = CreatePATRequest

// PostSigninAnonymousJSONRequestBody defines body for PostSigninAnonymous for application/json ContentType.
type PostSigninAnonymousJSONRequestBody = SigninAnonymousRequest

// PostSigninEmailPasswordJSONRequestBody defines body for PostSigninEmailPassword for application/json ContentType.
type PostSigninEmailPasswordJSONRequestBody = SignInEmailPasswordRequest

// PostSigninIdtokenJSONRequestBody defines body for PostSigninIdtoken for application/json ContentType.
type PostSigninIdtokenJSONRequestBody = SignInIdTokenRequest

// PostSigninMfaTotpJSONRequestBody defines body for PostSigninMfaTotp for application/json ContentType.
type PostSigninMfaTotpJSONRequestBody = SignInMfaTotpRequest

// PostSigninOtpEmailJSONRequestBody defines body for PostSigninOtpEmail for application/json ContentType.
type PostSigninOtpEmailJSONRequestBody = SignInOTPEmailRequest

// PostSigninOtpEmailVerifyJSONRequestBody defines body for PostSigninOtpEmailVerify for application/json ContentType.
type PostSigninOtpEmailVerifyJSONRequestBody = SignInOTPEmailVerifyRequest

// PostSigninPasswordlessEmailJSONRequestBody defines body for PostSigninPasswordlessEmail for application/json ContentType.
type PostSigninPasswordlessEmailJSONRequestBody = SignInPasswordlessEmailRequest

// PostSigninPasswordlessSmsJSONRequestBody defines body for PostSigninPasswordlessSms for application/json ContentType.
type PostSigninPasswordlessSmsJSONRequestBody = SignInPasswordlessSmsRequest

// PostSigninPasswordlessSmsOtpJSONRequestBody defines body for PostSigninPasswordlessSmsOtp for application/json ContentType.
type PostSigninPasswordlessSmsOtpJSONRequestBody = SignInPasswordlessSmsOtpRequest

// PostSigninPatJSONRequestBody defines body for PostSigninPat for application/json ContentType.
type PostSigninPatJSONRequestBody = SignInPATRequest

// PostSigninProviderProviderCallbackFormdataRequestBody defines body for PostSigninProviderProviderCallback for application/x-www-form-urlencoded ContentType.
type PostSigninProviderProviderCallbackFormdataRequestBody PostSigninProviderProviderCallbackFormdataBody

// PostSigninWebauthnJSONRequestBody defines body for PostSigninWebauthn for application/json ContentType.
type PostSigninWebauthnJSONRequestBody = SignInWebauthnRequest

// PostSigninWebauthnVerifyJSONRequestBody defines body for PostSigninWebauthnVerify for application/json ContentType.
type PostSigninWebauthnVerifyJSONRequestBody = SignInWebauthnVerifyRequest

// PostSignoutJSONRequestBody defines body for PostSignout for application/json ContentType.
type PostSignoutJSONRequestBody = SignOutRequest

// PostSignupEmailPasswordJSONRequestBody defines body for PostSignupEmailPassword for application/json ContentType.
type PostSignupEmailPasswordJSONRequestBody = SignUpEmailPasswordRequest

// PostSignupWebauthnJSONRequestBody defines body for PostSignupWebauthn for application/json ContentType.
type PostSignupWebauthnJSONRequestBody = SignUpWebauthnRequest

// PostSignupWebauthnVerifyJSONRequestBody defines body for PostSignupWebauthnVerify for application/json ContentType.
type PostSignupWebauthnVerifyJSONRequestBody = SignUpWebauthnVerifyRequest

// PostTokenJSONRequestBody defines body for PostToken for application/json ContentType.
type PostTokenJSONRequestBody = RefreshTokenRequest

// PostTokenVerifyJSONRequestBody defines body for PostTokenVerify for application/json ContentType.
type PostTokenVerifyJSONRequestBody = VerifyTokenRequest

// PostUserDeanonymizeJSONRequestBody defines body for PostUserDeanonymize for application/json ContentType.
type PostUserDeanonymizeJSONRequestBody = UserDeanonymizeRequest

// PostUserEmailChangeJSONRequestBody defines body for PostUserEmailChange for application/json ContentType.
type PostUserEmailChangeJSONRequestBody = UserEmailChangeRequest

// PostUserEmailSendVerificationEmailJSONRequestBody defines body for PostUserEmailSendVerificationEmail for application/json ContentType.
type PostUserEmailSendVerificationEmailJSONRequestBody = UserEmailSendVerificationEmailRequest

// PostUserMfaJSONRequestBody defines body for PostUserMfa for application/json ContentType.
type PostUserMfaJSONRequestBody = UserMfaRequest

// PostUserPasswordJSONRequestBody defines body for PostUserPassword for application/json ContentType.
type PostUserPasswordJSONRequestBody = UserPasswordRequest

// PostUserPasswordResetJSONRequestBody defines body for PostUserPasswordReset for application/json ContentType.
type PostUserPasswordResetJSONRequestBody = UserPasswordResetRequest

// PostUserWebauthnVerifyJSONRequestBody defines body for PostUserWebauthnVerify for application/json ContentType.
type PostUserWebauthnVerifyJSONRequestBody = UserAddSecurityKeyVerifyRequest

// Getter for additional properties for PostSigninProviderProviderCallbackFormdataBody. Returns the specified
// element and whether it was found
func (a PostSigninProviderProviderCallbackFormdataBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PostSigninProviderProviderCallbackFormdataBody
func (a *PostSigninProviderProviderCallbackFormdataBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PostSigninProviderProviderCallbackFormdataBody to handle AdditionalProperties
func (a *PostSigninProviderProviderCallbackFormdataBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["code"]; found {
		err = json.Unmarshal(raw, &a.Code)
		if err != nil {
			return fmt.Errorf("error reading 'code': %w", err)
		}
		delete(object, "code")
	}

	if raw, found := object["error"]; found {
		err = json.Unmarshal(raw, &a.Error)
		if err != nil {
			return fmt.Errorf("error reading 'error': %w", err)
		}
		delete(object, "error")
	}

	if raw, found := object["error_description"]; found {
		err = json.Unmarshal(raw, &a.ErrorDescription)
		if err != nil {
			return fmt.Errorf("error reading 'error_description': %w", err)
		}
		delete(object, "error_description")
	}

	if raw, found := object["error_uri"]; found {
		err = json.Unmarshal(raw, &a.ErrorUri)
		if err != nil {
			return fmt.Errorf("error reading 'error_uri': %w", err)
		}
		delete(object, "error_uri")
	}

	if raw, found := object["id_token"]; found {
		err = json.Unmarshal(raw, &a.IdToken)
		if err != nil {
			return fmt.Errorf("error reading 'id_token': %w", err)
		}
		delete(object, "id_token")
	}

	if raw, found := object["state"]; found {
		err = json.Unmarshal(raw, &a.State)
		if err != nil {
			return fmt.Errorf("error reading 'state': %w", err)
		}
		delete(object, "state")
	}

	if raw, found := object["user"]; found {
		err = json.Unmarshal(raw, &a.User)
		if err != nil {
			return fmt.Errorf("error reading 'user': %w", err)
		}
		delete(object, "user")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for PostSigninProviderProviderCallbackFormdataBody to handle AdditionalProperties
func (a PostSigninProviderProviderCallbackFormdataBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Code != nil {
		object["code"], err = json.Marshal(a.Code)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'code': %w", err)
		}
	}

	if a.Error != nil {
		object["error"], err = json.Marshal(a.Error)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'error': %w", err)
		}
	}

	if a.ErrorDescription != nil {
		object["error_description"], err = json.Marshal(a.ErrorDescription)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'error_description': %w", err)
		}
	}

	if a.ErrorUri != nil {
		object["error_uri"], err = json.Marshal(a.ErrorUri)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'error_uri': %w", err)
		}
	}

	if a.IdToken != nil {
		object["id_token"], err = json.Marshal(a.IdToken)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'id_token': %w", err)
		}
	}

	object["state"], err = json.Marshal(a.State)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'state': %w", err)
	}

	if a.User != nil {
		object["user"], err = json.Marshal(a.User)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'user': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
