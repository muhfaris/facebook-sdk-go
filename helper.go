package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// HashHmac is generate string
func HashHmac(secret string, data []byte) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write(data)

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
