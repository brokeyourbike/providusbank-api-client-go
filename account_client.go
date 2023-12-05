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
	CreateDynamicAccount(ctx context.Context, accountName string) (data CreateDynamicAccountResponse, err error)
	CreateReservedAccount(ctx context.Context, accountName, bvn string) (data CreateReservedAccountResponse, err error)
	UpdateAccountName(ctx context.Context, accountNumber, accountName string) (data AccountOperationResponse, err error)
	BlacklistAccount(ctx context.Context, accountNumber string) (data AccountOperationResponse, err error)
	VerifyTransaction(ctx context.Context, sessionID string) (data VerifyTransactionResponse, err error)
	VerifyTransactionWithSettlementID(ctx context.Context, settlementID string) (data VerifyTransactionResponse, err error)
	RepushTransaction(ctx context.Context, sessionID, settlementID string) (data RepushTransactionResponse, err error)
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

	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		req.Body = io.NopCloser(bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
	}

	if c.logger != nil {
		c.logger.WithContext(ctx).WithFields(logrus.Fields{
			"http.request.method":       req.Method,
			"http.request.url":          req.URL.String(),
			"http.request.body.content": string(b),
		}).Debug("providusbank.client -> request")
	}

	signature := sha512.Sum512([]byte(fmt.Sprintf("%s:%s", c.token, c.secret)))

	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Client-Id", c.token)
	req.Header.Set("X-Auth-Signature", fmt.Sprintf("%x", signature))
	return NewRequest(req), nil
}
