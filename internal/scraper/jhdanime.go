package scraper

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

func NewJhdAnime(log *zap.SugaredLogger) *JhdAnime {
	return &JhdAnime{
		col: colly.NewCollector(),
		log: log.Named("jhdanime"),
	}
}

type JhdAnime struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *JhdAnime) Read(url string) []string {
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
