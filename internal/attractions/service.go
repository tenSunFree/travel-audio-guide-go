package attractions

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

// NewService creates a new Service for attractions operations.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

var supportedLanguages = map[string]bool{
	"zh-tw": true,
	"zh-cn": true,
	"en":    true,
	"ja":    true,
	"ko":    true,
}

type ValidationError struct{ Message string }

// Error returns the error message for ValidationError.
func (e *ValidationError) Error() string { return e.Message }

// GetAll retrieves all attractions for the specified language, validating the language parameter.
func (s *Service) GetAll(ctx context.Context, lang string, rawQuery string) (Response, error) {
	if !supportedLanguages[lang] {
		return Response{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetAll(ctx, lang, rawQuery)
}
