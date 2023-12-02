package providusbank

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type NIPBanksResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Banks           []struct {
		Code string `json:"bankCode"`
		Name string `json:"bankName"`
	} `json:"banks"`
}

func (c *transferClient) GetNIPBanks(ctx context.Context) (data NIPBanksResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/PiPCreateDynamicAccountNumber", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type requireAuth interface {
	SetAuthCredentials(username, password string)
}

type authPayload struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

func (p *authPayload) SetAuthCredentials(username, password string) {
	p.Username = username
	p.Password = password
}

type getBvnDetailsPayload struct {
	authPayload
	BVN string `json:"bvn"`
}

type BVNDetailsResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (c *transferClient) GetBVNDetails(ctx context.Context, bvn string) (data BVNDetailsResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetBVNDetails", &getBvnDetailsPayload{BVN: bvn})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type getTransactionStatusPayload struct {
	authPayload
	Reference string `json:"transactionReference"`
}

type TransactionStatusResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	CreditAccount   string `json:"creditAccount"`
	DebitAccount    string `json:"debitAccount"`
	Reference       string `json:"transactionReference"`
	Date            Time   `json:"transactionDateTime"`
}

func (r TransactionStatusResponse) ParseAmount() float64 {
	v, _ := strconv.ParseFloat(r.Amount, 64)
	return v
}

func (c *transferClient) GetTransactionStatus(ctx context.Context, reference string) (data TransactionStatusResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetProvidusTransactionStatus", &getTransactionStatusPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
