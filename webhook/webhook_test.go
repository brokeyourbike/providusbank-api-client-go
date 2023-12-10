package webhook_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go/webhook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/webhook-1.json
var webhook1 []byte

func TestWebhookResponse(t *testing.T) {
	r1 := webhook.NewResponse("123", webhook.CodeSuccess, "yes")
	assert.True(t, r1.RequestSuccessful)

	r2 := webhook.NewResponse("123", webhook.CodeSystemFailure, "yes")
	assert.False(t, r2.RequestSuccessful)
}

func TestWebhookRequestShort(t *testing.T) {
	var r1 webhook.RequestShort
	require.NoError(t, json.Unmarshal(webhook1, &r1))

	assert.Equal(t, 100.0, r1.Amount)
	assert.Equal(t, 100.0, r1.SettledAmount)
	assert.Equal(t, 1.2, r1.FeeAmount)
	assert.Equal(t, 0.0, r1.VATAmount)
}
