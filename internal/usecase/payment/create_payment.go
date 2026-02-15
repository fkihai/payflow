package payment

import (
	"context"

	e "github.com/fkihai/payflow/internal/domain/entity"
	r "github.com/fkihai/payflow/internal/domain/repository"
)

type PaymentUsecase struct {
	repo r.TransactionRepository
}

func NewPaymentUsecase(repo r.TransactionRepository) *PaymentUsecase {
	return &PaymentUsecase{
		repo: repo,
	}
}

func (p *PaymentUsecase) CreatePayment(ctx context.Context, trx e.Transaction) (*e.Transaction, error) {
	t, err := p.repo.Create(ctx, trx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
