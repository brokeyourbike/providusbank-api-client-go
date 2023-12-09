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

//go:embed testdata/CreateDynamicAccount-auth-failed.json
var createDynamicAccountAuthFailed []byte

//go:embed testdata/CreateDynamicAccount-success.json
var createDynamicAccountSuccess []byte

//go:embed testdata/CreateReservedAccount-success.json
var createReservedAccountSuccess []byte

//go:embed testdata/UpdateAccountName-fail.json
var updateAccountNameFail []byte

//go:embed testdata/UpdateAccountName-success.json
var updateAccountNameSuccess []byte

//go:embed testdata/BlacklistAccount-fail.json
var blacklistAccountFail []byte

//go:embed testdata/BlacklistAccount-success.json
var blacklistAccountSuccess []byte

//go:embed testdata/VerifyTransaction-fail.json
var verifyTransactionFail []byte

//go:embed testdata/VerifyTransaction-success.json
var verifyTransactionSuccess []byte

//go:embed testdata/RepushTransaction-fail.json
var repushTransactionFail []byte

func TestCreateDynamicAccount_AuthFailed(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountAuthFailed))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateDynamicAccount(context.TODO(), "name")
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestCreateDynamicAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)

	logger, hook := logrustest.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient), providusbank.WithLogger(logger))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createDynamicAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateDynamicAccount(context.TODO(), "name")
	require.NoError(t, err)

	assert.True(t, got.Success)

	require.Equal(t, 2, len(hook.Entries))
	require.Contains(t, hook.Entries[0].Data, "http.request.method")
	require.Contains(t, hook.Entries[0].Data, "http.request.url")
	require.Contains(t, hook.Entries[0].Data, "http.request.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.status_code")
	require.Contains(t, hook.Entries[1].Data, "http.response.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.headers")
}

func TestCreateReservedAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(createReservedAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateReservedAccount(context.TODO(), "name", "bvn")
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestUpdateAccountName_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(updateAccountNameFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.UpdateAccountName(context.TODO(), "number", "name")
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestUpdateAccountName_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(updateAccountNameSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.UpdateAccountName(context.TODO(), "number", "name")
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestBlacklistAccount_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(blacklistAccountFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.BlacklistAccount(context.TODO(), "number")
	require.NoError(t, err)

	assert.False(t, got.Success)
}

func TestBlacklistAccount_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(blacklistAccountSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.BlacklistAccount(context.TODO(), "number")
	require.NoError(t, err)

	assert.True(t, got.Success)
}

func TestVerifyTransaction_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransaction(context.TODO(), "")
	require.NoError(t, err)

	assert.Equal(t, "", got.SessionID)
}

func TestVerifyTransaction_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransaction(context.TODO(), "123456789")
	require.NoError(t, err)

	assert.Equal(t, "123456789", got.SessionID)
}

func TestVerifyTransactionWithSettlementID_Success(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(verifyTransactionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.VerifyTransactionWithSettlementID(context.TODO(), "204210202000000700001")
	require.NoError(t, err)

	assert.Equal(t, "204210202000000700001", got.SettlementID)
}

func TestRepushTransaction_Fail(t *testing.T) {
	mockHttpClient := providusbank.NewMockHttpClient(t)
	client := providusbank.NewAccountClient("a.com", "john", "pass", providusbank.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(repushTransactionFail))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.RepushTransaction(context.TODO(), "session_id", "settlement_id")
	require.NoError(t, err)

	assert.False(t, got.Success)
}
