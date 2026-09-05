package media

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

// NewService creates a new Service for media operations.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type ValidationError struct{ Message string }

// Error returns the error message for ValidationError.
func (e *ValidationError) Error() string { return e.Message }

var audioSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
}

// GetAudio retrieves audio media for the specified language, validating the language parameter.
func (s *Service) GetAudio(ctx context.Context, lang string, rawQuery string) (AudioResponse, error) {
	if !audioSupportedLanguages[lang] {
		return AudioResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetAudio(ctx, lang, rawQuery)
}
