package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/sql"
	"github.com/oapi-codegen/runtime/types"
)

type HIBPClient interface {
	IsPasswordPwned(ctx context.Context, password string) (bool, error)
}

type Workflows struct {
	cfg                  *Config
	jwtGetter            JWTGetter
	db                   DBClient
	hibp                 HIBPClient
	redirectURLValidator func(redirectTo string) bool
	ValidateEmail        func(email string) bool
}

func NewWorkflows(
	cfg *Config, jwtGetter JWTGetter, db DBClient, hibp HIBPClient,
) (*Workflows, error) {
	allowedURLs := make([]*url.URL, len(cfg.AllowedRedirectURLs)+1)
	allowedURLs[0] = cfg.ClientURL
	for i, u := range cfg.AllowedRedirectURLs {
		allowedURLs[i+1] = u
	}

	redirectURLValidator, err := ValidateRedirectTo(allowedURLs)
	if err != nil {
		return nil, fmt.Errorf("error creating redirect URL wf: %w", err)
	}

	emailValidator := ValidateEmail(
		cfg.BlockedEmailDomains,
		cfg.BlockedEmails,
		cfg.AllowedEmailDomains,
		cfg.AllowedEmails,
	)

	return &Workflows{
		cfg:                  cfg,
		jwtGetter:            jwtGetter,
		db:                   db,
		hibp:                 hibp,
		redirectURLValidator: redirectURLValidator,
		ValidateEmail:        emailValidator,
	}, nil
}

func (wf *Workflows) ValidateSignupEmail(
	ctx context.Context, email types.Email, logger *slog.Logger,
) *APIError {
	_, err := wf.db.GetUserByEmail(ctx, sql.Text(email))
	if err == nil {
		logger.Warn("email already in use")
		return &APIError{api.EmailAlreadyInUse}
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		logger.Error("error getting user by email", logError(err))
		return &APIError{api.InternalServerError}
	}

	if !wf.ValidateEmail(string(email)) {
		logger.Warn("email didn't pass access control checks")
		return &APIError{api.InvalidEmailPassword}
	}

	return nil
}

func (wf *Workflows) ValidatePassword(
	ctx context.Context, password string, logger *slog.Logger,
) *APIError {
	if len(password) < wf.cfg.PasswordMinLength {
		logger.Warn("password too short")
		return &APIError{api.PasswordTooShort}
	}

	if wf.cfg.PasswordHIBPEnabled {
		if pwned, err := wf.hibp.IsPasswordPwned(ctx, password); err != nil {
			logger.Error("error checking password with HIBP", logError(err))
			return &APIError{api.InternalServerError}
		} else if pwned {
			logger.Warn("password is in HIBP database")
			return &APIError{api.PasswordInHibpDatabase}
		}
	}

	return nil
}

func (wf *Workflows) ValidateSignUpOptions( //nolint:cyclop
	options *api.SignUpOptions, defaultName string, logger *slog.Logger,
) (*api.SignUpOptions, *APIError) {
	if options == nil {
		options = &api.SignUpOptions{} //nolint:exhaustruct
	}

	if options.DefaultRole == nil {
		options.DefaultRole = ptr(wf.cfg.DefaultRole)
	}

	if options.AllowedRoles == nil {
		options.AllowedRoles = ptr(wf.cfg.DefaultAllowedRoles)
	} else {
		for _, role := range deptr(options.AllowedRoles) {
			if !slices.Contains(wf.cfg.DefaultAllowedRoles, role) {
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
		options.Locale = ptr(wf.cfg.DefaultLocale)
	}
	if !slices.Contains(wf.cfg.AllowedLocales, deptr(options.Locale)) {
		logger.Warn(
			"locale not allowed, using default",
			slog.String("locale", deptr(options.Locale)),
		)
		options.Locale = ptr(wf.cfg.DefaultLocale)
	}

	if options.RedirectTo == nil {
		options.RedirectTo = ptr(wf.cfg.ClientURL.String())
	} else if !wf.redirectURLValidator(deptr(options.RedirectTo)) {
		logger.Warn("redirect URL not allowed", slog.String("redirectTo", deptr(options.RedirectTo)))
		return nil, &APIError{api.RedirecToNotAllowed}
	}

	return options, nil
}

func (wf *Workflows) ValidateUser(
	user sql.AuthUser,
	logger *slog.Logger,
) *APIError {
	if !wf.ValidateEmail(user.Email.String) {
		logger.Warn("email didn't pass access control checks")
		return &APIError{api.InvalidEmailPassword}
	}

	if user.Disabled {
		logger.Warn("user is disabled")
		return &APIError{api.DisabledUser}
	}

	if !user.EmailVerified && wf.cfg.RequireEmailVerification {
		logger.Warn("user is unverified")
		return &APIError{api.UnverifiedUser}
	}

	return nil
}

func (wf *Workflows) ValidateOptionsRedirectTo(
	options *api.OptionsRedirectTo,
	logger *slog.Logger,
) (*api.OptionsRedirectTo, *APIError) {
	if options == nil {
		options = &api.OptionsRedirectTo{} //nolint:exhaustruct
	}

	if options.RedirectTo == nil {
		options.RedirectTo = ptr(wf.cfg.ClientURL.String())
	} else if !wf.redirectURLValidator(deptr(options.RedirectTo)) {
		logger.Warn("redirect URL not allowed", slog.String("redirectTo", deptr(options.RedirectTo)))
		return nil, &APIError{api.RedirecToNotAllowed}
	}

	return options, nil
}

func (wf *Workflows) GetUser(
	ctx context.Context,
	id uuid.UUID,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	user, err := wf.db.GetUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("user not found")
		return sql.AuthUser{}, &APIError{api.InvalidEmailPassword} //nolint:exhaustruct
	}
	if err != nil {
		logger.Error("error getting user by email", logError(err))
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if err := wf.ValidateUser(user, logger); err != nil {
		return sql.AuthUser{}, err //nolint:exhaustruct
	}

	return user, nil
}

func (wf *Workflows) GetUserByEmail(
	ctx context.Context,
	email string,
	logger *slog.Logger,
) (sql.AuthUser, bool, *APIError) {
	user, err := wf.db.GetUserByEmail(ctx, sql.Text(email))
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("user not found")
		return sql.AuthUser{}, true, &APIError{api.InvalidEmailPassword} //nolint:exhaustruct
	}
	if err != nil {
		logger.Error("error getting user by email", logError(err))
		return sql.AuthUser{}, false, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	if err := wf.ValidateUser(user, logger); err != nil {
		return sql.AuthUser{}, false, err //nolint:exhaustruct
	}

	return user, false, nil
}

func (wf *Workflows) GetUserByRefreshTokenHash(
	ctx context.Context,
	refreshToken string,
	refreshTokenType sql.RefreshTokenType,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	user, err := wf.db.GetUserByRefreshTokenHash(
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

	if apiErr := wf.ValidateUser(user, logger); apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	return user, nil
}

func (wf *Workflows) NewSession(
	ctx context.Context,
	user sql.AuthUser,
	logger *slog.Logger,
) (*api.Session, error) {
	userRoles, err := wf.db.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting roles by user id: %w", err)
	}
	allowedRoles := make([]string, len(userRoles))
	for i, role := range userRoles {
		allowedRoles[i] = role.Role
	}

	refreshToken := uuid.New()
	expiresAt := time.Now().Add(time.Duration(wf.cfg.RefreshTokenExpiresIn) * time.Second)
	if _, apiErr := wf.InsertRefreshtoken(
		ctx, user.ID, refreshToken.String(), expiresAt, sql.RefreshTokenTypeRegular, nil, logger,
	); apiErr != nil {
		return nil, apiErr
	}

	if _, err := wf.db.UpdateUserLastSeen(ctx, user.ID); err != nil {
		return nil, fmt.Errorf("error updating last seen: %w", err)
	}

	accessToken, expiresIn, err := wf.jwtGetter.GetToken(
		ctx, user.ID, allowedRoles, user.DefaultRole, logger,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting jwt: %w", err)
	}

	var metadata map[string]any
	if err := json.Unmarshal(user.Metadata, &metadata); err != nil {
		return nil, fmt.Errorf("error unmarshalling user metadata: %w", err)
	}
	return &api.Session{
		AccessToken:          accessToken,
		AccessTokenExpiresIn: expiresIn,
		RefreshToken:         refreshToken.String(),
		User: &api.User{
			AvatarUrl:           user.AvatarUrl,
			CreatedAt:           user.CreatedAt.Time,
			DefaultRole:         user.DefaultRole,
			DisplayName:         user.DisplayName,
			Email:               types.Email(user.Email.String),
			EmailVerified:       user.EmailVerified,
			Id:                  user.ID.String(),
			IsAnonymous:         false,
			Locale:              user.Locale,
			Metadata:            metadata,
			PhoneNumber:         user.PhoneNumber.String,
			PhoneNumberVerified: user.PhoneNumberVerified,
			Roles:               allowedRoles,
		},
	}, nil
}

func (wf *Workflows) GetUserFromJWTInContext(
	ctx context.Context,
	logger *slog.Logger,
) (sql.AuthUser, *APIError) {
	jwtToken, ok := wf.jwtGetter.FromContext(ctx)
	if !ok {
		logger.Error(
			"jwt token not found in context, this should not be possilble due to middleware",
		)
		return sql.AuthUser{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	sub, err := jwtToken.Claims.GetSubject()
	if err != nil {
		logger.Error("error getting user id from jwt token", logError(err))
		return sql.AuthUser{}, &APIError{api.InvalidRequest} //nolint:exhaustruct
	}
	logger = logger.With(slog.String("user_id", sub))

	userID, err := uuid.Parse(sub)
	if err != nil {
		logger.Error("error parsing user id from jwt token's subject", logError(err))
		return sql.AuthUser{}, &APIError{api.InvalidRequest} //nolint:exhaustruct
	}

	user, apiErr := wf.GetUser(ctx, userID, logger)
	if apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	if apiErr := wf.ValidateUser(user, logger); apiErr != nil {
		return sql.AuthUser{}, apiErr //nolint:exhaustruct
	}

	return user, nil
}

func (wf *Workflows) InsertRefreshtoken(
	ctx context.Context,
	userID uuid.UUID,
	refreshToken string,
	refreshTokenExpiresAt time.Time,
	refreshTokenType sql.RefreshTokenType,
	metadata map[string]any,
	logger *slog.Logger,
) (uuid.UUID, *APIError) {
	var b []byte
	var err error
	if metadata != nil {
		b, err = json.Marshal(metadata)
		if err != nil {
			logger.Error("error marshalling metadata", logError(err))
			return uuid.UUID{}, &APIError{api.InternalServerError} //nolint:exhaustruct
		}
	}

	refreshTokenID, err := wf.db.InsertRefreshtoken(ctx, sql.InsertRefreshtokenParams{
		UserID:           userID,
		RefreshTokenHash: sql.Text(hashRefreshToken([]byte(refreshToken))),
		ExpiresAt:        sql.TimestampTz(refreshTokenExpiresAt),
		Type:             refreshTokenType,
		Metadata:         b,
	})
	if err != nil {
		return uuid.UUID{}, &APIError{api.InternalServerError} //nolint:exhaustruct
	}

	return refreshTokenID, nil
}
