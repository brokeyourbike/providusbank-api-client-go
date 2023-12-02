package providusbank

import (
	"context"
	"fmt"
	"net/http"
)

type dynamicAccountPayload struct {
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

func (c *accountClient) CreateDynamicAccount(ctx context.Context, accountName string) (data CreateDynamicAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPCreateDynamicAccountNumber", dynamicAccountPayload{AccountName: accountName})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type reservedAccountPayload struct {
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

func (c *accountClient) CreateReservedAccount(ctx context.Context, accountName, bvn string) (data CreateReservedAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPCreateReservedAccountNumber", reservedAccountPayload{AccountName: accountName, BVN: bvn})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type updateAccountNamePayload struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

type AccountOperationResponse struct {
	Success         bool   `json:"requestSuccessful"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (c *accountClient) UpdateAccountName(ctx context.Context, accountNumber, accountName string) (data AccountOperationResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPUpdateAccountName", updateAccountNamePayload{AccountNumber: accountNumber, AccountName: accountName})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type blacklistAccountPayload struct {
	AccountNumber string `json:"account_number"`
}

func (c *accountClient) BlacklistAccount(ctx context.Context, accountNumber string) (data AccountOperationResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiPBlacklistAccount", blacklistAccountPayload{AccountNumber: accountNumber})
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

type repushTransactionPayload struct {
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

func (c *accountClient) RepushTransaction(ctx context.Context, sessionID, settlementID string) (data RepushTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/PiP_RepushTransaction_SettlementId", repushTransactionPayload{SessionID: sessionID, SettlementID: settlementID})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
