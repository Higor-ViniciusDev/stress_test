package service

import (
	"context"
	"net/http"
	"time"
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
		Client: &http.Client{
			//Adicionado esses parametros, pois os primeiros teste indicou altos ms para concluir, mesmo tendo asta goroutines
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
				DisableKeepAlives:   false,
				DisableCompression:  false,
			},
		},
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
