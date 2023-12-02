package providusbank

import (
	"context"
	"fmt"
	"net/http"
)

type DynamicAccountPayload struct {
	AccountName string `json:"account_name"`
}

type CreateDynamicAccountResponse struct {
	Success         bool   `json:"requestSuccessful"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccountName     string `json:"account_name"`
	AccountNumber   string `json:"account_number"`
	Reference       string `json:"initiationTranRef"`
}

func (c *accountClient) CreateDynamicAccount(ctx context.Context, payload DynamicAccountPayload) (data CreateDynamicAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPCreateDynamicAccountNumber", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type ReservedAccountPayload struct {
	AccountName string `json:"account_name"`
	BVN         string `json:"bvn"`
}

type CreateReservedAccountResponse struct {
	Success         bool   `json:"requestSuccessful"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccountName     string `json:"account_name"`
	AccountNumber   string `json:"account_number"`
	BVN             string `json:"bvn"`
}

func (c *accountClient) CreateReservedAccount(ctx context.Context, payload ReservedAccountPayload) (data CreateReservedAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPCreateReservedAccountNumber", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type UpdateAccountNamePayload struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

type AccountOperationResponse struct {
	Success         bool   `json:"requestSuccessful"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (c *accountClient) UpdateAccountName(ctx context.Context, payload UpdateAccountNamePayload) (data AccountOperationResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPUpdateAccountName", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type BlacklistAccountPayload struct {
	AccountNumber string `json:"account_number"`
}

func (c *accountClient) BlacklistAccount(ctx context.Context, payload BlacklistAccountPayload) (data AccountOperationResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPBlacklistAccount", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type VerifyTransactionResponse struct {
	SessionID           string  `json:"sessionId"`
	SettlementID        string  `json:"settlementId"`
	ChannelID           string  `json:"channelId"`
	AccountNumber       string  `json:"accountNumber"`
	SourceAccountNumber string  `json:"sourceAccountNumber"`
	SourceAccountName   string  `json:"sourceAccountName"`
	SourceBankName      string  `json:"sourceBankName"`
	Currency            string  `json:"currency"`
	Amount              float64 `json:"transactionAmount"`
	SettledAmount       float64 `json:"settledAmount"`
	FeeAmount           float64 `json:"feeAmount"`
	VATAmount           float64 `json:"vatAmount"`
	InitiationReference string  `json:"initiationTranRef"`
	Remarks             string  `json:"tranRemarks"`
	Date                Time    `json:"tranDateTime"`
}

func (c *accountClient) VerifyTransaction(ctx context.Context, sessionID string) (data VerifyTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/PiPverifyTransaction", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddQueryParam("session_id", sessionID)

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

func (c *accountClient) VerifyTransactionWithSettlementID(ctx context.Context, settlementID string) (data VerifyTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/api/PiPverifyTransaction_settlementid", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddQueryParam("settlement_id", settlementID)

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type RepushTransactionPayload struct {
	SessionID    string `json:"session_id"`
	SettlementID string `json:"settlement_id"`
}

type RepushTransactionResponse struct {
	Success             bool   `json:"requestSuccessful"`
	ResponseCode        string `json:"responseCode"`
	ResponseMessage     string `json:"responseMessage"`
	AccountNumber       string `json:"account_number"`
	AccountName         string `json:"account_name"`
	InitiationReference string `json:"initiationTranRef"`
}

func (c *accountClient) RepushTransaction(ctx context.Context, payload RepushTransactionPayload) (data RepushTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiP_RepushTransaction_SettlementId", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
