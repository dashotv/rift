package server

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type visitService struct {
	db  *Connection
	log *zap.SugaredLogger
}

func (s *visitService) Index(c echo.Context, req *Request) (*VisitsResponse, error) {
	limit, skip := reqLimitSkip(req)

	count, err := s.db.Visit.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Visit.Query().Desc("created_at").Limit(limit).Skip(skip).Run()
	if err != nil {
		return nil, err
	}

	return &VisitsResponse{Total: count, Results: list}, nil
}
func (s *visitService) Show(c echo.Context, req *Request) (*VisitResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Create(c echo.Context, req *Visit) (*VisitResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Update(c echo.Context, req *Visit) (*VisitResponse, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Delete(c echo.Context, req *Request) (*VisitResponse, error) {
	return nil, errors.New("not implemented")
}

func (c *Connection) IsVisited(page *Page, url string) (bool, error) {
	count, err := c.Visit.Query().Where("page_id", page.ID).Where("url", url).Count()
	if err != nil {
		return false, fmt.Errorf("is_visited: counting visit: %w", err)
	}
	if count > 0 {
		return true, nil
	}

	if err := c.Visit.Save(&Visit{PageID: page.ID, URL: url}); err != nil {
		return false, fmt.Errorf("is_visited: saving visit: %w", err)
	}

	return false, nil
}
