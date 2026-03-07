package paymentuc

import (
	cx "context"

	conf "github.com/fkihai/payflow/internal/infrastructure/config"
)

type ConfirmCharge struct {
	repo ChargeRepo
	gw   Gateway
	cfg  *conf.PaymentGatewayConfig
}

func NewConfirmCharge(repo ChargeRepo, gw Gateway, cfg *conf.PaymentGatewayConfig) *ConfirmCharge {
	return &ConfirmCharge{
		repo: repo,
		gw:   gw,
		cfg:  cfg,
	}
}

func (confirm *ConfirmCharge) Execute(ctx cx.Context, payload []byte) error {
	he, err := confirm.gw.ConfirmCharge(ctx, payload)
	if err != nil {
		return err
	}

	// push worker

	// update database
	if err := confirm.repo.Update(ctx, he.OID, he.ChgStatus, he.ChgPaid); err != nil {
		return err
	}

	return nil
}
