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
	// l := app.Log.Named("scrape")

	pages, err := app.DB.Page.Query().Limit(-1).Desc("name").Run()
	if err != nil {
		return fae.Errorf("scrape: %s", err)
	}

	// l.Debugf("scraping all %d pages", len(pages))
	for _, p := range pages {
		// l.Debugf("page: %s", p.Name)
		if err := app.Workers.Enqueue(&ScrapePage{Title: p.Name, Page: p}); err != nil {
			return fae.Errorf("scrape_pages: enqueuing scrape_page: %w", err)
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
	if app.Config.Production {
		scr := scraper.NewMyAnime(l)
		urls := scr.Read(p.Url)

		for _, url := range urls {
			ok, err := app.DB.IsVisited(p, url)
			if err != nil {
				return fae.Errorf("scrape_page: is_visited: %w", err)
			}
			if ok {
				continue
			}

			// l.Debugf("'%s' %s", p.Name, url)
			if err := app.Workers.Enqueue(&YtdlpList{Name: p.Name, PageID: p.ID, URL: url}); err != nil {
				return fae.Errorf("scrape_page_url: enqueuing ytdlp_list: %w", err)
			}
		}

		p.ProcessedAt = time.Now()
		if err := app.DB.Page.Save(p); err != nil {
			return fae.Errorf("scrape_page: saving page: %w", err)
		}
	}
	return nil
}
