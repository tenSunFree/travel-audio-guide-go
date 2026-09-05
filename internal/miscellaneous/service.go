package miscellaneous

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

// NewService creates a new Service for miscellaneous operations.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type ValidationError struct{ Message string }

// Error returns the error message for ValidationError.
func (e *ValidationError) Error() string { return e.Message }

var categoriesSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
}

// categoriesSupportedTypes is intentionally restricted to type values whose
// upstream response shape has actually been verified. The upstream API
// documents six type values (Activity, Calendar, Pictorial, Attractions,
// Accommodation, Tours), but only "Tours" has been confirmed to use the
// {"data": {"Category": [...]}} wrapper this DTO expects. Enabling the
// other values here without first confirming their response shape would
// let unverified upstream formats silently decode into an empty/zero-value
// CategoriesResponse and still return 200 — see internal/taipeitravel/dto.go
// for the DTO this depends on.
//
// To add a new type: call the upstream endpoint with that type, confirm the
// wrapper key matches "Category" (or extend CategoriesDataDTO/mapper if it
// doesn't), then add it here.
var categoriesSupportedTypes = map[string]bool{
	"Tours": true,
}

// GetCategories retrieves categories for the specified language and type, validating both parameters.
func (s *Service) GetCategories(ctx context.Context, lang string, typeParam string, rawQuery string) (CategoriesResponse, error) {
	if !categoriesSupportedLanguages[lang] {
		return CategoriesResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	if !categoriesSupportedTypes[typeParam] {
		return CategoriesResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported or unverified type: %s", typeParam)}
	}
	return s.repo.GetCategories(ctx, lang, rawQuery)
}
