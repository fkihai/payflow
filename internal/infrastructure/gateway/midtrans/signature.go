package midtrans

import (
	"crypto/sha512"
	"encoding/hex"
)

func generateSignature(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}
