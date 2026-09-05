package attractions

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

// NewRepository creates a new Repository for attractions data access.
func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

// GetAll fetches attractions data from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetAll(ctx context.Context, lang string, rawQuery string) (Response, error) {
	result, err := r.client.GetAttractions(ctx, lang, rawQuery)
	if err != nil {
		return Response{}, err
	}
	return fromTaipeiTravel(result), nil
}
