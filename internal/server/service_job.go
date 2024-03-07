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

func (s *jobService) Index(c echo.Context, req *Request) (*JobsResponse, error) {
	limit, skip := reqLimitSkip(req)

	count, err := s.db.Job.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Job.Query().Limit(limit).Skip(skip).Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return &JobsResponse{Total: count, Results: list}, nil
}

func (s *jobService) Create(c echo.Context, req *Request) (*JobResponse, error) {
	id := req.ID
	switch id {
	case "scrape_all":
		if err := s.bg.Enqueue(&ScrapeAll{}); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unknown job")
	}
	return &JobResponse{Job: &Job{Kind: id}}, nil
}

func (s *jobService) Update(c echo.Context, req *Job) (*JobResponse, error) {
	return nil, errors.New("not implemented")
}
