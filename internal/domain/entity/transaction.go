package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	TransactionPending    TransactionStatus = "PENDING"
	TransactionSettlement TransactionStatus = "SETTLEMENT"
	TransactionExpired    TransactionStatus = "EXPIRED"
	TransactionFailed     TransactionStatus = "FAILED"
)

type Transaction struct {
	ID        uuid.UUID         `db:"id"`
	OrderId   string            `db:"order_id"`
	Amount    int               `db:"amount"`
	Status    TransactionStatus `db:"status"`
	ExpiresAt int               `db:"expires_at"`
	PaidAt    *time.Time        `db:"paid_at"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}
