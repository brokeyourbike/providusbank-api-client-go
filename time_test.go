package providusbank_test

import (
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"unsupported", "2023-07-21", true},
		{"VerifyTransaction", "", false},
		{"VerifyTransaction", "2/11/2021 6:08:34 PM", false},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var d providusbank.Time

			err := d.UnmarshalJSON([]byte(test.value))
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestZeroTime(t *testing.T) {
	var d providusbank.Time
	require.NoError(t, d.UnmarshalJSON([]byte("")))
	require.True(t, d.IsZero())
}