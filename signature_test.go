package checkr

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidSignature(t *testing.T) {
	testBody := "test-body"
	req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(testBody)))
	require.NoError(t, err)

	// key: test-body
	// generated https://www.freeformatter.com/hmac-generator.html#ad-output
	signature := "92de73173e5e3e69c275c9407caf617f2006c0b02fa2db4b98ed47cff0db2b87"
	apiKey := "1"

	req.Header.Add("X-Checkr-Signature", signature)

	valid, err := IsValidSignature(req, []byte(apiKey))
	require.NoError(t, err)
	require.True(t, valid)
}
