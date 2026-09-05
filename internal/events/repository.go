package events

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

// NewRepository creates a new Repository for events data access.
func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

// GetNews fetches event news from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetNews(ctx context.Context, lang string, rawQuery string) (NewsResponse, error) {
	result, err := r.client.GetEventsNews(ctx, lang, rawQuery)
	if err != nil {
		return NewsResponse{}, err
	}
	return fromTaipeiTravelNews(result), nil
}

// GetActivity fetches event activities from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetActivity(ctx context.Context, lang string, rawQuery string) (ActivityResponse, error) {
	result, err := r.client.GetEventsActivity(ctx, lang, rawQuery)
	if err != nil {
		return ActivityResponse{}, err
	}
	return fromTaipeiTravelActivity(result), nil
}

// GetCalendar fetches event calendar from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetCalendar(ctx context.Context, lang string, rawQuery string) (CalendarResponse, error) {
	result, err := r.client.GetEventsCalendar(ctx, lang, rawQuery)
	if err != nil {
		return CalendarResponse{}, err
	}
	return fromTaipeiTravelCalendar(result), nil
}
