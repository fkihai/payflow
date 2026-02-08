package repository

import e "github.com/fkihai/payflow/internal/domain/entity"

type TransactionRepository interface {
	Create(trx e.Transaction) error
	Update(trx e.Transaction) error
}
