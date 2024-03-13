package server

import "github.com/dashotv/rift/internal/scraper"

type Scraper interface {
	Read(url string) []string
}

func ScraperFromPage(page *Page) Scraper {
	switch page.Scraper {
	case "myanime":
		return scraper.NewMyAnime()
	default:
		return nil
	}
}
