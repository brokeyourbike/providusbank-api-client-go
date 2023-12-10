package signature

import (
	"context"
	"crypto/sha512"
	"fmt"
)

type Generator interface {
	// Generate, generates a new signature.
	Generate(ctx context.Context) string
}

type Verifier interface {
	// Verify verifies the given message against signature.
	Verify(ctx context.Context, message string) bool
}

type signer struct {
	token  string
	secret string
}

func NewSigner(token, secret string) *signer {
	return &signer{token: token, secret: secret}
}

func (s *signer) Generate(ctx context.Context) string {
	signature := sha512.Sum512([]byte(fmt.Sprintf("%s:%s", s.token, s.secret)))
	return fmt.Sprintf("%x", signature)
}

func (s *signer) Verify(ctx context.Context, message string) bool {
	return message == s.Generate(ctx)
}
