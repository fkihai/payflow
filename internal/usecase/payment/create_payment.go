package payment

import (
	"context"

	e "github.com/fkihai/payflow/internal/domain/entity"
	r "github.com/fkihai/payflow/internal/domain/repository"
	p "github.com/fkihai/payflow/internal/infrastructure/payment"
)

type PaymentUsecase struct {
	repo    r.TransactionRepository
	gateway p.Gateway
	gen     p.QrisOrderIDGenerator
}

func NewPaymentUsecase(repo r.TransactionRepository, gateway p.Gateway) *PaymentUsecase {
	return &PaymentUsecase{
		repo:    repo,
		gateway: gateway,
	}
}

func (pu *PaymentUsecase) CreatePayment(ctx context.Context, trx p.CreateTransactionInput) (*e.Transaction, error) {

	orderID, err := pu.gen.Generate()
	if err != nil {
		return nil, err
	}
	req := p.GatewayRequest{
		OrderID: orderID,
		Amount:  trx.Amount,
	}

	// create trx payment gateway
	resp, err := pu.gateway.CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	// create to db
	trStore := e.Transaction{
		OrderId:   resp.OrderID,
		Amount:    resp.Amount,
		Status:    resp.Status,
		ExpiresAt: 0,
		QrUrl:     resp.QrCodeUrl,
	}

	t, err := pu.repo.Create(ctx, trStore)
	if err != nil {
		return nil, err
	}

	return t, nil
}
