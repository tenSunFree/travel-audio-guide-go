package events

type NewsResponse struct {
	Total int    `json:"total"`
	Data  []News `json:"data"`
}

type News struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Begin       *string `json:"begin"`
	End         *string `json:"end"`
	Posted      string  `json:"posted"`
	Modified    string  `json:"modified"`
	URL         string  `json:"url"`
	Files       []File  `json:"files"`
	Links       []Link  `json:"links"`
}

type ActivityResponse struct {
	Total int        `json:"total"`
	Data  []Activity `json:"data"`
}

type Activity struct {
	District    string  `json:"distric"`
	Address     string  `json:"address"`
	Nlat        string  `json:"nlat"`
	Elong       string  `json:"elong"`
	Organizer   string  `json:"organizer"`
	CoOrganizer string  `json:"co_rganizer"`
	Contact     string  `json:"contact"`
	Tel         string  `json:"tel"`
	Fax         string  `json:"fax"`
	Ticket      string  `json:"ticket"`
	Traffic     string  `json:"traffic"`
	Parking     string  `json:"parking"`
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Begin       *string `json:"begin"`
	End         *string `json:"end"`
	Posted      string  `json:"posted"`
	Modified    string  `json:"modified"`
	URL         string  `json:"url"`
	Files       []File  `json:"files"`
	Links       []Link  `json:"links"`
}

type CalendarResponse struct {
	Total int        `json:"total"`
	Data  []Calendar `json:"data"`
}

type Calendar struct {
	District    string  `json:"distric"`
	Address     string  `json:"address"`
	Nlat        string  `json:"nlat"`
	Elong       string  `json:"elong"`
	Tel         string  `json:"tel"`
	Fax         string  `json:"fax"`
	Ticket      string  `json:"ticket"`
	Traffic     string  `json:"traffic"`
	Parking     string  `json:"parking"`
	IsMajor     bool    `json:"is_major"`
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Begin       *string `json:"begin"`
	End         *string `json:"end"`
	Posted      string  `json:"posted"`
	Modified    string  `json:"modified"`
	URL         string  `json:"url"`
	Files       []File  `json:"files"`
	Links       []Link  `json:"links"`
}

type File struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
	Ext     string `json:"ext"`
}

type Link struct {
	Src     string `json:"src"`
	Subject string `json:"subject"`
}
