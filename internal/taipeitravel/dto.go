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

// ---------- Events / News ----------

type NewsResponseDTO struct {
	Total int       `json:"total"`
	Data  []NewsDTO `json:"data"`
}

type NewsDTO struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Begin       *string    `json:"begin"`
	End         *string    `json:"end"`
	Posted      string     `json:"posted"`
	Modified    string     `json:"modified"`
	URL         string     `json:"url"`
	Files       []ImageDTO `json:"files"`
	Links       []LinkDTO  `json:"links"`
}

// ---------- Events / Activity ----------

type ActivityResponseDTO struct {
	Total int           `json:"total"`
	Data  []ActivityDTO `json:"data"`
}

type ActivityDTO struct {
	District    string     `json:"distric"`
	Address     string     `json:"address"`
	Nlat        string     `json:"nlat"` // upstream sends this as a string, not a number
	Elong       string     `json:"elong"`
	Organizer   string     `json:"organizer"`
	CoOrganizer string     `json:"co_rganizer"` // upstream field name, kept as-is
	Contact     string     `json:"contact"`
	Tel         string     `json:"tel"`
	Fax         string     `json:"fax"`
	Ticket      string     `json:"ticket"`
	Traffic     string     `json:"traffic"`
	Parking     string     `json:"parking"`
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Begin       *string    `json:"begin"`
	End         *string    `json:"end"`
	Posted      string     `json:"posted"`
	Modified    string     `json:"modified"`
	URL         string     `json:"url"`
	Files       []ImageDTO `json:"files"`
	Links       []LinkDTO  `json:"links"`
}

// ---------- Events / Calendar ----------

type CalendarResponseDTO struct {
	Total int           `json:"total"`
	Data  []CalendarDTO `json:"data"`
}

type CalendarDTO struct {
	District    string     `json:"distric"`
	Address     string     `json:"address"`
	Nlat        string     `json:"nlat"`
	Elong       string     `json:"elong"`
	Tel         string     `json:"tel"`
	Fax         string     `json:"fax"`
	Ticket      string     `json:"ticket"`
	Traffic     string     `json:"traffic"`
	Parking     string     `json:"parking"`
	IsMajor     bool       `json:"is_major"`
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Begin       *string    `json:"begin"`
	End         *string    `json:"end"`
	Posted      string     `json:"posted"`
	Modified    string     `json:"modified"`
	URL         string     `json:"url"`
	Files       []ImageDTO `json:"files"`
	Links       []LinkDTO  `json:"links"`
}

// ---------- Miscellaneous / Categories ----------

type CategoriesResponseDTO struct {
	Total int               `json:"total"`
	Data  CategoriesDataDTO `json:"data"`
}

// Confirmed via type=Tours sample: the wrapper key is "Category" (singular,
// capitalized) regardless of the "type" query value. This has only been
// verified for type=Tours — if other type values (Activity, Calendar,
// Pictorial, Attractions, Accommodation) turn out to use a different key,
// this struct will need per-type handling.
type CategoriesDataDTO struct {
	Category []TagDTO `json:"Category"`
}

// ---------- Media / Audio ----------

type AudioResponseDTO struct {
	Total int        `json:"total"`
	Data  []AudioDTO `json:"data"`
}

type AudioDTO struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Summary  *string `json:"summary"`
	URL      string  `json:"url"`
	FileExt  *string `json:"file_ext"`
	Modified string  `json:"modified"`
}

// ---------- Tours / Theme ----------

type TourThemeResponseDTO struct {
	Total int            `json:"total"`
	Data  []TourThemeDTO `json:"data"`
}

type TourThemeDTO struct {
	ID          int      `json:"id"`
	Seasons     []string `json:"seasons"`
	Months      []string `json:"months"`
	Days        int      `json:"days"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Consume     string   `json:"consume"`
	Remark      string   `json:"remark"`
	Note        string   `json:"note"`
	URL         string   `json:"url"`
	// The sample data always returns null for these three fields and the
	// real schema is unknown. Use RawMessage to preserve the payload as-is.
	Category  json.RawMessage `json:"category"`
	Transport json.RawMessage `json:"transport"`
	Users     json.RawMessage `json:"users"`
	Modified  string          `json:"modified"`
	Files     []ImageDTO      `json:"files"`
}
