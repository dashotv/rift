package scraper

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
)

type Scraper interface {
	Read(url string) []string
	Parse(url string) []*Result
}

func New(name string, log *zap.SugaredLogger) Scraper {
	ua := "Mozilla/5.0 colly/v2 rift/" + VERSION + " " + name
	col := colly.NewCollector(colly.UserAgent(ua))
	switch name {
	case "myanime":
		return NewMyAnime(log, col)
	case "jhdanime":
		return NewJhdAnime(log, col)
	case "animexin":
		return NewAnimeXin(log, col)
	case "animekhor":
		return NewAnimeKhor(log, col)
	case "naruldonghua":
		return NewNarulDonghua(log, col)
	case "donghuastream":
		return NewDonghuaStream(log, col)
	default:
		return nil
	}
}

type Result struct {
	Title   string
	Season  int
	Episode int
	URL     string
}
