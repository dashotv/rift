package server

import "github.com/labstack/echo/v4"

type workerService struct {
	bg *Workers
}

func (s *workerService) Enqueue(c echo.Context, req *Request) (*Response, error) {
	switch req.ID {
	case "scrape":
		s.bg.Enqueue(s.bg.ScrapeJob())
	}
	return &Response{}, nil
}
