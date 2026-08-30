package media

type AudioResponse struct {
	Total int     `json:"total"`
	Data  []Audio `json:"data"`
}

type Audio struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Summary  *string `json:"summary"`
	URL      string  `json:"url"`
	FileExt  *string `json:"file_ext"`
	Modified string  `json:"modified"`
}
