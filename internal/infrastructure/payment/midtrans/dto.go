package midtrans

import "github.com/fkihai/payflow/internal/domain/entity"

type midtransResponse struct {
	TransactionID     string                   `json:"transaction_id"`
	OrderID           string                   `json:"order_id"`
	GrossAmount       string                   `json:"gross_amount"`
	TransactionStatus entity.TransactionStatus `json:"transaction_status"`
	Actions           []actionsResponse        `json:"actions"`
}

type actionsResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
