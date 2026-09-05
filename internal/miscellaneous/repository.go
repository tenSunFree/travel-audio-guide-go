package miscellaneous

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

// NewRepository creates a new Repository for miscellaneous data access.
func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

// GetCategories fetches categories from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetCategories(ctx context.Context, lang string, rawQuery string) (CategoriesResponse, error) {
	result, err := r.client.GetMiscellaneousCategories(ctx, lang, rawQuery)
	if err != nil {
		return CategoriesResponse{}, err
	}
	return fromTaipeiTravelCategories(result), nil
}
