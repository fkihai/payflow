package payment

import (
	"github.com/fkihai/payflow/internal/domain/entity"
	"github.com/fkihai/payflow/internal/domain/payment"
)

type GatewayResult struct {
	OrderID   string
	Amount    int64
	Status    entity.TransactionStatus
	QrCodeUrl string
}

type GatewayRequest struct {
	OrderID payment.OrderID
	Amount  int64
}

type CreateTransactionInput struct {
	Amount int64
}
