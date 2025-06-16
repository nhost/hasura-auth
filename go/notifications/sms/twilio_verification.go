package sms

import (
	"fmt"
	"strings"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioVerificationService struct {
	client                *twilio.RestClient
	isVerificationService bool
	from                  string
}

func NewTwilioVerificationService(
	accountSid string, authToken string, messageServiceSid string,
) *TwilioVerificationService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{ //nolint:exhaustruct
		Username: accountSid,
		Password: authToken,
	})
	return &TwilioVerificationService{
		client:                client,
		from:                  messageServiceSid,
		isVerificationService: strings.HasPrefix(messageServiceSid, "VA"),
	}
}

func (s *TwilioVerificationService) SendVerificationCode(
	to string, locale string,
) (string, time.Time, error) {
	if _, err := s.client.VerifyV2.CreateVerification(
		s.from,
		&openapi.CreateVerificationParams{ //nolint:exhaustruct
			To:      &to,
			Locale:  &locale,
			Channel: ptr("sms"),
		},
	); err != nil {
		return "", time.Time{}, fmt.Errorf("failed to fetch verification service: %w", err)
	}

	return "", time.Now().Add(in10Minutes), nil
}

func (s *TwilioVerificationService) CheckVerificationCode(to string, code string) (bool, error) {
	resp, err := s.client.VerifyV2.CreateVerificationCheck(
		s.from,
		&openapi.CreateVerificationCheckParams{ //nolint:exhaustruct
			To:   &to,
			Code: &code,
		},
	)
	if err != nil {
		return false, fmt.Errorf("failed to check verification code: %w", err)
	}

	return deptr(resp.Status) == "approved", nil
}
