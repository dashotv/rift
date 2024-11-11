package scraper

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

// https://donghuastream.org/
var donghuaStreamRegex = regexp.MustCompile(`(?i)^http(?:s)*://donghuastream.org/([\w-]+?)(?:-season-(\d+))*-episode(?:s)*-(\d+)(?:-(\d+))*`)

func NewDonghuaStream(log *zap.SugaredLogger, col *colly.Collector) *DonghuaStream {
	return &DonghuaStream{
		col: col,
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
	m.col.OnHTML(".eplister li", func(e *colly.HTMLElement) {
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
		s := e.ChildAttr("iframe", "src")
		if s == "" || s == "about:blank" {
			return
		}

		urls = append(urls, s)
	})
	m.col.OnHTML("option", func(e *colly.HTMLElement) {
		s := e.Attr("value")
		if s == "" {
			return
		}

		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			m.log.Errorf("base64: %s", err)
			return
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
		if err != nil {
			m.log.Errorf("goquery: %s", err)
			return
		}

		doc.Find("iframe").Each(func(i int, s *goquery.Selection) {
			u, ok := s.Attr("src")
			if ok {
				urls = append(urls, u)
			}
		})
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
			// fmt.Printf("video: %+v\n", videos)
			// fmt.Printf("match: %+v\n", match[0])
			title := strings.Replace(match[0][1], "-", " ", -1)
			season, _ := strconv.Atoi(match[0][2])
			episode, _ := strconv.Atoi(match[0][3])
			if len(match[0]) > 4 && match[0][4] != "" {
				episode, _ = strconv.Atoi(match[0][4])
			}
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
