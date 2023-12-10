package signature_test

import (
	"context"
	"testing"

	"github.com/brokeyourbike/providusbank-api-client-go/signature"
	"github.com/stretchr/testify/assert"
)

func TestSigner(t *testing.T) {
	want := "2ea3a61a94b44ec66a2c02e2ef7351612470451ce5441f89341f0f331844d130d4207cd9638d04415ed61f66368c99e684f1b1b34edde1aeec98787d4d902e15"

	s := signature.NewSigner("123", "456")

	assert.Equal(t, want, s.Generate(context.TODO()))
	assert.True(t, s.Verify(context.TODO(), want))
	assert.False(t, s.Verify(context.TODO(), "iamnotsignature"))
}
