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

var khorPagesNum = 25
var khorUrlsNum = 1
var khorDomains = regexp.MustCompile(`(?i)^http(?:s)*://.*\.(?:dailymotion|ok\.ru)`)
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
	if len(urls) > khorPagesNum {
		return urls[:khorPagesNum]
	}
	return urls
}

func (m *AnimeKhor) ReadPage(url string) []string {
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
	if len(urls) > khorUrlsNum {
		return urls[:khorUrlsNum]
	}
	return urls
}

func (m *AnimeKhor) Parse(url string) []*Result {
	results := []*Result{}
	pages := m.Read(url)
	for _, p := range pages {
		// m.log.Infof("readpage: %s", p)

		match := khorRegex.FindAllStringSubmatch(p, -1)
		if len(match) == 0 {
			continue
		}

		title := strings.ReplaceAll(match[0][1], "-", " ")
		season, _ := strconv.Atoi(match[0][2])
		episode, _ := strconv.Atoi(match[0][3])

		list := m.ReadPage(p)
		for _, l := range list {
			// m.log.Infof("list: %s", l)
			if !khorDomains.MatchString(l) {
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

//
// func (m *AnimeKhor) Read(url string) []string {
// 	urls := []string{}
// 	m.col.OnHTML("article", func(e *colly.HTMLElement) {
// 		urls = append(urls, e.ChildAttr("a", "href"))
// 	})
// 	m.col.OnHTML(".eplister li", func(e *colly.HTMLElement) {
// 		urls = append(urls, e.ChildAttr("a", "href"))
// 	})
// 	m.col.OnError(func(r *colly.Response, err error) {
// 		m.log.Errorf("scraping: %s\n", err)
// 	})
// 	m.col.Visit(url)
// 	return urls
// }
//
// func (m *AnimeKhor) Parse(url string) []*Result {
// 	results := []*Result{}
// 	list := m.Read(url)
// 	for _, l := range list {
// 		eps := m.Read(l)
// 		if len(eps) > 25 {
// 			eps = eps[:25]
// 		}
// 		for _, e := range eps {
// 			match := khorRegex.FindAllStringSubmatch(e, -1)
// 			if len(match) > 0 {
// 				// m.log.Infof("match: %v\n", match[0])
// 				title := strings.Replace(match[0][1], "-", " ", -1)
// 				season, _ := strconv.Atoi(match[0][2])
// 				episode, _ := strconv.Atoi(match[0][3])
// 				results = append(results, &Result{
// 					Title:   title,
// 					Season:  season,
// 					Episode: episode,
// 					URL:     e,
// 				})
// 			}
// 		}
// 	}
// 	return results
// }
