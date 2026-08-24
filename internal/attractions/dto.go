package attractions

import "encoding/json"

type Response struct {
	Total int          `json:"total"`
	Data  []Attraction `json:"data"`
}

type Attraction struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	NameZh       *string `json:"name_zh"`
	OpenStatus   int     `json:"open_status"`
	Introduction string  `json:"introduction"`
	OpenTime     string  `json:"open_time"`
	Zipcode      string  `json:"zipcode"`
	District     string  `json:"distric"`
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

	Category []Tag   `json:"category"`
	Target   []Tag   `json:"target"`
	Friendly []Tag   `json:"friendly"`
	Images   []Image `json:"images"`
	Links    []Link  `json:"links"`

	Service []json.RawMessage `json:"service"`
	Files   []json.RawMessage `json:"files"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
	Ext     string `json:"ext"`
}

type Link struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
}
