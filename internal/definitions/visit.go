package definitions

type VisitService interface {
	Index(Request) Response
	Show(Request) Response
	Create(Visit) Response
	Update(Visit) Response
	Delete(Request) Response
}

// Visit represents a web page to be scraped and downloaded
// model: true
type Visit struct {
	PageID string `json:"page_id"`
	URL    string `json:"url"`
}
