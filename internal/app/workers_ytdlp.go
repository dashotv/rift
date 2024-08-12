package app

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/ytdlp"
)

type YtdlpList struct {
	minion.WorkerDefaults[*YtdlpList]
	PageID  primitive.ObjectID
	Name    string
	Season  int
	Episode int
	URL     string
}

func (j *YtdlpList) Kind() string { return "ytdlp_list" }
func (j *YtdlpList) Work(ctx context.Context, job *minion.Job[*YtdlpList]) (err error) {
	a := ContextApp(ctx)
	l := a.Log.Named("ytdlp.list")
	pid := job.Args.PageID
	name := job.Args.Name
	season := job.Args.Season
	episode := job.Args.Episode
	url := job.Args.URL

	// defer func() {
	// 	if err != nil {
	// 		if e := app.DB.Visit.Save()
	// 	}
	// }()

	list, err := ytdlp.ProcessURL(url)
	if err != nil {
		return fae.Wrapf(err, "ytdlp-list: %s %s", name, url)
	}

	if len(list) == 0 {
		return fae.Errorf("ytdlp-list: %s %s: no entries", name, url)
	}

	for _, e := range list {
		l.Warnf("ytdlp-list: %s", e.WebpageURL)
		if err := app.Workers.Enqueue(&YtdlpParse{Name: name, PageID: pid, URL: url, Info: e, Season: season, Episode: episode}); err != nil {
			return fae.Wrapf(err, "ytdlp-list: info: %s", e.WebpageURL)
		}
	}

	return nil
}

type YtdlpParse struct {
	minion.WorkerDefaults[*YtdlpParse]
	PageID  primitive.ObjectID
	Source  string
	URL     string
	Info    *ytdlp.Info
	Name    string
	Season  int
	Episode int
}

func (j *YtdlpParse) Kind() string { return "ytdlp_parse" }
func (j *YtdlpParse) Work(ctx context.Context, job *minion.Job[*YtdlpParse]) error {
	a := ContextApp(ctx)
	l := a.Log.Named("ytdlp.parse")
	name := job.Args.Name
	season := job.Args.Season
	episode := job.Args.Episode
	url := job.Args.URL
	info := job.Args.Info
	pid := job.Args.PageID

	l.Warnf("%s %d %s [%s] URL:%s", info.Fulltitle, info.Height, info.EXT, info.DisplayID, info.WebpageURL)

	page := &Page{}
	err := app.DB.Page.FindByID(pid, page)
	if err != nil {
		return fae.Wrap(err, "finding page")
	}

	if season == 0 || episode == 0 {
		season, episode = ParseFulltitle(info.Fulltitle)

		if season == 0 || episode == 0 {
			name = info.Fulltitle
		}
	}

	// count, err := app.DB.Video.Query().Where("display_id", info.DisplayID).Count()
	// if err != nil {
	// 	return fae.Wrap(err, "couting")
	// }
	// if count > 0 {
	// 	return nil
	// }

	for _, format := range info.Formats {
		if format.Height < 720 || format.FilesizeApprox == 0 {
			continue
		}

		video, err := app.DB.VideoFindOrCreate(fmt.Sprintf("%s-%d", info.DisplayID, format.Height))
		if err != nil {
			return fae.Wrap(err, "finding or creating")
		}

		video.PageID = pid
		video.Title = name
		video.Season = season
		video.Episode = episode
		video.Raw = info.Fulltitle
		video.Resolution = int(format.Height)
		video.Extension = format.EXT
		video.DisplayID = info.DisplayID
		video.Download = format.URL
		video.View = url
		video.Size = format.FilesizeApprox
		video.Source = page.Name

		if err := app.DB.Video.Save(video); err != nil {
			return fae.Wrap(err, "saving")
		}
	}

	return nil
}
