package tours

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

// NewRepository creates a new Repository for tours data access.
func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

// GetTheme fetches tour themes from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetTheme(ctx context.Context, lang string, rawQuery string) (ThemeResponse, error) {
	result, err := r.client.GetToursTheme(ctx, lang, rawQuery)
	if err != nil {
		return ThemeResponse{}, err
	}
	return fromTaipeiTravelTheme(result), nil
}
