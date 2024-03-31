package scraper

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

func NewMyAnime(log *zap.SugaredLogger) *MyAnime {
	return &MyAnime{
		col: colly.NewCollector(),
		log: log.Named("myanime"),
	}
}

type MyAnime struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *MyAnime) Read(url string) []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		m.log.Errorf("scraping: %s\n", err)
	})
	m.col.Visit(url)
	return urls
}
