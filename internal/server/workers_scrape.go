package server

import (
	"context"
	"fmt"

	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/scraper"
)

type ScrapePages struct {
	minion.WorkerDefaults[*ScrapePages]
}

func (j *ScrapePages) Kind() string { return "scrape_pages" }
func (j *ScrapePages) Work(ctx context.Context, job *minion.Job[*ScrapePages]) error {
	s := getServer(ctx)

	l := s.Logger.Named("scrape")
	pages, err := s.db.Page.Query().Run()
	if err != nil {
		return fmt.Errorf("scrape: %s", err)
	}

	for _, p := range pages {
		l.Debugf("scrape: %s", p.Name)
		s.bg.Enqueue(&ScrapePage{Page: p})
	}
	return nil
}

type ScrapePage struct {
	minion.WorkerDefaults[*ScrapePage]
	Page *Page
}

func (j *ScrapePage) Kind() string { return "scrape_page" }
func (j *ScrapePage) Work(ctx context.Context, job *minion.Job[*ScrapePage]) error {
	p := job.Args.Page
	s := getServer(ctx)
	l := s.Logger.Named("scrape.page")
	scr := scraper.NewMyAnime()
	urls := scr.Read(p.URL)

	for _, url := range urls {
		l.Debugf("scrape: %s %s", p.Name, url)
		s.bg.Enqueue(&ScrapePageURL{Page: p, URL: url})
	}
	return nil
}

type ScrapePageURL struct {
	minion.WorkerDefaults[*ScrapePageURL]
	Page *Page
	URL  string
}

func (j *ScrapePageURL) Kind() string { return "scrape_page_url" }
func (j *ScrapePageURL) Work(ctx context.Context, job *minion.Job[*ScrapePageURL]) error {
	p := job.Args.Page
	u := job.Args.URL
	s := getServer(ctx)
	l := s.Logger.Named("scrape.page.url")
	count, err := s.db.Video.Query().Where("page_id", p.ID).Where("url", u).Count()
	if err != nil {
		return fmt.Errorf("scrape: %s", err)
	}
	if count > 0 {
		return nil
	}

	l.Debugf("scrape: %s %s", p.Name, u)
	s.bg.Enqueue(&YtdlpListJob{Name: p.Name, URL: u})
	return nil
}
