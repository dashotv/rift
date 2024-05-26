package scraper

import "go.uber.org/zap"

type Scraper interface {
	Read(url string) []string
	Parse(url string) []*Result
}

func New(name string, log *zap.SugaredLogger) Scraper {
	switch name {
	case "myanime":
		return NewMyAnime(log)
	case "jhdanime":
		return NewJhdAnime(log)
	case "animexin":
		return NewAnimeXin(log)
	case "animekhor":
		return NewAnimeKhor(log)
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
