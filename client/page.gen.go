// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"context"
	"fmt"

	"github.com/dashotv/fae"
)

type PageService struct {
	client *Client
}

// NewPage makes a new client for accessing Page services.
func NewPageService(client *Client) *PageService {
	return &PageService{
		client: client,
	}
}

type PageIndexRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PageIndexResponse struct {
	*Response
	Result []*Page `json:"result"`
	Total  int64   `json:"total"`
}

func (s *PageService) Index(ctx context.Context, req *PageIndexRequest) (*PageIndexResponse, error) {
	result := &PageIndexResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetQueryParam("page", fmt.Sprintf("%v", req.Page)).
		SetQueryParam("limit", fmt.Sprintf("%v", req.Limit)).
		Get("/page/")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageCreateRequest struct {
	Subject *Page `json:"subject"`
}

type PageCreateResponse struct {
	*Response
	Result *Page `json:"result"`
}

func (s *PageService) Create(ctx context.Context, req *PageCreateRequest) (*PageCreateResponse, error) {
	result := &PageCreateResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		Post("/page/")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageShowRequest struct {
	Id string `json:"id"`
}

type PageShowResponse struct {
	*Response
	Result *Page `json:"result"`
}

func (s *PageService) Show(ctx context.Context, req *PageShowRequest) (*PageShowResponse, error) {
	result := &PageShowResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetPathParam("id", fmt.Sprintf("%v", req.Id)).
		Get("/page/{id}")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageUpdateRequest struct {
	Id      string `json:"id"`
	Subject *Page  `json:"subject"`
}

type PageUpdateResponse struct {
	*Response
	Result *Page `json:"result"`
}

func (s *PageService) Update(ctx context.Context, req *PageUpdateRequest) (*PageUpdateResponse, error) {
	result := &PageUpdateResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetPathParam("id", fmt.Sprintf("%v", req.Id)).
		Put("/page/{id}")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageSettingsRequest struct {
	Id      string   `json:"id"`
	Setting *Setting `json:"setting"`
}

type PageSettingsResponse struct {
	*Response
	Result *Page `json:"result"`
}

func (s *PageService) Settings(ctx context.Context, req *PageSettingsRequest) (*PageSettingsResponse, error) {
	result := &PageSettingsResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetPathParam("id", fmt.Sprintf("%v", req.Id)).
		Patch("/page/{id}")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageDeleteRequest struct {
	Id string `json:"id"`
}

type PageDeleteResponse struct {
	*Response
	Result *Page `json:"result"`
}

func (s *PageService) Delete(ctx context.Context, req *PageDeleteRequest) (*PageDeleteResponse, error) {
	result := &PageDeleteResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetPathParam("id", fmt.Sprintf("%v", req.Id)).
		Delete("/page/{id}")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}

type PageVideosRequest struct {
	Id    string `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

type PageVideosResponse struct {
	*Response
	Result []*Video `json:"result"`
}

func (s *PageService) Videos(ctx context.Context, req *PageVideosRequest) (*PageVideosResponse, error) {
	result := &PageVideosResponse{Response: &Response{}}
	resp, err := s.client.Resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(result).
		SetQueryParam("page", fmt.Sprintf("%v", req.Page)).
		SetQueryParam("limit", fmt.Sprintf("%v", req.Limit)).
		SetPathParam("id", fmt.Sprintf("%v", req.Id)).
		Get("/page/{id}/videos")
	if err != nil {
		return nil, fae.Wrap(err, "failed to make request")
	}
	if !resp.IsSuccess() {
		return nil, fae.Errorf("%d: %v", resp.StatusCode(), resp.String())
	}
	if result.Error {
		return nil, fae.New(result.Message)
	}

	return result, nil
}
