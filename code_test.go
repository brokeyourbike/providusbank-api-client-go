package providusbank_test

import (
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go"
	"github.com/stretchr/testify/require"
)

func TestCode_Empty(t *testing.T) {
	c := providusbank.Code("")
	require.False(t, c.IsSuccesful())
	require.False(t, c.IsFailed())
}
