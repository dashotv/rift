package server

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type visitService struct {
	db  *Connection
	log *zap.SugaredLogger
}

func (s *visitService) Index(c echo.Context, req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Show(c echo.Context, req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Create(c echo.Context, req *Visit) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Update(c echo.Context, req *Visit) (*Response, error) {
	return nil, errors.New("not implemented")
}
func (s *visitService) Delete(c echo.Context, req *Request) (*Response, error) {
	return nil, errors.New("not implemented")
}
