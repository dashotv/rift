package definitions

type PageService interface {
	Index(Request) Response
	Show(Request) Response
	Create(Page) Response
	Update(Page) Response
	Delete(Request) Response
}

// Page represents a web page to be scraped and downloaded
// model: true
type Page struct {
	Name       string `json:"name" bson:"name" grimoire:"index"`
	URL        string `json:"url" bson:"url"`
	Scraper    string `json:"scraper" bson:"scraper"`
	Downloader string `json:"downloader" bson:"downloader"`
}
