package providusbank

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type AccountClient interface {
}

var _ AccountClient = (*accountClient)(nil)

type accountClient struct {
	client
	baseURL string
	token   string
	secret  string
}

func NewAccountClient(baseURL, token, secret string, options ...ClientOption) *accountClient {
	c := &accountClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		token:   token,
		secret:  secret,
	}

	c.httpClient = http.DefaultClient

	for _, option := range options {
		option(&c.client)
	}

	return c
}

func (c *accountClient) newRequest(ctx context.Context, method, url string, body interface{}) (*request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var b []byte

	switch req.Method {
	case http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete:
		b, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		req.Body = io.NopCloser(bytes.NewReader(b))
	}

	if c.logger != nil {
		c.logger.WithContext(ctx).WithFields(logrus.Fields{
			"http.request.method":       req.Method,
			"http.request.url":          req.URL.String(),
			"http.request.body.content": string(b),
		}).Debug("providusbank.client -> request")
	}

	signature := sha512.Sum512([]byte(fmt.Sprintf("%s:%s", c.token, c.secret)))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", c.token)
	req.Header.Set("X-Auth-Signature", fmt.Sprintf("%x", signature))
	return NewRequest(req), nil
}

func (c *accountClient) do(ctx context.Context, req *request) error {
	resp, err := c.httpClient.Do(req.req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(b))

	if c.logger != nil {
		c.logger.WithContext(ctx).WithFields(logrus.Fields{
			"http.response.status_code":  resp.StatusCode,
			"http.response.body.content": string(b),
			"http.response.headers":      resp.Header,
		}).Debug("providusbank.client -> response")
	}

	if req.decodeTo != nil {
		if err := json.NewDecoder(resp.Body).Decode(req.decodeTo); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
