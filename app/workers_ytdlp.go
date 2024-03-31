package app

import (
	"context"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/ytdlp"
)

type YtdlpList struct {
	minion.WorkerDefaults[*YtdlpList]
	Name string
	URL  string
}

func (j *YtdlpList) Kind() string { return "ytdlp_list" }
func (j *YtdlpList) Work(ctx context.Context, job *minion.Job[*YtdlpList]) error {
	// l := s.Logger.Named("ytdlp.list")
	name := job.Args.Name
	url := job.Args.URL

	list, err := ytdlp.ProcessURL(url)
	if err != nil {
		return fae.Errorf("ytdlp-list: %s %s: %s", name, url, err)
	}

	if len(list) == 0 {
		return fae.Errorf("ytdlp-list: %s %s: no entries", name, url)
	}

	for _, e := range list {
		// l.Warnf("ytdlp-list: %s", e.WebpageURL)
		if err := app.Workers.Enqueue(&YtdlpParse{Name: name, Source: "myanime", Info: e}); err != nil {
			return fae.Errorf("ytdlp-list: info: %s: %w", e.WebpageURL, err)
		}
	}

	return nil
}

type YtdlpParse struct {
	minion.WorkerDefaults[*YtdlpParse]
	Name   string
	Source string
	Info   *ytdlp.Info
}

func (j *YtdlpParse) Kind() string { return "ytdlp_parse" }
func (j *YtdlpParse) Work(ctx context.Context, job *minion.Job[*YtdlpParse]) error {
	// l := s.Logger.Named("ytdlp.parse")
	name := job.Args.Name
	source := job.Args.Source
	info := job.Args.Info

	// l.Warnf("%s %d %s [%s] URL:%s", info.Fulltitle, info.Height, info.EXT, info.DisplayID, info.WebpageURL)

	count, err := app.DB.Video.Query().Where("display_id", info.DisplayID).Count()
	if err != nil {
		return fae.Errorf("couting: %w", err)
	}
	if count > 0 {
		return nil
	}

	_, season, episode, err := ParseFulltitle(info.Fulltitle)
	if err != nil {
		return fae.Errorf("parsing: %w", err)
	}

	video := &Video{}
	video.Title = name
	video.Season = season
	video.Episode = episode
	video.Raw = info.Fulltitle
	video.Resolution = int(info.Height)
	video.Extension = info.EXT
	video.DisplayId = info.DisplayID
	video.Download = info.WebpageURL
	video.View = info.WebpageURL
	video.Size = info.FilesizeApprox
	video.Source = source

	if err := app.DB.Video.Save(video); err != nil {
		return fae.Errorf("saving: %w", err)
	}

	return nil
}
