package providusbank

import (
	"context"
	"fmt"
	"net/http"
)

type NIPBanksResponse struct {
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Banks           []struct {
		Code string `json:"bankCode"`
		Name string `json:"bankName"`
	} `json:"banks"`
}

// GetNIPBanks returns the list of institutions currently enrolled on NIP and their respective NIP bank codes.
func (c *transferClient) GetNIPBanks(ctx context.Context) (data NIPBanksResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/PiPCreateDynamicAccountNumber", nil)
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
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	FirstName       string `json:"firstName"`
	MiddleName      string `json:"middleName"`
	LastName        string `json:"surname"`
	DOB             Time   `json:"dateOfBirth"`
}

// GetBVNDetails validates the supplied single BVN and returns the full demography details associated with the BVN.
func (c *transferClient) GetBVNDetails(ctx context.Context, bvn string) (data BVNDetailsResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/GetBVNDetails", &getBvnDetailsPayload{BVN: bvn})
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
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	CreditAccount   string `json:"creditAccount"`
	DebitAccount    string `json:"debitAccount"`
	Reference       string `json:"transactionReference"`
	Date            Time   `json:"transactionDateTime"`
}

// GetTransactionStatus validates the supplied single transaction reference and returns the current status of the transaction.
// This status is of Providus-to-Providus transactions.
func (c *transferClient) GetTransactionStatus(ctx context.Context, reference string) (data TransactionStatusResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/GetProvidusTransactionStatus", &getTransactionStatusPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type NIPTransactionStatusResponse struct {
	ResponseCode           Code   `json:"responseCode"`
	ResponseMessage        string `json:"responseMessage"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	RecipientBankCode      string `json:"recipientBankCode"`
	RecipientAccountNumber string `json:"recipientAccountNumber"`
	Reference              string `json:"transactionReference"`
	Date                   Time   `json:"transactionDateTime"`
}

// GetNIPTransactionStatus validates the supplied single transaction reference and returns the current status of the transaction.
func (c *transferClient) GetNIPTransactionStatus(ctx context.Context, reference string) (data NIPTransactionStatusResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/GetNIPTransactionStatus", &getTransactionStatusPayload{Reference: reference})
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
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccountNumber   string `json:"accountNumber"`
	AccountName     string `json:"accountName"`
	AccountStatus   string `json:"accountStatus"`
	Balance         string `json:"availableBalance"`
	Email           string `json:"emailAddress"`
	Phone           string `json:"phoneNumber"`
	BVN             string `json:"bvn"`
}

// GetAccount returns the details tied to your account including the balance.
// This account is the one tied to the username making the call.
func (c *transferClient) GetAccount(ctx context.Context, accountNumber string) (data AccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/GetProvidusAccount", &getAccountPayload{AccountNumber: accountNumber})
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
	ResponseCode         Code   `json:"responseCode"`
	ResponseMessage      string `json:"responseMessage"`
	AccountNumber        string `json:"accountNumber"`
	AccountName          string `json:"accountName"`
	BankCode             string `json:"bankCode"`
	BVN                  string `json:"bvn"`
	TransactionReference string `json:"transactionReference"`
}

// GetNIPAccount validates the supplied account number and 3-digit bank code and returns the account details.
// It can also accept the 6-digit NIP bank code in place of the 3-digit.
func (c *transferClient) GetNIPAccount(ctx context.Context, bankCode, accountNumber string) (data NIPAccountResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/GetNIPAccount", &getNIPAccountPayload{BankCode: bankCode, AccountNumber: accountNumber})
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
	CreditAccount string  `json:"creditAccount"`
	DebitAccount  string  `json:"debitAccount"`
	Currency      string  `json:"currencyCode"`
	Amount        float64 `json:"transactionAmount"`
	Reference     string  `json:"transactionReference"`
	Narration     string  `json:"narration"`
}

type FundTransferResponse struct {
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	Reference       string `json:"transactionReference"`
}

// FundTransfer used to transfer fund from a specified Providus account number to another ProvidusBank account.
func (c *transferClient) FundTransfer(ctx context.Context, payload FundTransferPayload) (data FundTransferResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/ProvidusFundTransfer", &fundTransferPayload{FundTransferPayload: payload})
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
	SourceAccountName        string  `json:"sourceAccountName"`
	BeneficiaryAccountName   string  `json:"beneficiaryAccountName"`
	BeneficiaryAccountNumber string  `json:"beneficiaryAccountNumber"`
	BeneficiaryBank          string  `json:"beneficiaryBank"`
	Currency                 string  `json:"currencyCode"`
	Amount                   float64 `json:"transactionAmount"`
	Reference                string  `json:"transactionReference"`
	Narration                string  `json:"narration"`
}

type NIPFundTransferResponse struct {
	ResponseCode    Code   `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Reference       string `json:"transactionReference"`
	SessionID       string `json:"sessionId"`
}

// NIPFundTransfer used to transfer fund from a specified Providus account number to another account in a different bank.
func (c *transferClient) NIPFundTransfer(ctx context.Context, payload NIPFundTransferPayload) (data NIPFundTransferResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/NIPFundTransfer", &nipFundTransferPayload{NIPFundTransferPayload: payload})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
