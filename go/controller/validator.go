//go:generate mockgen -package mock -destination mock/validator.go --source=validator.go
package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"regexp"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/oapi-codegen/runtime/types"
)

type HIBPClient interface {
	IsPasswordPwned(ctx context.Context, password string) (bool, error)
}

type Validator struct {
	cfg                  *Config
	db                   DBClient
	hibp                 HIBPClient
	redirectURLValidator func(redirectTo string) bool
	ValidateEmail        func(email string) bool
}

func NewValidator(cfg *Config, db DBClient, hibp HIBPClient) (*Validator, error) {
	allowedURLs := make([]*url.URL, len(cfg.AllowedRedirectURLs)+1)
	allowedURLs[0] = cfg.ClientURL
	for i, u := range cfg.AllowedRedirectURLs {
		allowedURLs[i+1] = u
	}

	redirectURLValidator, err := ValidateRedirectTo(allowedURLs)
	if err != nil {
		return nil, fmt.Errorf("error creating redirect URL validator: %w", err)
	}

	emailValidator := ValidateEmail(
		cfg.BlockedEmailDomains,
		cfg.BlockedEmails,
		cfg.AllowedEmailDomains,
		cfg.AllowedEmails,
	)

	return &Validator{
		cfg:                  cfg,
		db:                   db,
		hibp:                 hibp,
		redirectURLValidator: redirectURLValidator,
		ValidateEmail:        emailValidator,
	}, nil
}

func (validator *Validator) SignupEmail(
	ctx context.Context, email types.Email, logger *slog.Logger,
) *APIError {
	_, err := validator.db.GetUserByEmail(ctx, sql.Text(email))
	if err == nil {
		logger.Warn("email already in use")
		return &APIError{api.EmailAlreadyInUse}
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		logger.Error("error getting user by email", logError(err))
		return &APIError{api.InternalServerError}
	}

	if !validator.ValidateEmail(string(email)) {
		logger.Warn("email didn't pass access control checks")
		return &APIError{api.InvalidEmailPassword}
	}

	return nil
}

func (validator *Validator) Password(
	ctx context.Context, password string, logger *slog.Logger,
) *APIError {
	if len(password) < validator.cfg.PasswordMinLength {
		logger.Warn("password too short")
		return &APIError{api.PasswordTooShort}
	}

	if validator.cfg.PasswordHIBPEnabled {
		if pwned, err := validator.hibp.IsPasswordPwned(ctx, password); err != nil {
			logger.Error("error checking password with HIBP", logError(err))
			return &APIError{api.InternalServerError}
		} else if pwned {
			logger.Warn("password is in HIBP database")
			return &APIError{api.PasswordInHibpDatabase}
		}
	}

	return nil
}

func (validator *Validator) SignUpOptions( //nolint:cyclop
	options *api.SignUpOptions, defaultName string, logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	if options == nil {
		options = &api.SignUpOptions{} //nolint:exhaustruct
	}

	if options.DefaultRole == nil {
		options.DefaultRole = ptr(validator.cfg.DefaultRole)
	}

	if options.AllowedRoles == nil {
		options.AllowedRoles = ptr(validator.cfg.DefaultAllowedRoles)
	} else {
		for _, role := range deptr(options.AllowedRoles) {
			if !slices.Contains(validator.cfg.DefaultAllowedRoles, role) {
				logger.Warn("role not allowed", slog.String("role", role))
				return nil, &APIError{api.RoleNotAllowed}
			}
		}
	}

	if !slices.Contains(deptr(options.AllowedRoles), deptr(options.DefaultRole)) {
		logger.Warn("default role not in allowed roles")
		return nil, &APIError{api.DefaultRoleMustBeInAllowedRoles}
	}

	if options.DisplayName == nil {
		options.DisplayName = &defaultName
	}

	if options.Locale == nil {
		options.Locale = ptr(validator.cfg.DefaultLocale)
	}
	if !slices.Contains(validator.cfg.AllowedLocales, deptr(options.Locale)) {
		logger.Warn(
			"locale not allowed, using default",
			slog.String("locale", deptr(options.Locale)),
		)
		options.Locale = ptr(validator.cfg.DefaultLocale)
	}

	if options.RedirectTo == nil {
		options.RedirectTo = ptr(validator.cfg.ClientURL.String())
	} else if !validator.redirectURLValidator(deptr(options.RedirectTo)) {
		logger.Warn("redirect URL not allowed", slog.String("redirectTo", deptr(options.RedirectTo)))
		return nil, &APIError{api.RedirecToNotAllowed}
	}

	return options, nil
}

func (validator *Validator) GetUser(
	ctx context.Context,
	id uuid.UUID,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	user, err := validator.db.GetUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("user not found")
		return sql.AuthUser{}, &APIError{api.InvalidEmailPassword} //nolint:exhaustruct
	}
	if err != nil {
		logger.Error("error getting user by email", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if err := validator.User(user, logger); err != nil {
		return sql.AuthUser{}, err //nolint:exhaustruct
	}

	return user, nil
}

func (validator *Validator) GetUserByEmail(
	ctx context.Context,
	email string,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	user, err := validator.db.GetUserByEmail(ctx, sql.Text(email))
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("user not found")
		return sql.AuthUser{}, &APIError{api.InvalidEmailPassword} //nolint:exhaustruct
	}
	if err != nil {
		logger.Error("error getting user by email", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if err := validator.User(user, logger); err != nil {
		return sql.AuthUser{}, err //nolint:exhaustruct
	}

	return user, nil
}

func (validator *Validator) GetUserByRefreshTokenHash(
	ctx context.Context,
	refreshToken string,
	refreshTokenType sql.RefreshTokenType,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	user, err := validator.db.GetUserByRefreshTokenHash(
		ctx,
		sql.GetUserByRefreshTokenHashParams{
			RefreshTokenHash: sql.Text(hashRefreshToken([]byte(refreshToken))),
			Type:             refreshTokenType,
		},
	)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Error("could not find user by refresh token")
		return sql.AuthUser{}, &APIError{api.InvalidPat} //nolint:exhaustruct
	}
	if err != nil {
		logger.Error("could not get user by refresh token", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if apiErr := validator.User(user, logger); apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	return user, nil
}

func (validator *Validator) User(
	user sql.AuthUser,
	logger *slog.Logger,
) *APIError {
	if !validator.ValidateEmail(user.Email.String) {
		logger.Warn("email didn't pass access control checks")
		return &APIError{api.InvalidEmailPassword}
	}

	if user.Disabled {
		logger.Warn("user is disabled")
		return &APIError{api.DisabledUser}
	}

	if !user.EmailVerified && validator.cfg.RequireEmailVerification {
		logger.Warn("user is unverified")
		return &APIError{api.UnverifiedUser}
	}

	return nil
}

func (validator *Validator) OptionsRedirectTo(
	options *api.OptionsRedirectTo,
	logger *slog.Logger,
) (*api.OptionsRedirectTo, *APIError) {
	if options == nil {
		options = &api.OptionsRedirectTo{} //nolint:exhaustruct
	}

	if options.RedirectTo == nil {
		options.RedirectTo = ptr(validator.cfg.ClientURL.String())
	} else if !validator.redirectURLValidator(deptr(options.RedirectTo)) {
		logger.Warn("redirect URL not allowed", slog.String("redirectTo", deptr(options.RedirectTo)))
		return nil, &APIError{api.RedirecToNotAllowed}
	}

	return options, nil
}

type urlMatcher struct {
	Scheme   string
	Hostname *regexp.Regexp
	Port     string
	Path     string
}

func newURLMatcher(u *url.URL) (urlMatcher, error) {
	if u.Scheme == "" {
		return urlMatcher{}, fmt.Errorf( //nolint:goerr113
			"scheme is required for allowed redirect URL",
		)
	}

	port := u.Port()
	if port == "" && u.Scheme == "http" {
		port = "80"
	} else if port == "" && u.Scheme == "https" {
		port = "443"
	}

	r := regexp.QuoteMeta(u.Hostname())
	r = "^" + strings.ReplaceAll(r, `\*`, `(\w+|\w+-\w+)+`) + `$`
	re, err := regexp.Compile(r)
	if err != nil {
		return urlMatcher{}, fmt.Errorf("error compiling regex: %w", err)
	}

	return urlMatcher{
		Scheme:   u.Scheme,
		Hostname: re,
		Port:     port,
		Path:     u.Path,
	}, nil
}

func (m urlMatcher) Matches(scheme, host, port, path string) bool {
	return m.Scheme == scheme &&
		m.Port == port &&
		m.Hostname.MatchString(host) &&
		strings.HasPrefix(path, m.Path)
}

func ValidateRedirectTo( //nolint:cyclop
	allowedRedirectURLs []*url.URL,
) (
	func(redirectTo string) bool,
	error,
) {
	matchers := make([]urlMatcher, len(allowedRedirectURLs))
	for i, u := range allowedRedirectURLs {
		m, err := newURLMatcher(u)
		if err != nil {
			return nil, err
		}

		matchers[i] = m
	}

	return func(redirectTo string) bool {
		u, err := url.Parse(redirectTo)
		if err != nil {
			return false
		}

		if u.Scheme == "" || u.Hostname() == "" {
			return false
		}

		if len(allowedRedirectURLs) == 0 {
			return true
		}

		port := u.Port()
		if port == "" && u.Scheme == "http" {
			port = "80"
		} else if port == "" && u.Scheme == "https" {
			port = "443"
		}

		for _, m := range matchers {
			if m.Matches(u.Scheme, u.Hostname(), port, u.Path) {
				return true
			}
		}

		return false
	}, nil
}

func ValidateEmail(
	blockedEmailDomains []string,
	blockedEmails []string,
	allowedEmailDomains []string,
	allowedEmails []string,
) func(email string) bool {
	return func(email string) bool {
		parts := strings.Split(email, "@")
		if len(parts) != 2 { //nolint:gomnd
			return false
		}

		domain := parts[1]

		if slices.Contains(blockedEmailDomains, domain) {
			return false
		}

		if slices.Contains(blockedEmails, email) {
			return false
		}

		if len(allowedEmailDomains) > 0 && !slices.Contains(allowedEmailDomains, domain) {
			return false
		}

		if len(allowedEmails) > 0 && !slices.Contains(allowedEmails, email) {
			return false
		}

		return true
	}
}
