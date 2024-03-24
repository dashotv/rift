package server

import (
	"context"
	"fmt"

	"github.com/dashotv/minion"
)

type YtdlpListJob struct {
	minion.WorkerDefaults[*YtdlpListJob]
	Name string
	URL  string
}

func (j *YtdlpListJob) Kind() string { return "ytdlp_list" }
func (j *YtdlpListJob) Work(ctx context.Context, job *minion.Job[*YtdlpListJob]) error {
	s := getServer(ctx)
	// l := s.Logger.Named("ytdlp.list")
	name := job.Args.Name
	url := job.Args.URL

	list, err := ProcessURL(url)
	if err != nil {
		return fmt.Errorf("ytdlp-list: %s %s: %s", name, url, err)
	}

	if len(list) == 0 {
		return fmt.Errorf("ytdlp-list: %s %s: no entries", name, url)
	}

	for _, e := range list {
		// l.Warnf("ytdlp-list: %s", e.WebpageURL)
		if err := s.bg.Enqueue(&YtdlpParseJob{Name: name, Source: "myanime", Info: e}); err != nil {
			return fmt.Errorf("ytdlp-list: info: %s: %w", e.WebpageURL, err)
		}
	}

	return nil
}

type YtdlpParseJob struct {
	minion.WorkerDefaults[*YtdlpParseJob]
	Name   string
	Source string
	Info   *YtdlpInfo
}

func (j *YtdlpParseJob) Kind() string { return "ytdlp_parse" }
func (j *YtdlpParseJob) Work(ctx context.Context, job *minion.Job[*YtdlpParseJob]) error {
	s := getServer(ctx)
	// l := s.Logger.Named("ytdlp.parse")
	name := job.Args.Name
	source := job.Args.Source
	info := job.Args.Info

	// l.Warnf("%s %d %s [%s] URL:%s", info.Fulltitle, info.Height, info.EXT, info.DisplayID, info.WebpageURL)

	count, err := s.db.Video.Query().Where("display_id", info.DisplayID).Count()
	if err != nil {
		return fmt.Errorf("couting: %w", err)
	}
	if count > 0 {
		return nil
	}

	_, season, episode, err := ParseFulltitle(info.Fulltitle)
	if err != nil {
		return fmt.Errorf("parsing: %w", err)
	}

	video := &Video{}
	video.Title = name
	video.Season = season
	video.Episode = episode
	video.Raw = info.Fulltitle
	video.Resolution = int(info.Height)
	video.Extension = info.EXT
	video.DisplayID = info.DisplayID
	video.Download = info.WebpageURL
	video.View = info.WebpageURL
	video.Size = info.FilesizeApprox
	video.Source = source

	if err := s.db.Video.Save(video); err != nil {
		return fmt.Errorf("saving: %w", err)
	}

	return nil
}

type YtdlpList struct {
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

type YtdlpInfo struct {
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
