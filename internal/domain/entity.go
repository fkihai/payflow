package domain

import (
	"time"

	"github.com/google/uuid"
)

type ChargeStatus string

const (
	ChargePending    ChargeStatus = "PENDING"
	ChargeSettlement ChargeStatus = "SETTLEMENT"
	ChargeExpired    ChargeStatus = "EXPIRED"
	ChargeFailed     ChargeStatus = "FAILED"
)

type Charge struct {
	ID        uuid.UUID    `db:"id"`
	OID       OID          `db:"order_id"`
	Amount    int64        `db:"amount"`
	Status    ChargeStatus `db:"status"`
	ExpiresAt int          `db:"expires_at"`
	PaidAt    *time.Time   `db:"paid_at"`
	QrUrl     string       `db:"qr_url"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
