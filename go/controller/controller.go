package controller

import "github.com/nhost/hasura-auth/go/sql"

type Config struct {
	ConcealErrors       bool
	DisableSignup       bool
	DisableNewUsers     bool
	DefaultAllowedRoles []string
	DefaultRole         string
	DefaultLocale       string
	GravatarEnabled     bool
	GravatarDefault     string
	GravatarRating      string
}

type Controller struct {
	db          *sql.Queries
	validator   *Validator
	config      Config
	gravatarURL func(string) string
}

func New(db *sql.Queries, config Config) *Controller {
	return &Controller{
		db:        db,
		config:    config,
		validator: NewValidator(&config),
		gravatarURL: GravatarURLFunc(
			config.GravatarEnabled, config.GravatarDefault, config.GravatarRating,
		),
	}
}
