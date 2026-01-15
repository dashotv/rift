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

var narulRegex = regexp.MustCompile(`(?i)^http(?:s)*://naruldonghua\.(?:[\w]+)/([\w-]+?)(?:-season-(\d+))*-episode(?:s)*-(\d+)`)
var narulDomains = regexp.MustCompile(`(?i)^http(?:s)*://.*\.(?:dailymotion|ok\.ru)`)

func NewNarulDonghua(log *zap.SugaredLogger, col *colly.Collector) *NarulDonghua {
	return &NarulDonghua{
		col: col,
		log: log.Named("naruldonghua"),
	}
}

type NarulDonghua struct {
	col *colly.Collector
	log *zap.SugaredLogger
}

func (m *NarulDonghua) Read(url string) []string {
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

func (m *NarulDonghua) ReadPage(url string) []string {
	urls := []string{}
	m.col.OnHTML("select.mirror > option", func(e *colly.HTMLElement) {
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

func (m *NarulDonghua) Parse(url string) []*Result {
	results := []*Result{}
	pages := m.Read(url)
	for _, p := range pages {
		// m.log.Infof("readpage: %s", p)

		match := narulRegex.FindAllStringSubmatch(p, -1)
		if len(match) == 0 {
			continue
		}

		title := strings.ReplaceAll(match[0][1], "-", " ")
		season, _ := strconv.Atoi(match[0][2])
		episode, _ := strconv.Atoi(match[0][3])

		list := m.ReadPage(p)
		for _, l := range list {
			m.log.Infof("list: %s", l)
			if !narulDomains.MatchString(l) {
				continue
			}
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
