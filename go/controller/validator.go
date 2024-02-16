package controller

import "github.com/nhost/hasura-auth/go/api"

type Validator struct {
	cfg *Config
}

func NewValidator(cfg *Config) *Validator {
	return &Validator{cfg: cfg}
}

func (validator *Validator) SignUpOptions(
	options *api.SignUpOptions, defaultName string,
) (*api.SignUpOptions, error) {
	if options == nil {
		options = new(api.SignUpOptions)
	}

	if options.AllowedRoles == nil {
		options.AllowedRoles = ptr(validator.cfg.DefaultAllowedRoles)
	}

	if options.DefaultRole == nil {
		options.DefaultRole = ptr(validator.cfg.DefaultRole)
	}

	if options.DisplayName == nil {
		options.DisplayName = &defaultName
	}

	if options.Locale == nil {
		options.Locale = ptr(validator.cfg.DefaultLocale)
	}

	return options, nil
}
