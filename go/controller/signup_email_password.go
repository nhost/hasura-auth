package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nhost/hasura-auth/go/api"
	"github.com/nhost/hasura-auth/go/sql"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hash), nil
}

func deptr[T any](x *T) T { //nolint:ireturn
	if x == nil {
		return *new(T)
	}
	return *x
}

func ptr[T any](x T) *T {
	return &x
}

func (ctrl *Controller) PostSignupEmailPassword( //nolint:ireturn
	ctx context.Context,
	req api.PostSignupEmailPasswordRequestObject,
) (api.PostSignupEmailPasswordResponseObject, error) {
	// TODO: middleware?
	if ctrl.config.DisableSignup {
		return ctrl.sendError(api.SignupDisabled), nil
	}

	email := sql.Text(req.Body.Email)
	_, err := ctrl.db.GetUserByEmail(ctx, email)
	if err == nil {
		return ctrl.sendError(api.EmailAlreadyInUse), nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	hashedPassword, err := hashPassword(req.Body.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	options, err := ctrl.validator.SignUpOptions(req.Body.Options, string(req.Body.Email))
	if err != nil {
		return nil, fmt.Errorf("error getting signup options: %w", err)
	}

	metadata, err := json.Marshal(options.Metadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling metadata: %w", err)
	}

	gravatarURL := ctrl.gravatarURL(email.String)

	user, err := ctrl.db.InsertUser(
		ctx, sql.InsertUserParams{
			Disabled:        ctrl.config.DisableNewUsers,
			DisplayName:     deptr(options.DisplayName),
			AvatarUrl:       gravatarURL,
			Email:           email,
			PasswordHash:    sql.Text(hashedPassword),
			Ticket:          sql.Text("verifyEmail:" + uuid.NewString()),
			TicketExpiresAt: sql.TimestampTz(time.Now().Add(30 * 24 * time.Hour)),
			EmailVerified:   false,
			Locale:          deptr(options.Locale),
			DefaultRole:     deptr(options.DefaultRole),
			Metadata:        metadata,
			Roles:           deptr(options.AllowedRoles),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	return api.PostSignupEmailPassword200JSONResponse{
		Session: &api.Session{
			AccessToken:          "",
			AccessTokenExpiresIn: 0,
			RefreshToken:         "",
			User: &api.User{
				AvatarUrl:           gravatarURL,
				CreatedAt:           user.CreatedAt.Time,
				DefaultRole:         *options.DefaultRole,
				DisplayName:         deptr(options.DisplayName),
				Email:               openapi_types.Email(email.String),
				EmailVerified:       false,
				Id:                  ptr(user.UserID.String()),
				IsAnonymous:         false,
				Locale:              deptr(options.Locale),
				Metadata:            deptr(options.Metadata),
				PhoneNumber:         "",
				PhoneNumberVerified: false,
				Roles:               deptr(options.AllowedRoles),
			},
		},
	}, nil
}
