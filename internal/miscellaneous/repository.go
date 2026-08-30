package miscellaneous

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

func (r *Repository) GetCategories(ctx context.Context, lang string, rawQuery string) (CategoriesResponse, error) {
	result, err := r.client.GetMiscellaneousCategories(ctx, lang, rawQuery)
	if err != nil {
		return CategoriesResponse{}, err
	}
	return fromTaipeiTravelCategories(result), nil
}
