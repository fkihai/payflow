package handler

import (
	"time"

	"github.com/fkihai/payflow/internal/domain"
	"github.com/google/uuid"
)

type ChargerResponse struct {
	ID        uuid.UUID           `json:"id" binding:"required"`
	OID       domain.OID          `json:"order_id" binding:"required"`
	Amount    int64               `json:"amount" binding:"required"`
	Status    domain.ChargeStatus `json:"status" binding:"required"`
	QrUrl     string              `json:"qr_url" binding:"required"`
	ExpiresAt time.Time           `json:"expires_at" binding:"required"`
	PaidAt    time.Time           `json:"paid_at" binding:"required"`
}

type ChargeRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}
