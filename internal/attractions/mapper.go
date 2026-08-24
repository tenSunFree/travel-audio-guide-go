package attractions

import "github.com/tenSunFree/travel-audio-guide-go/internal/taipeitravel"

func fromTaipeiTravel(source taipeitravel.AttractionsResponseDTO) Response {
	items := make([]Attraction, 0, len(source.Data))
	for _, item := range source.Data {
		items = append(items, Attraction{
			ID:           item.ID,
			Name:         item.Name,
			NameZh:       item.NameZh,
			OpenStatus:   item.OpenStatus,
			Introduction: item.Introduction,
			OpenTime:     item.OpenTime,
			Zipcode:      item.Zipcode,
			District:     item.District,
			Address:      item.Address,
			Tel:          item.Tel,
			Fax:          item.Fax,
			Email:        item.Email,
			Months:       item.Months,
			NLat:         item.NLat,
			ELong:        item.ELong,
			OfficialSite: item.OfficialSite,
			Facebook:     item.Facebook,
			Ticket:       item.Ticket,
			Remind:       item.Remind,
			StayTime:     item.StayTime,
			Modified:     item.Modified,
			URL:          item.URL,

			Category: mapTags(item.Category),
			Target:   mapTags(item.Target),
			Friendly: mapTags(item.Friendly),
			Images:   mapImages(item.Images),
			Links:    mapLinks(item.Links),
			Service:  item.Service,
			Files:    item.Files,
		})
	}
	return Response{Total: source.Total, Data: items}
}

func mapTags(source []taipeitravel.TagDTO) []Tag {
	out := make([]Tag, 0, len(source))
	for _, t := range source {
		out = append(out, Tag{ID: t.ID, Name: t.Name})
	}
	return out
}

func mapImages(source []taipeitravel.ImageDTO) []Image {
	out := make([]Image, 0, len(source))
	for _, i := range source {
		out = append(out, Image{Src: i.Src, Subject: i.Subject, Ext: i.Ext})
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
