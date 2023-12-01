package providusbank_test

import (
	_ "embed"
	"testing"
)

//go:embed testdata/CreateDynamicAccount-auth-failed.json
var createDynamicAccountAuthFailed []byte

//go:embed testdata/CreateDynamicAccount-success.json
var createDynamicAccountSuccess []byte

func TestCreateDynamicAccount_AuthFailed(t *testing.T) {}

func TestCreateDynamicAccount_Success(t *testing.T) {}
