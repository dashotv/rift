package server

import (
	"fmt"

	"github.com/dashotv/rift/internal/scraper"
)

func (w *Workers) ScrapeJob() JobFunc {
	return func() error {
		l := w.logger.Named("scrape")

		pages, err := w.db.Page.Query().Run()
		if err != nil {
			return fmt.Errorf("scrape: %s", err)
		}

		for _, p := range pages {
			l.Debugf("scrape: %s", p.Name)
			w.Enqueue(w.ScrapePageJob(p))
		}
		return nil
	}
}

func (w *Workers) ScrapePageJob(p *Page) JobFunc {
	return func() error {
		l := w.logger.Named("scrape.page")
		scr := scraper.NewMyAnime()
		urls := scr.Read(p.URL)

		for _, url := range urls {
			l.Debugf("scrape: %s %s", p.Name, url)
			w.ScrapePageURLJob(p, url)
		}
		return nil
	}
}

func (w *Workers) ScrapePageURLJob(p *Page, url string) JobFunc {
	return func() error {
		l := w.logger.Named("scrape.page.url")
		count, err := w.db.Video.Query().Where("page_id", p.ID).Where("url", url).Count()
		if err != nil {
			return fmt.Errorf("scrape: %s", err)
		}
		if count > 0 {
			return nil
		}

		l.Debugf("scrape: %s %s", p.Name, url)
		w.Enqueue(w.YtdlpListJob(p.Name, url))
		return nil
	}
}
