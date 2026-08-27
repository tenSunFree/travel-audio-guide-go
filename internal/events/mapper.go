package events

import "github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"

func fromTaipeiTravelNews(source taipeitravel.NewsResponseDTO) NewsResponse {
	items := make([]News, 0, len(source.Data))
	for _, item := range source.Data {
		items = append(items, News{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Begin:       item.Begin,
			End:         item.End,
			Posted:      item.Posted,
			Modified:    item.Modified,
			URL:         item.URL,
			Files:       mapFiles(item.Files),
			Links:       mapLinks(item.Links),
		})
	}
	return NewsResponse{Total: source.Total, Data: items}
}

func fromTaipeiTravelActivity(source taipeitravel.ActivityResponseDTO) ActivityResponse {
	items := make([]Activity, 0, len(source.Data))
	for _, item := range source.Data {
		items = append(items, Activity{
			District:    item.District,
			Address:     item.Address,
			Nlat:        item.Nlat,
			Elong:       item.Elong,
			Organizer:   item.Organizer,
			CoOrganizer: item.CoOrganizer,
			Contact:     item.Contact,
			Tel:         item.Tel,
			Fax:         item.Fax,
			Ticket:      item.Ticket,
			Traffic:     item.Traffic,
			Parking:     item.Parking,
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Begin:       item.Begin,
			End:         item.End,
			Posted:      item.Posted,
			Modified:    item.Modified,
			URL:         item.URL,
			Files:       mapFiles(item.Files),
			Links:       mapLinks(item.Links),
		})
	}
	return ActivityResponse{Total: source.Total, Data: items}
}

func fromTaipeiTravelCalendar(source taipeitravel.CalendarResponseDTO) CalendarResponse {
	items := make([]Calendar, 0, len(source.Data))
	for _, item := range source.Data {
		items = append(items, Calendar{
			District:    item.District,
			Address:     item.Address,
			Nlat:        item.Nlat,
			Elong:       item.Elong,
			Tel:         item.Tel,
			Fax:         item.Fax,
			Ticket:      item.Ticket,
			Traffic:     item.Traffic,
			Parking:     item.Parking,
			IsMajor:     item.IsMajor,
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Begin:       item.Begin,
			End:         item.End,
			Posted:      item.Posted,
			Modified:    item.Modified,
			URL:         item.URL,
			Files:       mapFiles(item.Files),
			Links:       mapLinks(item.Links),
		})
	}
	return CalendarResponse{Total: source.Total, Data: items}
}

func mapFiles(source []taipeitravel.ImageDTO) []File {
	out := make([]File, 0, len(source))
	for _, f := range source {
		out = append(out, File{Src: f.Src, Subject: f.Subject, Ext: f.Ext})
	}
	return out
}

func mapLinks(source []taipeitravel.LinkDTO) []Link {
	out := make([]Link, 0, len(source))
	for _, l := range source {
		out = append(out, Link{Src: l.Src, Subject: l.Subject})
	}
	return out
}
