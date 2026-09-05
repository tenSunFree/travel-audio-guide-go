package miscellaneous

type CategoriesResponse struct {
	Total int            `json:"total"`
	Data  CategoriesData `json:"data"`
}

type CategoriesData struct {
	Category []Category `json:"Category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
