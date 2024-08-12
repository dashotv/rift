package app

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/ytdlp"
)

type YtdlpList struct {
	minion.WorkerDefaults[*YtdlpList]
	Name   string
	PageID primitive.ObjectID
	URL    string
}

func (j *YtdlpList) Kind() string { return "ytdlp_list" }
func (j *YtdlpList) Work(ctx context.Context, job *minion.Job[*YtdlpList]) (err error) {
	a := ContextApp(ctx)
	l := a.Log.Named("ytdlp.list")
	name := job.Args.Name
	url := job.Args.URL
	pid := job.Args.PageID

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
		if err := app.Workers.Enqueue(&YtdlpParse{Name: name, PageID: pid, URL: url, Info: e}); err != nil {
			return fae.Wrapf(err, "ytdlp-list: info: %s", e.WebpageURL)
		}
	}

	return nil
}

type YtdlpParse struct {
	minion.WorkerDefaults[*YtdlpParse]
	Name   string
	PageID primitive.ObjectID
	Source string
	URL    string
	Info   *ytdlp.Info
}

func (j *YtdlpParse) Kind() string { return "ytdlp_parse" }
func (j *YtdlpParse) Work(ctx context.Context, job *minion.Job[*YtdlpParse]) error {
	a := ContextApp(ctx)
	l := a.Log.Named("ytdlp.parse")
	name := job.Args.Name
	url := job.Args.URL
	info := job.Args.Info
	pid := job.Args.PageID

	l.Warnf("%s %d %s [%s] URL:%s", info.Fulltitle, info.Height, info.EXT, info.DisplayID, info.WebpageURL)

	page := &Page{}
	err := app.DB.Page.FindByID(pid, page)
	if err != nil {
		return fae.Wrap(err, "finding page")
	}

	season, episode := ParseFulltitle(info.Fulltitle)
	if episode == 0 {
		season, episode = ParseURL(url)
	}

	if season == 0 || episode == 0 {
		name = info.Fulltitle
	}

	// count, err := app.DB.Video.Query().Where("display_id", info.DisplayID).Count()
	// if err != nil {
	// 	return fae.Wrap(err, "couting")
	// }
	// if count > 0 {
	// 	return nil
	// }

	video, err := app.DB.VideoFindOrCreate(info.DisplayID)
	if err != nil {
		return fae.Wrap(err, "finding or creating")
	}

	video.PageID = pid
	video.Title = name
	video.Season = season
	video.Episode = episode
	video.Raw = info.Fulltitle
	video.Resolution = int(info.Height)
	video.Extension = info.EXT
	video.DisplayID = info.DisplayID
	video.Download = info.WebpageURL
	video.View = url
	video.Size = info.FilesizeApprox
	video.Source = page.Name

	if err := app.DB.Video.Save(video); err != nil {
		return fae.Wrap(err, "saving")
	}

	return nil
}
