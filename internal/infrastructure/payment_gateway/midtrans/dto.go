package midtrans

import "github.com/fkihai/payflow/internal/domain/entity"

type SnapResponse struct {
	TransactionID     string                   `json:"transaction_id"`
	OrderID           string                   `json:"order_id"`
	GrossAmount       string                   `json:"gross_amount"`
	TransactionStatus entity.TransactionStatus `json:"transaction_status"`
	Actions           []ActionsResponse        `json:"actions"`
}

type ActionsResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
