package facebook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func GenerateSecretProof(token, secret string) string {
	key := []byte(secret)
	message := token

	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))

	return hex.EncodeToString(sig.Sum(nil))
}

func isEmptyString(param string) bool {
	if len(param) < 1 {
		return false
	}
	return true
}

func isErrorResponse(jsonBytes []byte, resp *http.Response) (ErrorResponse, bool) {
	var tempResult ErrorResponse
	json.Unmarshal(jsonBytes, &tempResult)

	if (tempResult == ErrorResponse{}) {
		return ErrorResponse{}, false
	}

	return tempResult, true
}
