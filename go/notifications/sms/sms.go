package sms

import (
	"fmt"
	"time"

	"github.com/nhost/hasura-auth/go/notifications"
)

const (
	in5Minutes  = 5 * 60 * time.Second
	in10Minutes = 10 * 60 * time.Second
)

func ptr[T any](v T) *T {
	return &v
}

func deptr[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

//go:generate mockgen -source=sms.go -destination=mock/sms_backend_mock.go -package=mock GenericSMSProvider
type GenericSMSProvider interface {
	SendSMS(to string, body string) error
}

type SMS struct {
	backend      GenericSMSProvider
	otpGenerator func() (string, string, error)
	templates    *notifications.Templates
}

func NewSMS(
	backend GenericSMSProvider,
	otpGenerator func() (string, string, error),
	templates *notifications.Templates,
) *SMS {
	return &SMS{
		backend:      backend,
		otpGenerator: otpGenerator,
		templates:    templates,
	}
}

func (s *SMS) SendVerificationCode(to string, locale string) (string, time.Time, error) {
	code, hash, err := s.otpGenerator()
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error generating OTP: %w", err)
	}

	body, err := s.templates.RenderSMS(locale, notifications.TemplateSMSData{
		Code: code,
	})
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error rendering email template: %w", err)
	}

	if err := s.backend.SendSMS(to, body); err != nil {
		return "", time.Time{}, fmt.Errorf("error sending SMS: %w", err)
	}

	return hash, time.Now().Add(in5Minutes), nil
}

func (s *SMS) CheckVerificationCode(to string, code string) error {
	return nil
}
