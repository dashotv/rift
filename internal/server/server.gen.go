// Code generated by oto; DO NOT EDIT.

package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageService interface {
	Create(echo.Context, *Page) (*PageResponse, error)
	Delete(echo.Context, *Request) (*PageResponse, error)
	Index(echo.Context, *Request) (*PagesResponse, error)
	Show(echo.Context, *Request) (*PageResponse, error)
	Update(echo.Context, *Page) (*PageResponse, error)
}

type VideoService interface {
	Create(echo.Context, *Video) (*VideoResponse, error)
	Delete(echo.Context, *Request) (*VideoResponse, error)
	Index(echo.Context, *Request) (*VideosResponse, error)
	Show(echo.Context, *Request) (*VideoResponse, error)
	Update(echo.Context, *Video) (*VideoResponse, error)
}

type VisitService interface {
	Create(echo.Context, *Visit) (*VisitResponse, error)
	Delete(echo.Context, *Request) (*VisitResponse, error)
	Index(echo.Context, *Request) (*VisitsResponse, error)
	Show(echo.Context, *Request) (*VisitResponse, error)
	Update(echo.Context, *Visit) (*VisitResponse, error)
}

type pageServiceServer struct {
	pageService PageService
}

// Register adds the PageService to the otohttp.Server.
func RegisterPageService(e *echo.Group, pageService PageService) {
	handler := &pageServiceServer{
		pageService: pageService,
	}
	e.POST("/PageService.Create", handler.handleCreate)
	e.POST("/PageService.Delete", handler.handleDelete)
	e.POST("/PageService.Index", handler.handleIndex)
	e.POST("/PageService.Show", handler.handleShow)
	e.POST("/PageService.Update", handler.handleUpdate)
}

func (s *pageServiceServer) handleCreate(c echo.Context) error {
	request := &Page{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.pageService.Create(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *pageServiceServer) handleDelete(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.pageService.Delete(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *pageServiceServer) handleIndex(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.pageService.Index(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *pageServiceServer) handleShow(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.pageService.Show(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *pageServiceServer) handleUpdate(c echo.Context) error {
	request := &Page{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.pageService.Update(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

type videoServiceServer struct {
	videoService VideoService
}

// Register adds the VideoService to the otohttp.Server.
func RegisterVideoService(e *echo.Group, videoService VideoService) {
	handler := &videoServiceServer{
		videoService: videoService,
	}
	e.POST("/VideoService.Create", handler.handleCreate)
	e.POST("/VideoService.Delete", handler.handleDelete)
	e.POST("/VideoService.Index", handler.handleIndex)
	e.POST("/VideoService.Show", handler.handleShow)
	e.POST("/VideoService.Update", handler.handleUpdate)
}

func (s *videoServiceServer) handleCreate(c echo.Context) error {
	request := &Video{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.videoService.Create(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *videoServiceServer) handleDelete(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.videoService.Delete(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *videoServiceServer) handleIndex(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.videoService.Index(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *videoServiceServer) handleShow(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.videoService.Show(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *videoServiceServer) handleUpdate(c echo.Context) error {
	request := &Video{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.videoService.Update(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

type visitServiceServer struct {
	visitService VisitService
}

// Register adds the VisitService to the otohttp.Server.
func RegisterVisitService(e *echo.Group, visitService VisitService) {
	handler := &visitServiceServer{
		visitService: visitService,
	}
	e.POST("/VisitService.Create", handler.handleCreate)
	e.POST("/VisitService.Delete", handler.handleDelete)
	e.POST("/VisitService.Index", handler.handleIndex)
	e.POST("/VisitService.Show", handler.handleShow)
	e.POST("/VisitService.Update", handler.handleUpdate)
}

func (s *visitServiceServer) handleCreate(c echo.Context) error {
	request := &Visit{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.visitService.Create(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *visitServiceServer) handleDelete(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.visitService.Delete(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *visitServiceServer) handleIndex(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.visitService.Index(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *visitServiceServer) handleShow(c echo.Context) error {
	request := &Request{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.visitService.Show(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *visitServiceServer) handleUpdate(c echo.Context) error {
	request := &Visit{}
	if err := c.Bind(request); err != nil {
		return fmt.Errorf("binding request: %w", err)
	}

	response, err := s.visitService.Update(c, request)
	if err != nil {
		return fmt.Errorf("handling request: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}
