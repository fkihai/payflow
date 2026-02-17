package paymentuc

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/fkihai/payflow/internal/domain"
)

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type QrisOID struct{}

func NewQriOrderIDGenerator() OIDGenerator {
	return &QrisOID{}
}

func (g *QrisOID) Generate() (domain.OID, error) {
	timestamp := time.Now().Format("20060102150405")

	randomStr, err := randomBase64(8)

	if err != nil {
		return "", nil
	}

	return domain.NewOID(timestamp, randomStr)
}

func randomBase64(length int) (string, error) {
	result := make([]byte, length)
	for i := range length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62))))
		if err != nil {
			return "", err
		}
		result[i] = base62[num.Int64()]
	}
	return string(result), nil
}
