package definitions

type PageService interface {
	Index(Request) PagesResponse
	Show(Request) PageResponse
	Create(Page) PageResponse
	Update(Page) PageResponse
	Delete(Request) PageResponse
}

// Page represents a web page to be scraped and downloaded
// model: true
type Page struct {
	Name       string `json:"name" bson:"name" grimoire:"index"`
	URL        string `json:"url" bson:"url"`
	Scraper    string `json:"scraper" bson:"scraper"`
	Downloader string `json:"downloader" bson:"downloader"`
}

type PagesResponse struct {
	Total   int64   `json:"total"`
	Results []*Page `json:"results"`
}
type PageResponse struct {
	Page *Page `json:"page"`
}
