package taipeitravel

import "encoding/json"

type AttractionsResponseDTO struct {
	Total int             `json:"total"`
	Data  []AttractionDTO `json:"data"`
}

type AttractionDTO struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	NameZh       *string `json:"name_zh"`
	OpenStatus   int     `json:"open_status"`
	Introduction string  `json:"introduction"`
	OpenTime     string  `json:"open_time"`
	Zipcode      string  `json:"zipcode"`
	District     string  `json:"distric"` // The third party misspells this as "distric"; do not "correct" it.
	Address      string  `json:"address"`
	Tel          string  `json:"tel"`
	Fax          string  `json:"fax"`
	Email        string  `json:"email"`
	Months       string  `json:"months"`
	NLat         float64 `json:"nlat"`
	ELong        float64 `json:"elong"`
	OfficialSite string  `json:"official_site"`
	Facebook     string  `json:"facebook"`
	Ticket       string  `json:"ticket"`
	Remind       string  `json:"remind"`
	StayTime     string  `json:"staytime"`
	Modified     string  `json:"modified"`
	URL          string  `json:"url"`

	Category []TagDTO   `json:"category"`
	Target   []TagDTO   `json:"target"`
	Friendly []TagDTO   `json:"friendly"`
	Images   []ImageDTO `json:"images"`
	Links    []LinkDTO  `json:"links"`

	// The sample data always uses empty arrays, and the real schema is still unknown.
	// Use RawMessage to preserve the payload as-is instead of guessing the structure and losing data.
	Service []json.RawMessage `json:"service"`
	Files   []json.RawMessage `json:"files"`
}

type TagDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ImageDTO struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
	Ext     string `json:"ext"`
}

type LinkDTO struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
}
