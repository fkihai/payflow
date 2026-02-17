package postgres

import (
	cx "context"
	"database/sql"
	"errors"
	"time"

	dm "github.com/fkihai/payflow/internal/domain"
	pu "github.com/fkihai/payflow/internal/usecase/payment"
)

type PostgresChargeRepo struct {
	db *sql.DB
}

func NewPostgresTransactionRepositoy(db *sql.DB) pu.ChargeRepo {
	return &PostgresChargeRepo{
		db: db,
	}
}

func (repo *PostgresChargeRepo) Create(ctx cx.Context, chgIn *dm.Charge) (*dm.Charge, error) {
	q := `
		INSERT INTO transactions(order_id, amount, status, expires_at, paid_at, qr_url)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, order_id, amount, status, expires_at, paid_at, created_at, updated_at, qr_url
	`

	var chgOut dm.Charge

	if err := repo.db.QueryRowContext(
		ctx,
		q,
		chgIn.OID,
		chgIn.Amount,
		chgIn.Status,
		chgIn.ExpiresAt,
		chgIn.PaidAt,
		chgIn.QrUrl,
	).Scan(
		&chgOut.ID,
		&chgOut.OID,
		&chgOut.Amount,
		&chgOut.Status,
		&chgOut.ExpiresAt,
		&chgOut.PaidAt,
		&chgOut.CreatedAt,
		&chgOut.UpdatedAt,
		&chgOut.QrUrl,
	); err != nil {
		return nil, err
	}

	return &chgOut, nil
}

func (repo *PostgresChargeRepo) FindByOID(ctx cx.Context, OID dm.OID) (*dm.Charge, error) {
	q := `
		SELECT id, order_id, amount, status, expires_at, paid_at, created_at, updated_at, qr_url
		FROM transactions
		WHERE order_id=$1
	`

	row := repo.db.QueryRowContext(ctx, q, OID)

	var chgOut dm.Charge
	if err := row.Scan(
		&chgOut.ID,
		&chgOut.OID,
		&chgOut.Amount,
		&chgOut.Status,
		&chgOut.ExpiresAt,
		&chgOut.PaidAt,
		&chgOut.CreatedAt,
		&chgOut.UpdatedAt,
		&chgOut.QrUrl,
	); err != nil {
		return nil, err
	}

	return &chgOut, nil
}

func (repo *PostgresChargeRepo) Update(ctx cx.Context, OID dm.OID, status dm.ChargeStatus, paidAt *time.Time) error {

	q := `
		UPDATE transactions
		SET status = $1, paid_at = $2
		WHERE order_id = $3 
		AND status = $4
	`
	result, err := repo.db.ExecContext(
		ctx,
		q,
		status,
		paidAt,
		OID,
		dm.ChargePending,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsAffected == 0 {
		return errors.New("transaction not updated (already process and not found)")
	}

	return nil
}
