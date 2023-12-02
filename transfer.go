package providusbank

import (
	"context"
	"fmt"
	"net/http"
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

func (c *transferClient) GetTransactionStatus(ctx context.Context, reference string) (data TransactionStatusResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetProvidusTransactionStatus", &getTransactionStatusPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

func (c *transferClient) GetNIPTransactionStatus(ctx context.Context, reference string) (data TransactionStatusResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetNIPTransactionStatus", &getTransactionStatusPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type getAccountPayload struct {
	authPayload
	AccountNumber string `json:"accountNumber"`
}

type AccountResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (c *transferClient) GetAccount(ctx context.Context, accountNumber string) (data AccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetProvidusAccount", &getAccountPayload{AccountNumber: accountNumber})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type getNIPAccountPayload struct {
	authPayload
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"beneficiaryBank"`
}

type NIPAccountResponse struct {
	ResponseCode         string `json:"responseCode"`
	ResponseMessage      string `json:"responseMessage"`
	AccountNumber        string `json:"accountNumber"`
	AccountName          string `json:"accountName"`
	BankCode             string `json:"bankCode"`
	BVN                  string `json:"bvn"`
	TransactionReference string `json:"transactionReference"`
}

func (c *transferClient) GetNIPAccount(ctx context.Context, bankCode, accountNumber string) (data NIPAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/GetNIPAccount", &getNIPAccountPayload{BankCode: bankCode, AccountNumber: accountNumber})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type fundTransferPayload struct {
	authPayload
	FundTransferPayload
}

type FundTransferPayload struct {
	CreditAccount string `json:"creditAccount"`
	DebitAccount  string `json:"debitAccount"`
	Currency      string `json:"currencyCode"`
	Amount        string `json:"transactionAmount"`
	Reference     string `json:"transactionReference"`
	Narration     string `json:"narration"`
}

type FundTransferResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	Reference       string `json:"transactionReference"`
}

func (c *transferClient) FundTransfer(ctx context.Context, payload FundTransferPayload) (data FundTransferResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/ProvidusFundTransfer", &fundTransferPayload{FundTransferPayload: payload})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type nipFundTransferPayload struct {
	authPayload
	NIPFundTransferPayload
}

type NIPFundTransferPayload struct {
	SourceAccountName        string `json:"sourceAccountName"`
	BeneficiaryAccountName   string `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
	BeneficiaryBank          string `json:"beneficiaryBank"`
	Currency                 string `json:"currencyCode"`
	Amount                   string `json:"transactionAmount"`
	Reference                string `json:"transactionReference"`
	Narration                string `json:"narration"`
}

type NIPFundTransferResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Reference       string `json:"transactionReference"`
	SessionID       string `json:"sessionId"`
}

func (c *transferClient) NIPFundTransfer(ctx context.Context, payload NIPFundTransferPayload) (data NIPFundTransferResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/api/NIPFundTransfer", &nipFundTransferPayload{NIPFundTransferPayload: payload})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
