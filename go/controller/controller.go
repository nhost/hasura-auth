package controller

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
)

type Config struct {
	ClientURL                *url.URL
	ConcealErrors            bool
	DisableSignup            bool
	DisableNewUsers          bool
	DefaultAllowedRoles      []string
	DefaultRole              string
	DefaultLocale            string
	GravatarEnabled          bool
	GravatarDefault          string
	GravatarRating           string
	RefreshTokenExpiresIn    int
	AccessTokenExpiresIn     int
	JWTSecret                string
	RequireEmailVerification bool
	ServerURL                *url.URL
}

//go:generate mockgen -package mock -destination mock/emailer.go --source=controller.go Emailer
type Emailer interface {
	SendEmailVerify(to string, locale string, data notifications.EmailVerifyData) error
}

type Controller struct {
	db          *sql.Queries
	validator   *Validator
	config      Config
	gravatarURL func(string) string
	jwtGetter   func(uuid.UUID) (string, int64, error)
	email       Emailer
}

func New(db *sql.Queries, config Config, emailer Emailer) (*Controller, error) {
	jwtGetter, err := NewJWTGetter(
		[]byte(config.JWTSecret),
		time.Duration(config.AccessTokenExpiresIn)*time.Second,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating jwt getter: %w", err)
	}
	return &Controller{
		db:        db,
		config:    config,
		validator: NewValidator(&config),
		gravatarURL: GravatarURLFunc(
			config.GravatarEnabled, config.GravatarDefault, config.GravatarRating,
		),
		jwtGetter: jwtGetter,
		email:     emailer,
	}, nil
}
