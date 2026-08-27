package events

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) GetNews(ctx context.Context, lang string, rawQuery string) (NewsResponse, error) {
	result, err := r.client.GetEventsNews(ctx, lang, rawQuery)
	if err != nil {
		return NewsResponse{}, err
	}
	return fromTaipeiTravelNews(result), nil
}

func (r *Repository) GetActivity(ctx context.Context, lang string, rawQuery string) (ActivityResponse, error) {
	result, err := r.client.GetEventsActivity(ctx, lang, rawQuery)
	if err != nil {
		return ActivityResponse{}, err
	}
	return fromTaipeiTravelActivity(result), nil
}

func (r *Repository) GetCalendar(ctx context.Context, lang string, rawQuery string) (CalendarResponse, error) {
	result, err := r.client.GetEventsCalendar(ctx, lang, rawQuery)
	if err != nil {
		return CalendarResponse{}, err
	}
	return fromTaipeiTravelCalendar(result), nil
}
