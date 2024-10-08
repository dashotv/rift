package app

import (
	"context"
	"time"

	"github.com/dashotv/fae"
	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/scraper"
)

type ScrapeAll struct {
	minion.WorkerDefaults[*ScrapeAll]
}

func (j *ScrapeAll) Kind() string { return "scrape_all" }
func (j *ScrapeAll) Work(ctx context.Context, job *minion.Job[*ScrapeAll]) error {
	l := app.Log.Named("scrape")
	if !app.Config.Production {
		l.Warn("scrape_all: skipping")
		return nil
	}

	pages, err := app.DB.Page.Query().Limit(-1).Where("enabled", true).Desc("name").Run()
	if err != nil {
		return fae.Wrap(err, "loading pages")
	}

	// l.Debugf("scraping all %d pages", len(pages))
	for _, p := range pages {
		// l.Debugf("page: %s", p.Name)
		if err := app.Workers.Enqueue(&ScrapePage{Title: p.Name, Page: p}); err != nil {
			return fae.Wrap(err, "scrape_pages: enqueuing scrape_page")
		}
	}
	return nil
}

type ScrapePage struct {
	minion.WorkerDefaults[*ScrapePage]
	Title string `json:"title"`
	Page  *Page  `json:"page"`
}

func (j *ScrapePage) Kind() string { return "scrape_page" }
func (j *ScrapePage) Work(ctx context.Context, job *minion.Job[*ScrapePage]) error {
	p := job.Args.Page
	l := app.Log.Named("scrape.page")

	// l.Debugf("scrape: %s", p.Name)
	scr := scraper.New(p.Scraper, l)
	if scr == nil {
		return fae.Errorf("invalid scraper: %s", p.Scraper)
	}

	// l.Debugf("parse: %s", p.URL)
	results := scr.Parse(p.URL)

	// l.Debugf("results: %d", len(results))
	for _, result := range results {
		ok, err := app.DB.IsVisited(p, result.URL)
		if err != nil {
			return fae.Wrap(err, "scrape_page: is_visited")
		}
		if ok {
			// l.Debugf("'%s' already visited: %s", p.Name, url)
			continue
		}

		// l.Debugf("'%s' %s", p.Name, url)
		if err := app.Workers.Enqueue(&YtdlpList{PageID: p.ID, Name: result.Title, URL: result.URL, Season: result.Season, Episode: result.Episode}); err != nil {
			return fae.Wrap(err, "scrape_page_url: enqueuing ytdlp_list")
		}
	}

	p.ProcessedAt = time.Now()
	if err := app.DB.Page.Save(p); err != nil {
		return fae.Wrap(err, "scrape_page: saving page")
	}
	return nil
}
