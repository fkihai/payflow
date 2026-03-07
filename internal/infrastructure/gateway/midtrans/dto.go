package midtrans

import (
	"github.com/fkihai/payflow/internal/domain"
)

type Response struct {
	TrxID       string              `json:"transaction_id"`
	OID         domain.OID          `json:"order_id"`
	GrossAmount string              `json:"gross_amount"`
	TrxStatus   domain.ChargeStatus `json:"transaction_status"`
	Actions     []actionsResponse   `json:"actions"`
}

type actionsResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type midtransWebhookEvent struct {
	TrxID       string              `json:"transaction_id"`
	OID         domain.OID          `json:"order_id"`
	TrxStatus   domain.ChargeStatus `json:"transaction_status"`
	StatusCode  string              `json:"status_code"`
	GrossAmount string              `json:"gross_amount"`
	Signature   string              `json:"signature_key"`
	Issuer      string              `json:"issuer"`
	PaidAt      string              `json:"settlement_time"`
}
