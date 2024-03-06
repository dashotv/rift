// Code generated by oto; DO NOT EDIT.

package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Client is used to access Pace services.
type Client struct {
	// RemoteHost is the URL of the remote server that this Client should
	// access.
	RemoteHost string
	// HTTPClient is the http.Client to use when making HTTP requests.
	HTTPClient *http.Client
	// BeforeRequest is an optional hook that gives you the opportunity
	// to inspect or modify the request before it is made.
	// Useful for adding auth headers, for example.
	BeforeRequest func(r *http.Request) error
	// Debug writes a line of debug log output.
	Debug func(s string)
}

// New makes a new Client.
func New(remoteHost string) *Client {
	c := &Client{
		RemoteHost: remoteHost,
		Debug:      func(s string) {},
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	return c
}

type JobService struct {
	client *Client
}

// NewJobService makes a new client for accessing JobService services.
func NewJobService(client *Client) *JobService {
	return &JobService{
		client: client,
	}
}

func (s *JobService) Create(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Create: marshal Request")
	}
	url := s.client.RemoteHost + "JobService.Create"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Create: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Create")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "JobService.Create: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Create: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("JobService.Create: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *JobService) Index(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Index: marshal Request")
	}
	url := s.client.RemoteHost + "JobService.Index"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Index: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Index")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "JobService.Index: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "JobService.Index: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("JobService.Index: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

type PageService struct {
	client *Client
}

// NewPageService makes a new client for accessing PageService services.
func NewPageService(client *Client) *PageService {
	return &PageService{
		client: client,
	}
}

func (s *PageService) Create(ctx context.Context, r *Page) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Create: marshal Page")
	}
	url := s.client.RemoteHost + "PageService.Create"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Create: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Create")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "PageService.Create: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Create: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("PageService.Create: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *PageService) Delete(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Delete: marshal Request")
	}
	url := s.client.RemoteHost + "PageService.Delete"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Delete: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Delete")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "PageService.Delete: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Delete: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("PageService.Delete: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *PageService) Index(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Index: marshal Request")
	}
	url := s.client.RemoteHost + "PageService.Index"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Index: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Index")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "PageService.Index: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Index: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("PageService.Index: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *PageService) Show(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Show: marshal Request")
	}
	url := s.client.RemoteHost + "PageService.Show"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Show: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Show")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "PageService.Show: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Show: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("PageService.Show: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *PageService) Update(ctx context.Context, r *Page) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Update: marshal Page")
	}
	url := s.client.RemoteHost + "PageService.Update"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Update: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Update")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "PageService.Update: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "PageService.Update: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("PageService.Update: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

type VideoService struct {
	client *Client
}

// NewVideoService makes a new client for accessing VideoService services.
func NewVideoService(client *Client) *VideoService {
	return &VideoService{
		client: client,
	}
}

func (s *VideoService) Create(ctx context.Context, r *Video) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Create: marshal Video")
	}
	url := s.client.RemoteHost + "VideoService.Create"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Create: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Create")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VideoService.Create: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Create: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VideoService.Create: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VideoService) Delete(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Delete: marshal Request")
	}
	url := s.client.RemoteHost + "VideoService.Delete"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Delete: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Delete")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VideoService.Delete: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Delete: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VideoService.Delete: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VideoService) Index(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Index: marshal Request")
	}
	url := s.client.RemoteHost + "VideoService.Index"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Index: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Index")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VideoService.Index: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Index: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VideoService.Index: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VideoService) Show(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Show: marshal Request")
	}
	url := s.client.RemoteHost + "VideoService.Show"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Show: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Show")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VideoService.Show: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Show: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VideoService.Show: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VideoService) Update(ctx context.Context, r *Video) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Update: marshal Video")
	}
	url := s.client.RemoteHost + "VideoService.Update"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Update: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Update")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VideoService.Update: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VideoService.Update: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VideoService.Update: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

type VisitService struct {
	client *Client
}

// NewVisitService makes a new client for accessing VisitService services.
func NewVisitService(client *Client) *VisitService {
	return &VisitService{
		client: client,
	}
}

func (s *VisitService) Create(ctx context.Context, r *Visit) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Create: marshal Visit")
	}
	url := s.client.RemoteHost + "VisitService.Create"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Create: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Create")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VisitService.Create: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Create: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VisitService.Create: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VisitService) Delete(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Delete: marshal Request")
	}
	url := s.client.RemoteHost + "VisitService.Delete"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Delete: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Delete")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VisitService.Delete: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Delete: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VisitService.Delete: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VisitService) Index(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Index: marshal Request")
	}
	url := s.client.RemoteHost + "VisitService.Index"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Index: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Index")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VisitService.Index: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Index: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VisitService.Index: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VisitService) Show(ctx context.Context, r *Request) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Show: marshal Request")
	}
	url := s.client.RemoteHost + "VisitService.Show"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Show: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Show")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VisitService.Show: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Show: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VisitService.Show: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

func (s *VisitService) Update(ctx context.Context, r *Visit) (*Response, error) {
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Update: marshal Visit")
	}
	url := s.client.RemoteHost + "VisitService.Update"
	s.client.Debug(fmt.Sprintf("POST %s", url))
	s.client.Debug(fmt.Sprintf(">> %s", string(requestBodyBytes)))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Update: NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req = req.WithContext(ctx)
	if s.client.BeforeRequest != nil {
		err = s.client.BeforeRequest(req)
		if err != nil {
			// don't wrap this error, it belongs to the user
			return nil, err
		}
	}
	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Update")
	}
	defer resp.Body.Close()
	var response struct {
		Response
		Error string
	}
	var bodyReader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		decodedBody, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "VisitService.Update: new gzip reader")
		}
		defer decodedBody.Close()
		bodyReader = decodedBody
	}
	respBodyBytes, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "VisitService.Update: read response body")
	}
	if err := json.Unmarshal(respBodyBytes, &response); err != nil {
		if resp.StatusCode != http.StatusOK {
			return nil, errors.Errorf("VisitService.Update: (%d) %v", resp.StatusCode, string(respBodyBytes))
		}
		return nil, err
	}
	if response.Error != "" {
		return nil, errors.New(response.Error)
	}
	return &response.Response, nil
}

// Job tracks jobs in the system
type Job struct {
	Kind string `json:"kind"`

	Args string `json:"args"`

	Status string `json:"status"`

	Queue string `json:"queue"`

	Attempts []*JobAttempt `json:"attempts"`
}

// JobAttempt tracks the attempts made to process a job
type JobAttempt struct {
	StartedAt string `json:"started_at"`

	Duration float64 `json:"duration"`

	Status string `json:"status"`

	Stacktrace []string `json:"stacktrace"`
}

// Page represents a web page to be scraped and downloaded
type Page struct {
	Name string `json:"name"`

	URL string `json:"url"`

	Scraper string `json:"scraper"`

	Downloader string `json:"downloader"`
}

type Request struct {
	ID string `json:"id"`

	Limit int `json:"limit"`

	Skip int `json:"skip"`
}

type Response struct {
	Total int64 `json:"total"`

	Results interface{} `json:"results"`
}

// Video represents a video to be downloaded
type Video struct {
	Title string `json:"title"`

	Season int `json:"season"`

	Episode int `json:"episode"`

	Raw string `json:"raw"`

	DisplayID string `json:"display_id"`

	Extension string `json:"extension"`

	Resolution int `json:"resolution"`

	Size int64 `json:"size"`

	Download string `json:"download"`

	View string `json:"view"`

	Source string `json:"source"`
}

// Visit represents a web page to be scraped and downloaded
type Visit struct {
	PageID primitive.ObjectID `json:"page_id"`

	URL string `json:"url"`
}
