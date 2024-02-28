//go:generate mockgen -package mock -destination mock/controller.go --source=controller.go
package controller

import (
	"fmt"

	"github.com/nhost/hasura-auth/go/notifications"
	"github.com/nhost/hasura-auth/go/sql"
)

type Emailer interface {
	SendEmailVerify(to string, locale string, data notifications.EmailVerifyData) error
}

type Controller struct {
	db          *sql.Queries
	validator   *Validator
	config      Config
	gravatarURL func(string) string
	jwtGetter   *JWTGetter
	email       Emailer
}

func New(
	db *sql.Queries,
	config Config,
	jwtGetter *JWTGetter,
	emailer Emailer,
	hibp HIBPClient,
) (*Controller, error) {
	validator, err := NewValidator(&config, db, hibp)
	if err != nil {
		return nil, fmt.Errorf("error creating validator: %w", err)
	}

	return &Controller{
		db:        db,
		config:    config,
		validator: validator,
		gravatarURL: GravatarURLFunc(
			config.GravatarEnabled, config.GravatarDefault, config.GravatarRating,
		),
		jwtGetter: jwtGetter,
		email:     emailer,
	}, nil
}
