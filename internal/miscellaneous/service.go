package miscellaneous

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type ValidationError struct{ Message string }

func (e *ValidationError) Error() string { return e.Message }

var categoriesSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
}

var categoriesSupportedTypes = map[string]bool{
	"Activity": true, "Calendar": true, "Pictorial": true,
	"Attractions": true, "Accommodation": true, "Tours": true,
}

func (s *Service) GetCategories(ctx context.Context, lang string, typeParam string, rawQuery string) (CategoriesResponse, error) {
	if !categoriesSupportedLanguages[lang] {
		return CategoriesResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	if !categoriesSupportedTypes[typeParam] {
		return CategoriesResponse{}, &ValidationError{Message: fmt.Sprintf("invalid parameter: type (%s)", typeParam)}
	}
	return s.repo.GetCategories(ctx, lang, rawQuery)
}
