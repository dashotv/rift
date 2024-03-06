package definitions

type VideoService interface {
	Index(Request) Response
	Show(Request) Response
	Create(Video) Response
	Update(Video) Response
	Delete(Request) Response
}

// Video represents a video to be downloaded
// model: true
type Video struct {
	Title      string `json:"title" bson:"title" grimoire:"index"`
	Season     int    `json:"season" bson:"season" grimoire:"index"`
	Episode    int    `json:"episode" bson:"episode" grimoire:"index"`
	Raw        string `json:"raw" bson:"raw"`
	DisplayID  string `json:"display_id" bson:"display_id" grimoire:"index"`
	Extension  string `json:"extension" bson:"extension"`
	Resolution int    `json:"resolution" bson:"resolution"`
	Size       int64  `json:"size" bson:"size"`
	Download   string `json:"download" bson:"download"`
	View       string `json:"view" bson:"view"`
	Source     string `json:"source" bson:"source"`
}
