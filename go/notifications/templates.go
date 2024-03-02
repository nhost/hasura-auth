package notifications

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/valyala/fasttemplate"
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
	templateName string, locale string,
) (
	*fasttemplate.Template, *fasttemplate.Template, error,
) {
	path := filepath.Join(locale, templateName, "body.html")
	template, ok := t.templates[path]
	if !ok {
		return nil, nil, ErrTemplateNotFound
	}

	path = filepath.Join(locale, templateName, "subject.txt")
	subject, ok := t.templates[path]
	if !ok {
		return nil, nil, ErrTemplateNotFound
	}

	return template, subject, nil
}

type EmailVerifyData struct {
	Link        string
	DisplayName string
	Email       string
	Ticket      string
	RedirectTo  string
	ServerURL   string
	ClientURL   string
}

func (data EmailVerifyData) ToMap(extra map[string]any) map[string]any {
	m := map[string]any{
		"link":        data.Link,
		"displayName": data.DisplayName,
		"email":       data.Email,
		"ticket":      data.Ticket,
		"redirectTo":  data.RedirectTo,
		"serverUrl":   data.ServerURL,
		"clientUrl":   data.ClientURL,
	}

	for k, v := range extra {
		m[k] = v
	}

	return m
}

func (t *Templates) RenderEmailVerify(
	locale string,
	data EmailVerifyData,
) (string, string, error) {
	bodyTemplate, subjectTemplate, err := t.GetTemplate("email-verify", locale)
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