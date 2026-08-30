package media

import "github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"

func fromTaipeiTravelAudio(source taipeitravel.AudioResponseDTO) AudioResponse {
	items := make([]Audio, 0, len(source.Data))
	for _, item := range source.Data {
		items = append(items, Audio{
			ID:       item.ID,
			Title:    item.Title,
			Summary:  item.Summary,
			URL:      item.URL,
			FileExt:  item.FileExt,
			Modified: item.Modified,
		})
	}
	return AudioResponse{Total: source.Total, Data: items}
}
