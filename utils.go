package facebook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSecretProof(token, secret string) string {
	key := []byte(secret)
	message := token

	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))

	return hex.EncodeToString(sig.Sum(nil))
}

func isEmptyString(param string) bool {
	if param != "" {
		return false
	}

	return true
}
