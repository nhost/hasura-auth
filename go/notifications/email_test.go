package notifications_test

import (
	"context"
	"log/slog"
	"net/smtp"
	"os"
	"strconv"
	"testing"

	"github.com/nhost/hasura-auth/go/notifications"
)

func TestEmailSend(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
	}{
		{
			name: "",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// tc := tc

			portStr := os.Getenv("AUTH_SMTP_PORT")
			port, err := strconv.Atoi(portStr)
			if err != nil {
				t.Fatalf("error parsing port: %v", err)
			}

			auth := notifications.PlainAuth(
				"",
				os.Getenv("AUTH_SMTP_USER"),
				os.Getenv("AUTH_SMTP_PASS"),
				os.Getenv("AUTH_SMTP_HOST"),
			)

			useTLSConnection := os.Getenv("AUTH_SMTP_SECURE") == "true"

			mail := notifications.NewEmail(
				os.Getenv("AUTH_SMTP_HOST"),
				uint16(port),
				useTLSConnection,
				auth,
				os.Getenv("AUTH_SMTP_SENDER"),
				map[string]string{
					"x-something": "asd",
				},
				nil,
			)
			headers := map[string]string{
				"x-another": "qwe",
			}
			if err := mail.Send("user@localhost", "some email", "contents", headers); err != nil {
				t.Fatalf("error sending email: %v", err)
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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// tc := tc

			if err != nil {
				t.Fatalf("error creating mailer: %v", err)
			}

			if err := mail.SendEmail(
				context.Background(),
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
