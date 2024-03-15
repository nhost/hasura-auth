package notifications

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/valyala/fasttemplate"
)

type TemplateName string

const (
	TemplateNameEmailVerify        TemplateName = "email-verify"
	TemplateNameEmailConfirmChange TemplateName = "email-confirm-change"
	TemplateNameSigninPasswordless TemplateName = "signin-passwordless"
	TemplateNamePasswordReset      TemplateName = "password-reset"
)

type Templates struct {
	templates     map[string]*fasttemplate.Template
	defaultLocale string
	logger        *slog.Logger
}

func NewTemplatesFromFilesystem(
	basePath string,
	defaultLocale string,
	logger *slog.Logger,
) (*Templates, error) {
	templates := make(map[string]*fasttemplate.Template)
	if err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relativePath, err := filepath.Rel(basePath, path)
			if err != nil {
				return fmt.Errorf("error getting relative path: %w", err)
			}

			if info.Name() == "body.html" || info.Name() == "body.txt" || info.Name() == "subject.txt" {
				f, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("error reading file: %w", err)
				}
				templates[relativePath] = fasttemplate.New(string(f), "${", "}")
			}
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("error walking the templates path (%s): %w", basePath, err)
	}

	return &Templates{
		templates:     templates,
		defaultLocale: defaultLocale,
		logger:        logger,
	}, nil
}

func (t *Templates) GetRawTemplates() map[string]*fasttemplate.Template {
	return t.templates
}

func (t *Templates) GetTemplate(
	templateName TemplateName, locale string,
) (
	*fasttemplate.Template, *fasttemplate.Template, error,
) {
	path := filepath.Join(locale, string(templateName), "body.html")
	template, ok := t.templates[path]
	if !ok {
		return nil, nil, ErrTemplateNotFound
	}

	path = filepath.Join(locale, string(templateName), "subject.txt")
	subject, ok := t.templates[path]
	if !ok {
		return nil, nil, ErrTemplateNotFound
	}

	return template, subject, nil
}

type TemplateData struct {
	Link        string
	DisplayName string
	Email       string
	NewEmail    string
	Ticket      string
	RedirectTo  string
	Locale      string
	ServerURL   string
	ClientURL   string
}

func (data TemplateData) ToMap(extra map[string]any) map[string]any {
	m := map[string]any{
		"link":        data.Link,
		"displayName": data.DisplayName,
		"email":       data.Email,
		"newEmail":    data.NewEmail,
		"ticket":      data.Ticket,
		"redirectTo":  data.RedirectTo,
		"locale":      data.Locale,
		"serverUrl":   data.ServerURL,
		"clientUrl":   data.ClientURL,
	}

	for k, v := range extra {
		m[k] = v
	}

	return m
}

func (t *Templates) Render(
	locale string,
	templateName TemplateName,
	data TemplateData,
) (string, string, error) {
	bodyTemplate, subjectTemplate, err := t.GetTemplate(templateName, locale)
	if errors.Is(err, ErrTemplateNotFound) {
		locale = t.defaultLocale
		t.logger.Warn("email-verify template not found, falling back to default locale",
			slog.String("locale", locale))
		bodyTemplate, subjectTemplate, err = t.GetTemplate("email-verify", locale)
	}
	if err != nil {
		return "", "", fmt.Errorf("error getting email template: %w", err)
	}

	m := data.ToMap(map[string]any{"locale": locale})
	body := bodyTemplate.ExecuteString(m)
	subject := subjectTemplate.ExecuteString(m)
	return body, subject, nil
}
