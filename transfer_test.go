package providusbank_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/exception.json
var exception []byte

//go:embed testdata/GetNIPBanks-success.json
var getNIPBanksSuccess []byte

//go:embed testdata/TransactionStatus-success.json
var transactionStatusSuccess []byte

//go:embed testdata/FundTransfer-fail.json
var fundTransferFail []byte

func TestGetNIPBanks_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetNIPBanks(nil) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestGetNIPBanks_Err(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusInternalServerError, Body: io.NopCloser(bytes.NewReader(exception))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	_, err := client.GetNIPBanks(context.TODO())
	require.Error(t, err)

	got, ok := err.(providusbank.ErrResponse)
	require.True(t, ok)
	require.Equal(t, 500, got.Status)
}

func TestGetNIPBanks_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)

	logger, hook := logrustest.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient), providusbank.WithLogger(logger))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(getNIPBanksSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.GetNIPBanks(context.TODO())
	require.NoError(t, err)

	assert.Len(t, got.Banks, 2)

	require.Equal(t, 2, len(hook.Entries))
	require.Contains(t, hook.Entries[0].Data, "http.request.method")
	require.Contains(t, hook.Entries[0].Data, "http.request.url")
	require.Contains(t, hook.Entries[0].Data, "http.request.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.status_code")
	require.Contains(t, hook.Entries[1].Data, "http.response.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.headers")
}

func TestGetBVNDetails_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetBVNDetails(nil, "bvn") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestGetBVNDetails_AuthFailed(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountAuthFailed))}
	mockHttpClient.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		b, err := io.ReadAll(req.Body)
		require.NoError(t, err)

		// username and passoword passed as part of request
		return string(b) == `{"userName":"john","password":"pass","bvn":"bvn"}`
	})).Return(resp, nil).Once()

	got, err := client.GetBVNDetails(context.TODO(), "bvn")
	require.NoError(t, err)

	assert.Equal(t, providusbank.CodeDoNotHonor, got.ResponseCode)
}

func TestGetTransactionStatus_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetTransactionStatus(nil, "reference") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestGetTransactionStatus_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(transactionStatusSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.GetTransactionStatus(context.TODO(), "ab003ed9")
	require.NoError(t, err)

	assert.Equal(t, "2.50", got.Amount)
}

func TestGetNIPTransactionStatus_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetNIPTransactionStatus(nil, "reference") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestGetAccount_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetAccount(nil, "number") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestGetNIPAccount_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.GetNIPAccount(nil, "bank_code", "number") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestFundTransfer_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.FundTransfer(nil, providusbank.FundTransferPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestFundTransfer_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(fundTransferFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.FundTransfer(context.TODO(), providusbank.FundTransferPayload{})
	require.NoError(t, err)

	assert.NotEqual(t, providusbank.CodeCompleted, got.ResponseCode)
}

func TestNIPFundTransfer_RequestErr(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("baseurl", "username", "password", providusbank.WithHTTPClient(mockHttpClient))

	_, err := client.NIPFundTransfer(nil, providusbank.NIPFundTransferPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}
