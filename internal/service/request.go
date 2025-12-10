package service

import (
	"context"
	"net/http"
)

type ServiceRequest struct {
	Client *http.Client
}

type RetornoRequest struct {
	Code  int
	Error error
}

func NewServiceRequest() *ServiceRequest {
	return &ServiceRequest{
		Client: &http.Client{},
	}
}

func (s *ServiceRequest) FazerRequest(ctx context.Context, url string) *RetornoRequest {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return &RetornoRequest{Code: 0, Error: err}
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return &RetornoRequest{Code: 0, Error: err}
	}
	defer resp.Body.Close()

	return &RetornoRequest{
		Code:  resp.StatusCode,
		Error: nil,
	}
}
