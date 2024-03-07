package definitions

import "go.mongodb.org/mongo-driver/bson/primitive"

type VisitService interface {
	Index(Request) VisitsResponse
	Show(Request) VisitResponse
	Create(Visit) VisitResponse
	Update(Visit) VisitResponse
	Delete(Request) VisitResponse
}

// Visit represents a web page to be scraped and downloaded
// model: true
type Visit struct {
	PageID primitive.ObjectID `json:"page_id" bson:"page_id" grimoire:"index"`
	URL    string             `json:"url" bson:"url"`
}

type VisitsResponse struct {
	Total   int64    `json:"total"`
	Results []*Visit `json:"results"`
}

type VisitResponse struct {
	Visit *Visit `json:"visit"`
}
