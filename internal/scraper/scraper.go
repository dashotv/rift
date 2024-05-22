package scraper

import "go.uber.org/zap"

type Scraper interface {
	Read(url string) []string
}

func New(name string, log *zap.SugaredLogger) Scraper {
	switch name {
	case "myanime":
		return NewMyAnime(log)
	case "jhdanime":
		return NewJhdAnime(log)
	default:
		return nil
	}
}
