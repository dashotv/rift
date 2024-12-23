package scraper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

var xinRegex = regexp.MustCompile(`(?i)^http(?:s)*://animexin\..*/([\w-]+?)(?:-season-(\d+))*-episode-(\d+)`)

func NewAnimeXin(log *zap.SugaredLogger, col *colly.Collector) *AnimeXin {
	return &AnimeXin{
		col: col,
		log: log.Named("animexin"),
	}
}

type AnimeXin struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *AnimeXin) Read(url string) []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnHTML("div.eplister li", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		m.log.Errorf("scraping: %s\n", err)
	})
	m.col.Visit(url)
	return urls
}

func (m *AnimeXin) Parse(url string) []*Result {
	results := []*Result{}
	list := m.Read(url)
	for _, l := range list {
		match := xinRegex.FindAllStringSubmatch(l, -1)
		if len(match) > 0 {
			// m.log.Infof("match: %v\n", match[0])
			title := strings.Replace(match[0][1], "-", " ", -1)
			season, _ := strconv.Atoi(match[0][2])
			episode, _ := strconv.Atoi(match[0][3])
			results = append(results, &Result{
				Title:   title,
				Season:  season,
				Episode: episode,
				URL:     l,
			})
		}
	}
	return results
}
