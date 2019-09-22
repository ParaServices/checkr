package checkr

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var ErrCheckrHeaderSignatureNotFound = errors.New("Checkr header X-Checkr-Signature not found")

func IsValidSignature(req *http.Request, apiKey []byte) (bool, error) {
	signature := req.Header.Get("X-Checkr-Signature")
	if signature == "" {
		return false, ErrCheckrHeaderSignatureNotFound
	}

	h := hmac.New(sha256.New, apiKey)

	var buf bytes.Buffer
	_, err := io.Copy(&buf, io.TeeReader(req.Body, h))
	if err != nil {
		return false, err
	}
	defer req.Body.Close()
	req.Body = ioutil.NopCloser(&buf)

	digest := hex.EncodeToString(h.Sum(nil))

	return signature == digest, nil
}
