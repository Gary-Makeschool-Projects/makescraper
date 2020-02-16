package site

// Site structure
type Site struct {
	URL string `json:"url" form:"url" query:"url"`
}
