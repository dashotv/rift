package server

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/dashotv/minion"
)

type jobService struct {
	db  *Connection
	log *zap.SugaredLogger
	bg  *minion.Minion
}

func (s *jobService) Index(c echo.Context, req *Request) (*Response, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 100
	}

	count, err := s.db.Job.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Job.Query().Limit(limit).Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return &Response{Total: count, Results: list}, nil
}

func (s *jobService) Create(c echo.Context, req *Request) (*Response, error) {
	id := req.ID
	switch id {
	case "scrape_all":
		if err := s.bg.Enqueue(&ScrapeAll{}); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unknown job")
	}
	return &Response{Results: true}, nil
}

func (s *jobService) Update(c echo.Context, req *Job) (*Response, error) {
	return nil, errors.New("not implemented")
}
