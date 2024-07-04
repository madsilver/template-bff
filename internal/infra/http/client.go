package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/madsilver/template-bff/internal/domain/dto"
)

type httpClient struct {
	client *http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		client: &http.Client{Timeout: time.Second * 60},
	}
}

func (h *httpClient) Send(ctx context.Context, req *dto.Request) ([]byte, int, error) {
	buf := bytes.NewBuffer([]byte{})
	if req.Body != nil {
		buf = bytes.NewBuffer(req.Body.([]byte))
	}
	request, err := http.NewRequest(req.Method, req.URL, buf)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response, err := h.client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	if req.Entity != nil && len(respBody) > 0 {
		err = json.Unmarshal(respBody, req.Entity)
		if err != nil {
			return nil, response.StatusCode, err
		}
	}

	return respBody, response.StatusCode, err
}
