package events

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

var newsSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
}

var activitySupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true, "ja": true, "ko": true,
	"es": true, "id": true, "th": true, "vi": true,
}

var calendarSupportedLanguages = map[string]bool{
	"zh-tw": true, "zh-cn": true, "en": true,
}

func (s *Service) GetNews(ctx context.Context, lang string, rawQuery string) (NewsResponse, error) {
	if !newsSupportedLanguages[lang] {
		return NewsResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetNews(ctx, lang, rawQuery)
}

func (s *Service) GetActivity(ctx context.Context, lang string, rawQuery string) (ActivityResponse, error) {
	if !activitySupportedLanguages[lang] {
		return ActivityResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetActivity(ctx, lang, rawQuery)
}

func (s *Service) GetCalendar(ctx context.Context, lang string, rawQuery string) (CalendarResponse, error) {
	if !calendarSupportedLanguages[lang] {
		return CalendarResponse{}, &ValidationError{Message: fmt.Sprintf("unsupported language: %s", lang)}
	}
	return s.repo.GetCalendar(ctx, lang, rawQuery)
}
