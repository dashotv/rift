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
	Title      string `json:"title"`
	Season     int    `json:"season"`
	Episode    int    `json:"episode"`
	Raw        string `json:"raw"`
	DisplayID  string `json:"display_id"`
	Extension  string `json:"extension"`
	Resolution int    `json:"resolution"`
	Size       int64  `json:"size"`
	Download   string `json:"download"`
	View       string `json:"view"`
	Source     string `json:"source"`
}
