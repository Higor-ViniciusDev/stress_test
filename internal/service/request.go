package service

import (
	"context"
	"net/http"
)

type ServiceRequest struct {
}

func (s *ServiceRequest) FazerRequest(ctx context.Context, url string) map[string]interface{} {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return map[string]interface{}{
			"status": false,
			"Code":   0,
		}
	}
	defer req.Body.Close()

	respons, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{
			"status": false,
			"Code":   0,
		}
	}

	return map[string]interface{}{
		"status": true,
		"Code":   respons.StatusCode,
	}
}
