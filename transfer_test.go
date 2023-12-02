package providusbank_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go"
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

func TestGetNIPBanks_Err(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

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
	client := providusbank.NewTransferClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(getNIPBanksSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.GetNIPBanks(context.TODO())
	require.NoError(t, err)

	assert.Len(t, got.Banks, 2)
}

func TestGetBVNDetails_AuthFailed(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountAuthFailed))}
	mockHttpClient.On("Do", mock.MatchedBy(func(req *http.Request) bool {
		b, err := io.ReadAll(req.Body)
		require.NoError(t, err)

		return string(b) == `{"userName":"john","password":"pass","bvn":"bvn"}`
	})).Return(resp, nil).Once()

	got, err := client.GetBVNDetails(context.TODO(), "bvn")
	require.NoError(t, err)

	assert.Equal(t, "05", got.ResponseCode)
}

func TestGetTransactionStatus_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewTransferClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(transactionStatusSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.GetTransactionStatus(context.TODO(), "ab003ed9")
	require.NoError(t, err)

	assert.Equal(t, "2.50", got.Amount)
}
