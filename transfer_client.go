package providusbank

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type TransferClient interface {
	GetNIPBanks(ctx context.Context) (data NIPBanksResponse, err error)
	GetBVNDetails(ctx context.Context, bvn string) (data BVNDetailsResponse, err error)
	GetTransactionStatus(ctx context.Context, reference string) (data TransactionStatusResponse, err error)
	GetNIPTransactionStatus(ctx context.Context, reference string) (data NIPTransactionStatusResponse, err error)
	GetAccount(ctx context.Context, accountNumber string) (data AccountResponse, err error)
	GetNIPAccount(ctx context.Context, bankCode, accountNumber string) (data NIPAccountResponse, err error)
	FundTransfer(ctx context.Context, payload FundTransferPayload) (data FundTransferResponse, err error)
	NIPFundTransfer(ctx context.Context, payload NIPFundTransferPayload) (data NIPFundTransferResponse, err error)
}

var _ TransferClient = (*transferClient)(nil)

type transferClient struct {
	client
	baseURL  string
	username string
	password string
}

func NewTransferClient(baseURL, username, password string, options ...ClientOption) *transferClient {
	c := &transferClient{
		baseURL:  strings.TrimSuffix(baseURL, "/"),
		username: username,
		password: password,
	}

	c.httpClient = http.DefaultClient

	for _, option := range options {
		option(&c.client)
	}

	return c
}

func (c *transferClient) newRequest(ctx context.Context, method, url string, body interface{}) (*request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if auth, ok := body.(requireAuth); ok {
		auth.SetAuthCredentials(c.username, c.password)
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

	req.Header.Set("Accept-Encoding", "identity")
	return NewRequest(req), nil
}
