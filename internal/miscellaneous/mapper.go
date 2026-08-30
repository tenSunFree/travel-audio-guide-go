package miscellaneous

import "github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"

func fromTaipeiTravelCategories(source taipeitravel.CategoriesResponseDTO) CategoriesResponse {
	items := make([]Category, 0, len(source.Data.Category))
	for _, item := range source.Data.Category {
		items = append(items, Category{ID: item.ID, Name: item.Name})
	}
	return CategoriesResponse{
		Total: source.Total,
		Data:  CategoriesData{Category: items},
	}
}
