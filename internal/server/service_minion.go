package server

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/dashotv/minion"
)

type minionService struct {
	db  *Connection
	log *zap.SugaredLogger
	bg  *minion.Minion
}

func (s *minionService) Index(c echo.Context, req *Request) (*Response, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 100
	}

	count, err := s.db.Minion.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Minion.Query().Limit(limit).Run()
	if err != nil {
		return nil, err
	}

	return &Response{Total: count, Results: list}, nil
}

func (s *minionService) Create(c echo.Context, req *Request) (*Response, error) {
	id := req.ID
	switch id {
	case "scrape_pages":
		if err := s.bg.Enqueue(&ScrapePages{}); err != nil {
			return nil, err
		}
	}
	return &Response{Results: true}, nil
}

func (s *minionService) Update(c echo.Context, req *Minion) (*Response, error) {
	return nil, errors.New("not implemented")
}
