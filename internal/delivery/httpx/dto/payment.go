package dto

import (
	"time"

	"github.com/fkihai/payflow/internal/domain/entity"
	"github.com/google/uuid"
)

type CreatePaymentResponse struct {
	ID        uuid.UUID                `json:"id" binding:"required"`
	OrderID   string                   `json:"order_id" binding:"required"`
	Amount    int64                    `json:"amount" binding:"required"`
	Status    entity.TransactionStatus `json:"status" binding:"required"`
	QrUrl     string                   `json:"qr_url" binding:"required"`
	ExpiresAt time.Time                `json:"expires_at" binding:"required"`
	PaidAt    time.Time                `json:"paid_at" binding:"required"`
}

type CreatePaymentRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}
