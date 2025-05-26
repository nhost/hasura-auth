package notifications_test

import (
	"log/slog"
	"net/smtp"
	"testing"

	"github.com/nhost/hasura-auth/go/notifications"
)

func TestExtractEmail(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple email",
			input:    "user@example.com",
			expected: "user@example.com",
		},
		{
			name:     "email with display name",
			input:    "Test User <user@example.com>",
			expected: "user@example.com",
		},
		{
			name:     "malformed email - no closing bracket",
			input:    "Test User <user@example.com",
			expected: "Test User <user@example.com",
		},
		{
			name:     "email with display name containing brackets",
			input:    "Test (User) <user@example.com>",
			expected: "user@example.com",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Now that ExtractEmail is public, we can test it directly
			result := notifications.ExtractEmail(tc.input)
			if result != tc.expected {
				t.Errorf("ExtractEmail(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestEmailSend(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{
		{
			name:     "plain email addresses",
			from:     "admin@localhost",
			to:       "recipient@localhost",
			expected: true,
		},
		{
			name:     "formatted email addresses",
			from:     "Administrator <admin@localhost>",
			to:       "Recipient <recipient@localhost>",
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mail := notifications.NewEmail(
				"localhost",
				1025,
				false,
				smtp.PlainAuth("", "user", "password", "localhost"),
				tc.from,
				map[string]string{
					"x-something": "asd",
				},
				nil,
			)
			headers := map[string]string{
				"x-custom-header": "custom-value",
			}
			err := mail.Send(tc.to, "test", "contents", headers)

			if tc.expected && err != nil {
				t.Fatalf("expected success but got error: %v", err)
			} else if !tc.expected && err == nil {
				t.Fatalf("expected error but got success")
			}
		})
	}
}

func TestEmailSendEmailVerify(t *testing.T) {
	t.Parallel()

	logger := slog.Default()
	templates, err := notifications.NewTemplatesFromFilesystem(
		"../../email-templates/", "en", logger,
	)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	mail := notifications.NewEmail(
		"localhost",
		1025,
		false,
		smtp.PlainAuth("", "user", "password", "localhost"),
		"admin@localhost",
		map[string]string{
			"x-something": "asd",
		},
		templates,
	)

	cases := []struct {
		name   string
		data   notifications.TemplateData
		locale string
	}{
		{
			name: "",
			data: notifications.TemplateData{
				Link:        "http://link",
				DisplayName: "Display Name",
				Email:       "user@email",
				NewEmail:    "",
				Ticket:      "ticket",
				RedirectTo:  "http://redirect-to",
				Locale:      "en",
				ServerURL:   "http://servier-url",
				ClientURL:   "http://client-url",
			},
			locale: "en",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if err != nil {
				t.Fatalf("error creating mailer: %v", err)
			}

			if err := mail.SendEmail(
				t.Context(),
				"user@localhost",
				tc.locale,
				notifications.TemplateNameEmailVerify,
				tc.data,
			); err != nil {
				t.Fatalf("error sending email: %v", err)
			}
		})
	}
}
