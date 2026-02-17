package payment

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/fkihai/payflow/internal/domain/payment"
)

type QrisOrderIDGenerator struct{}

func NewQriOrderIDGenerator() payment.OrderIDGenerator {
	return &QrisOrderIDGenerator{}
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func (g *QrisOrderIDGenerator) Generate() (payment.OrderID, error) {
	timestamp := time.Now().Format("20060102150405")

	randomStr, err := randomBase64(8)

	if err != nil {
		return "", nil
	}

	return payment.NewOrderID(timestamp, randomStr)
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
