package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/dashotv/minion"
)

type pageService struct {
	bg  *minion.Minion
	db  *Connection
	log *zap.SugaredLogger
}

func (s *pageService) Index(c echo.Context, req *Request) (*PagesResponse, error) {
	limit, skip := reqLimitSkip(req)

	// TODO: limit and offset
	count, err := s.db.Page.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Page.Query().Desc("created_at").Limit(limit).Skip(skip).Run()
	if err != nil {
		return nil, err
	}

	return &PagesResponse{Total: count, Results: list}, nil
}
func (s *pageService) Show(c echo.Context, req *Request) (*PageResponse, error) {
	page, err := s.dbGet(req.ID)
	if err != nil {
		return nil, err
	}

	return &PageResponse{Page: page}, nil
}
func (s *pageService) Update(c echo.Context, req *Page) (*PageResponse, error) {
	if !s.dbExists(req.Name) {
		return nil, echo.NewHTTPError(http.StatusNotFound, errors.New("page not found"))
	}

	if err := s.dbSet(req); err != nil {
		return nil, err
	}

	if err := s.bg.Enqueue(&ScrapePage{Page: req}); err != nil {
		return nil, err
	}

	return &PageResponse{Page: req}, nil
}
func (s *pageService) Create(c echo.Context, req *Page) (*PageResponse, error) {
	s.log.Debugf("creating page: %+v", req)
	if s.dbExists(req.Name) {
		s.log.Debugf("creating page: already exists: %s", req.Name)
		return nil, echo.NewHTTPError(http.StatusNotFound, errors.New("page already exists"))
	}

	if err := s.dbSet(req); err != nil {
		return nil, err
	}

	if err := s.bg.Enqueue(&ScrapePage{Page: req}); err != nil {
		return nil, err
	}

	return &PageResponse{Page: req}, nil
}

func (s *pageService) Delete(c echo.Context, req *Request) (*PageResponse, error) {
	page, err := s.dbGet(req.ID)
	if err != nil {
		return nil, fmt.Errorf("error finding id: %s: %w", req.ID, err)
	}

	if err := s.dbDelete(req.ID); err != nil {
		return nil, err
	}

	return &PageResponse{Page: page}, nil
}

func (s *pageService) dbList() ([]*Page, error) {
	list, err := s.db.Page.Query().Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *pageService) dbGet(key string) (*Page, error) {
	p := &Page{}
	err := s.db.Page.Find(key, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *pageService) dbSet(page *Page) error {
	return s.db.Page.Save(page)
}

func (s *pageService) dbDelete(key string) error {
	page, err := s.dbGet(key)
	if err != nil {
		return err
	}

	return s.db.Page.Delete(page)
}

func (s *pageService) dbExists(key string) bool {
	page, err := s.dbGet(key)
	if err != nil {
		return false
	}
	if page == nil {
		return false
	}
	return true
}
