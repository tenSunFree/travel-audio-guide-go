package tours

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

// NewService creates a new Service for tours operations.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type ValidationError struct{ Message string }

// Error returns the error message for ValidationError.
func (e *ValidationError) Error() string { return e.Message }

var themeSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
}

// GetTheme retrieves tour themes for the specified language, validating the language parameter.
func (s *Service) GetTheme(ctx context.Context, lang string, rawQuery string) (ThemeResponse, error) {
	if !themeSupportedLanguages[lang] {
		return ThemeResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetTheme(ctx, lang, rawQuery)
}
