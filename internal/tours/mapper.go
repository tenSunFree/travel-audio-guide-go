package tours

import "github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"

func fromTaipeiTravelTheme(source taipeitravel.TourThemeResponseDTO) ThemeResponse {
	var items []Theme
	if source.Data != nil {
		items = make([]Theme, 0, len(source.Data))
		for _, item := range source.Data {
			items = append(items, Theme{
				ID:          item.ID,
				Seasons:     item.Seasons,
				Months:      item.Months,
				Days:        item.Days,
				Title:       item.Title,
				Author:      item.Author,
				Description: item.Description,
				Consume:     item.Consume,
				Remark:      item.Remark,
				Note:        item.Note,
				URL:         item.URL,
				Category:    item.Category,
				Transport:   item.Transport,
				Users:       item.Users,
				Modified:    item.Modified,
				Files:       mapFiles(item.Files),
			})
		}
	}
	return ThemeResponse{Total: source.Total, Data: items}
}

func mapFiles(source []taipeitravel.ImageDTO) []File {
	if source == nil {
		return nil
	}
	out := make([]File, 0, len(source))
	for _, f := range source {
		out = append(out, File{Src: f.Src, Subject: f.Subject, Ext: f.Ext})
	}
	return out
}
