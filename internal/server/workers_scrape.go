package server

import (
	"context"
	"fmt"

	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/scraper"
)

type ScrapeAll struct {
	minion.WorkerDefaults[*ScrapeAll]
}

func (j *ScrapeAll) Kind() string { return "scrape_all" }
func (j *ScrapeAll) Work(ctx context.Context, job *minion.Job[*ScrapeAll]) error {
	s := getServer(ctx)
	l := s.Logger.Named("scrape")

	pages, err := s.db.Page.Query().Run()
	if err != nil {
		return fmt.Errorf("scrape: %s", err)
	}

	l.Debugf("scraping all %d pages", len(pages))
	for _, p := range pages {
		// l.Debugf("page: %s", p.Name)
		if err := s.bg.Enqueue(&ScrapePage{Page: p}); err != nil {
			return fmt.Errorf("scrape_pages: enqueuing scrape_page: %w", err)
		}
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
		if ok, err := s.db.IsVisited(p, url); err != nil {
			return fmt.Errorf("scrape_page: is_visited: %w", err)
		} else if ok {
			continue
		}
		l.Debugf("'%s' %s", p.Name, url)
		if err := s.bg.Enqueue(&YtdlpListJob{Name: p.Name, URL: url}); err != nil {
			return fmt.Errorf("scrape_page_url: enqueuing ytdlp_list: %w", err)
		}
	}
	return nil
}

// type ScrapePageURL struct {
// 	minion.WorkerDefaults[*ScrapePageURL]
// 	Page *Page
// 	URL  string
// }
//
// func (j *ScrapePageURL) Kind() string { return "scrape_page_url" }
// func (j *ScrapePageURL) Work(ctx context.Context, job *minion.Job[*ScrapePageURL]) error {
// 	p := job.Args.Page
// 	u := job.Args.URL
// 	s := getServer(ctx)
// 	l := s.Logger.Named("scrape.page.url")
//
// 	count, err := s.db.Visit.Query().Where("page_id", p.ID).Where("url", u).Count()
// 	if err != nil {
// 		return fmt.Errorf("scrape_page_url: counting visit: %w", err)
// 	}
// 	if count > 0 {
// 		return nil
// 	}
//
// 	if err := s.db.Visit.Save(&Visit{PageID: p.ID.Hex(), URL: u}); err != nil {
// 		return fmt.Errorf("scrape_page_url: saving visit: %w", err)
// 	}
//
// 	l.Debugf("scrape: %s %s", p.Name, u)
// 	if err := s.bg.Enqueue(&YtdlpListJob{Name: p.Name, URL: u}); err != nil {
// 		return fmt.Errorf("scrape_page_url: enqueuing ytdlp_list: %w", err)
// 	}
// 	return nil
// }
