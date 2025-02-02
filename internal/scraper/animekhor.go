package scraper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

var khorRegex = regexp.MustCompile(`(?i)^http(?:s)*://animekhor\..*/([\w-]+?)(?:-season-(\d+))*-episode(?:s)*-(\d+)`)

func NewAnimeKhor(log *zap.SugaredLogger, col *colly.Collector) *AnimeKhor {
	return &AnimeKhor{
		col: col,
		log: log.Named("animekhor"),
	}
}

type AnimeKhor struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *AnimeKhor) Read(url string) []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnHTML(".eplister li", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("a", "href"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		m.log.Errorf("scraping: %s\n", err)
	})
	m.col.Visit(url)
	return urls
}

func (m *AnimeKhor) Parse(url string) []*Result {
	results := []*Result{}
	list := m.Read(url)
	for _, l := range list {
		eps := m.Read(l)
		if len(eps) > 25 {
			eps = eps[:25]
		}
		for _, e := range eps {
			match := khorRegex.FindAllStringSubmatch(e, -1)
			if len(match) > 0 {
				// m.log.Infof("match: %v\n", match[0])
				title := strings.Replace(match[0][1], "-", " ", -1)
				season, _ := strconv.Atoi(match[0][2])
				episode, _ := strconv.Atoi(match[0][3])
				results = append(results, &Result{
					Title:   title,
					Season:  season,
					Episode: episode,
					URL:     e,
				})
			}
		}
	}
	return results
}
