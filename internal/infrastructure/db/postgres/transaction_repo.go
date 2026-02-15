package postgres

import (
	"context"
	"database/sql"

	e "github.com/fkihai/payflow/internal/domain/entity"
	r "github.com/fkihai/payflow/internal/domain/repository"
)

type PostgresTransactionRepository struct {
	db *sql.DB
}

func (trxRepo *PostgresTransactionRepository) Create(ctx context.Context, trx e.Transaction) (*e.Transaction, error) {
	q := `
		INSERT INTO transactions(order_id, amount, status, expires_at, paid_at, qr_url)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, order_id, amount, status, expires_at, paid_at, created_at, updated_at, qr_url
	`

	var t e.Transaction

	if err := trxRepo.db.QueryRowContext(
		ctx,
		q,
		trx.OrderId,
		trx.Amount,
		trx.Status,
		trx.ExpiresAt,
		trx.PaidAt,
		trx.QrUrl,
	).Scan(
		&t.ID,
		&t.OrderId,
		&t.Amount,
		&t.Status,
		&t.ExpiresAt,
		&t.PaidAt,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.QrUrl,
	); err != nil {
		return nil, err
	}

	return &t, nil
}

func (trxRepo *PostgresTransactionRepository) Update(ctx context.Context, trx e.Transaction) (*e.Transaction, error) {
	return nil, nil
}

func NewPostgresTransactionRepositoy(db *sql.DB) r.TransactionRepository {
	return &PostgresTransactionRepository{
		db: db,
	}
}
