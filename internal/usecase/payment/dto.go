package paymentuc

import (
	"time"

	"github.com/fkihai/payflow/internal/domain"
	"github.com/google/uuid"
)

// from client
type CreateChargeRequest struct {
	Amount int64
}

type CreateChargeResult struct {
	ID     uuid.UUID
	OID    domain.OID
	Amount int64
	Status domain.ChargeStatus
	QrUrl  string
}

// to payment gateway
type GatewayRequest struct {
	OID    domain.OID
	Amount int64
}

type GatewayResult struct {
	OID       domain.OID
	Amount    int64
	Status    domain.ChargeStatus
	ExpiresAt int64
	QrUrl     string
}

type WebhookEvent struct {
	ChgID       string
	ChgStatus   domain.ChargeStatus
	ChgPaid     *time.Time
	OID         domain.OID
	GrossAmount int64
	Issuer      string
}
