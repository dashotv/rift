package scraper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

// https://donghuastream.org/
var donghuaStreamRegex = regexp.MustCompile(`(?i)^http(?:s)*://donghuastream.org/([\w-]+?)(?:-season-(\d+))*-episode(?:s)*-(\d+)`)

func NewDonghuaStream(log *zap.SugaredLogger) *DonghuaStream {
	return &DonghuaStream{
		col: colly.NewCollector(),
		log: log.Named("donghuastream"),
	}
}

type DonghuaStream struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *DonghuaStream) Read(url string) []string {
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

func (m *DonghuaStream) ReadPage(url string) []string {
	urls := []string{}
	m.col.OnHTML("article", func(e *colly.HTMLElement) {
		urls = append(urls, e.ChildAttr("iframe", "src"))
	})
	m.col.OnError(func(r *colly.Response, err error) {
		m.log.Errorf("scraping: %s", err)
	})
	m.col.Visit(url)
	return urls
}

func (m *DonghuaStream) Parse(url string) []*Result {
	results := []*Result{}
	list := m.Read(url)
	for _, l := range list {
		match := donghuaStreamRegex.FindAllStringSubmatch(l, -1)
		if len(match) > 0 {
			videos := m.ReadPage(l)
			if len(videos) == 0 {
				continue
			}
			title := strings.Replace(match[0][1], "-", " ", -1)
			season, _ := strconv.Atoi(match[0][2])
			episode, _ := strconv.Atoi(match[0][3])
			results = append(results, &Result{
				Title:   title,
				Season:  season,
				Episode: episode,
				URL:     videos[0],
			})
		}
	}
	return results
}
