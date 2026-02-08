package postgres

import (
	"database/sql"

	e "github.com/fkihai/payflow/internal/domain/entity"
	r "github.com/fkihai/payflow/internal/domain/repository"
)

type TransactionRepository struct {
	db *sql.DB
}

func (trxRepo *TransactionRepository) Create(trx e.Transaction) error {
	return nil
}
func (trxRepo *TransactionRepository) Update(trx e.Transaction) error {
	return nil
}

func NewTransactionRepositoy(db *sql.DB) r.TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
