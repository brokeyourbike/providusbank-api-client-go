package webhook

import (
	"encoding/json"
	"strconv"

	"github.com/brokeyourbike/providusbank-api-client-go"
)

type Code string

const (
	CodeSuccess              Code = "00"
	CodeDuplicateTransaction Code = "01"
	CodeRejectedTransaction  Code = "02"
	CodeSystemFailure        Code = "03"
)

type RequestShort struct {
	SessionID           string            `json:"sessionId" binding:"required"`
	SettlementID        string            `json:"settlementId" binding:"required"`
	ChannelID           string            `json:"channelId"`
	AccountNumber       string            `json:"accountNumber"`
	Amount              float64           `json:"transactionAmount"`
	SettledAmount       float64           `json:"settledAmount"`
	FeeAmount           float64           `json:"feeAmount"`
	VATAmount           float64           `json:"vatAmount"`
	Currency            string            `json:"currency"`
	SourceAccountNumber string            `json:"sourceAccountNumber"`
	SourceAccountName   string            `json:"sourceAccountName"`
	SourceBankName      string            `json:"sourceBankName"`
	InitiationReference string            `json:"initiationTranRef"`
	Remarks             string            `json:"tranRemarks"`
	Date                providusbank.Time `json:"tranDateTime"`
}

func (r *RequestShort) UnmarshalJSON(data []byte) error {
	var requestShort struct {
		SessionID           string            `json:"sessionId"`
		SettlementID        string            `json:"settlementId"`
		ChannelID           string            `json:"channelId"`
		AccountNumber       string            `json:"accountNumber"`
		Amount              string            `json:"transactionAmount"`
		SettledAmount       string            `json:"settledAmount"`
		FeeAmount           string            `json:"feeAmount"`
		VATAmount           string            `json:"vatAmount"`
		Currency            string            `json:"currency"`
		SourceAccountNumber string            `json:"sourceAccountNumber"`
		SourceAccountName   string            `json:"sourceAccountName"`
		SourceBankName      string            `json:"sourceBankName"`
		InitiationReference string            `json:"initiationTranRef"`
		Remarks             string            `json:"tranRemarks"`
		Date                providusbank.Time `json:"tranDateTime"`
	}

	if err := json.Unmarshal(data, &requestShort); err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(requestShort.Amount, 64)
	if err != nil {
		return err
	}

	settledAmount, err := strconv.ParseFloat(requestShort.SettledAmount, 64)
	if err != nil {
		return err
	}

	feeAmount, err := strconv.ParseFloat(requestShort.FeeAmount, 64)
	if err != nil {
		return err
	}

	vatAmount, err := strconv.ParseFloat(requestShort.VATAmount, 64)
	if err != nil {
		return err
	}

	r.SessionID = requestShort.SessionID
	r.SettlementID = requestShort.SettlementID
	r.ChannelID = requestShort.ChannelID
	r.AccountNumber = requestShort.AccountNumber
	r.Amount = amount
	r.SettledAmount = settledAmount
	r.FeeAmount = feeAmount
	r.VATAmount = vatAmount
	r.Currency = requestShort.Currency
	r.SourceAccountNumber = requestShort.SourceAccountNumber
	r.SourceAccountName = requestShort.SourceAccountName
	r.SourceBankName = requestShort.SourceBankName
	r.InitiationReference = requestShort.InitiationReference
	r.Remarks = requestShort.Remarks
	r.Date = requestShort.Date

	return nil
}

type Response struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	SessionID         string `json:"sessionId,omitempty"`
	MesponseMessage   string `json:"responseMessage"`
	ResponseCode      Code   `json:"responseCode"`
}

func NewResponse(code Code, message string) Response {
	return Response{
		RequestSuccessful: code == CodeSuccess,
		MesponseMessage:   message,
		ResponseCode:      code,
	}
}

func NewResponseWithSession(sessionID string, code Code, message string) Response {
	return Response{
		RequestSuccessful: code == CodeSuccess,
		SessionID:         sessionID,
		MesponseMessage:   message,
		ResponseCode:      code,
	}
}
