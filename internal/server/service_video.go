package server

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type videoService struct {
	db  *Connection
	log *zap.SugaredLogger
}

func (s *videoService) Index(c echo.Context, req *Request) (*Response, error) {
	count, err := s.db.Video.Query().Count()
	if err != nil {
		return nil, err
	}

	list, err := s.db.Video.Query().Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return &Response{Total: count, Results: list}, nil
}
func (s *videoService) Show(c echo.Context, req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *videoService) Update(c echo.Context, req *Video) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *videoService) Create(c echo.Context, req *Video) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *videoService) Delete(c echo.Context, req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
