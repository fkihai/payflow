package repository

import (
	"context"

	e "github.com/fkihai/payflow/internal/domain/entity"
)

type TransactionRepository interface {
	Create(context context.Context, trx e.Transaction) (*e.Transaction, error)
	Update(context context.Context, trx e.Transaction) (*e.Transaction, error)
}
