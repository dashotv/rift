package ytdlp

type List struct {
	ID                 string      `json:"id"`
	Title              string      `json:"title"`
	Timestamp          float64     `json:"timestamp"`
	Description        string      `json:"description"`
	Thumbnail          string      `json:"thumbnail"`
	AgeLimit           int64       `json:"age_limit"`
	Type               string      `json:"_type"`
	Entries            []Entry     `json:"entries"`
	WebpageURL         string      `json:"webpage_url"`
	OriginalURL        string      `json:"original_url"`
	WebpageURLBasename string      `json:"webpage_url_basename"`
	WebpageURLDomain   string      `json:"webpage_url_domain"`
	Extractor          string      `json:"extractor"`
	ExtractorKey       string      `json:"extractor_key"`
	UploadDate         string      `json:"upload_date"`
	Thumbnails         []Thumbnail `json:"thumbnails"`
	PlaylistCount      int64       `json:"playlist_count"`
	Epoch              int64       `json:"epoch"`
	Version            Version     `json:"_version"`
}

type Thumbnail struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

type Version struct {
	Version        string      `json:"version"`
	CurrentGitHead interface{} `json:"current_git_head"`
	ReleaseGitHead string      `json:"release_git_head"`
	Repository     string      `json:"repository"`
}

type Info struct {
	ID                 string      `json:"id"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Thumbnails         []Thumbnail `json:"thumbnails"`
	Duration           int64       `json:"duration"`
	Timestamp          int64       `json:"timestamp"`
	Uploader           string      `json:"uploader"`
	UploaderID         string      `json:"uploader_id"`
	AgeLimit           int64       `json:"age_limit"`
	ViewCount          int64       `json:"view_count"`
	LikeCount          int64       `json:"like_count"`
	Formats            []Format    `json:"formats"`
	WebpageURL         string      `json:"webpage_url"`
	OriginalURL        string      `json:"original_url"`
	WebpageURLBasename string      `json:"webpage_url_basename"`
	WebpageURLDomain   string      `json:"webpage_url_domain"`
	Extractor          string      `json:"extractor"`
	ExtractorKey       string      `json:"extractor_key"`
	Thumbnail          string      `json:"thumbnail"`
	DisplayID          string      `json:"display_id"`
	Fulltitle          string      `json:"fulltitle"`
	DurationString     string      `json:"duration_string"`
	UploadDate         string      `json:"upload_date"`
	RequestedDownloads []Format    `json:"requested_downloads"`
	FormatID           string      `json:"format_id"`
	URL                string      `json:"url"`
	Tbr                float64     `json:"tbr"`
	EXT                string      `json:"ext"`
	Protocol           string      `json:"protocol"`
	Width              int64       `json:"width"`
	Height             int64       `json:"height"`
	Vcodec             string      `json:"vcodec"`
	Acodec             string      `json:"acodec"`
	DynamicRange       string      `json:"dynamic_range"`
	Resolution         string      `json:"resolution"`
	AspectRatio        float64     `json:"aspect_ratio"`
	FilesizeApprox     int64       `json:"filesize_approx"`
	VideoEXT           string      `json:"video_ext"`
	AudioEXT           string      `json:"audio_ext"`
	Vbr                float64     `json:"vbr"`
	ABR                float64     `json:"abr"`
	Format             string      `json:"format"`
	Epoch              int64       `json:"epoch"`
	Type               string      `json:"_type"`
	Version            Version     `json:"_version"`
}

type Entry struct {
	IeKey        string `json:"ie_key"`
	Type         string `json:"_type"`
	URL          string `json:"url"`
	Extractor    string `json:"extractor"`
	ExtractorKey string `json:"extractor_key"`
}
type Format struct {
	FormatID             string      `json:"format_id"`
	FormatIndex          interface{} `json:"format_index,omitempty"`
	URL                  string      `json:"url"`
	ManifestURL          *string     `json:"manifest_url,omitempty"`
	Tbr                  float64     `json:"tbr"`
	EXT                  string      `json:"ext"`
	FPS                  interface{} `json:"fps,omitempty"`
	Protocol             string      `json:"protocol"`
	Preference           interface{} `json:"preference,omitempty"`
	Quality              interface{} `json:"quality,omitempty"`
	HasDRM               interface{} `json:"has_drm,omitempty"`
	Width                int64       `json:"width"`
	Height               int64       `json:"height"`
	Vcodec               string      `json:"vcodec"`
	Acodec               string      `json:"acodec"`
	DynamicRange         string      `json:"dynamic_range"`
	Resolution           string      `json:"resolution"`
	AspectRatio          float64     `json:"aspect_ratio"`
	FilesizeApprox       int64       `json:"filesize_approx"`
	VideoEXT             string      `json:"video_ext"`
	AudioEXT             string      `json:"audio_ext"`
	Vbr                  float64     `json:"vbr"`
	ABR                  float64     `json:"abr"`
	Format               string      `json:"format"`
	Epoch                *int64      `json:"epoch,omitempty"`
	Filename             *string     `json:"_filename,omitempty"`
	WriteDownloadArchive *bool       `json:"__write_download_archive,omitempty"`
}
