package media

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"
)

type Repository struct {
	client *taipeitravel.Client
}

// NewRepository creates a new Repository for media data access.
func NewRepository(client *taipeitravel.Client) *Repository {
	return &Repository{client: client}
}

// GetAudio fetches audio media from the Taipei Travel API and maps it to the response format.
func (r *Repository) GetAudio(ctx context.Context, lang string, rawQuery string) (AudioResponse, error) {
	result, err := r.client.GetMediaAudio(ctx, lang, rawQuery)
	if err != nil {
		return AudioResponse{}, err
	}
	return fromTaipeiTravelAudio(result), nil
}
