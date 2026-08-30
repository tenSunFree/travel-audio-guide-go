package tours

import "encoding/json"

type ThemeResponse struct {
	Total int     `json:"total"`
	Data  []Theme `json:"data"`
}

type Theme struct {
	ID          int             `json:"id"`
	Seasons     []string        `json:"seasons"`
	Months      []string        `json:"months"`
	Days        int             `json:"days"`
	Title       string          `json:"title"`
	Author      string          `json:"author"`
	Description string          `json:"description"`
	Consume     string          `json:"consume"`
	Remark      string          `json:"remark"`
	Note        string          `json:"note"`
	URL         string          `json:"url"`
	Category    json.RawMessage `json:"category"`
	Transport   json.RawMessage `json:"transport"`
	Users       json.RawMessage `json:"users"`
	Modified    string          `json:"modified"`
	Files       []File          `json:"files"`
}

type File struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
	Ext     string `json:"ext"`
}
